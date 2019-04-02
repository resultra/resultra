// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/common/runtimeConfig"
	"strings"
	"time"

	"database/sql"

	_ "github.com/lib/pq"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"golang.org/x/net/context"
)

func promptUserInputString(prompt string, defaultInput string) string {

	readUserInput := func() string {

		if len(defaultInput) > 0 {
			fmt.Printf("%v [%v]:", prompt, defaultInput)
		} else {
			fmt.Printf("%v:", prompt)
		}

		reader := bufio.NewReader(os.Stdin)
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimRight(userInput, "\n")

		if len(userInput) == 0 && len(defaultInput) > 0 {
			return defaultInput
		} else if len(userInput) > 0 {
			return userInput
		} else {
			return ""
		}
	}

	userInput := ""
	for userInput = readUserInput(); userInput == ""; userInput = readUserInput() {
	}

	return userInput
}

func makeInstallDir(dirLabel string, dirName string, perm os.FileMode) error {
	_, statErr := os.Stat(dirName)
	if os.IsNotExist(statErr) {
		fmt.Printf("Creating install directory: %v: %v ... ", dirLabel, dirName)
		err := os.Mkdir(dirName, perm)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			return fmt.Errorf("error creating installation directory: %v: %v: %v", dirLabel, dirName, err)
		}
		fmt.Printf("done.\n")
	} else {
		fmt.Printf("Skipping directory creation. Directory already exists: %v: %v ... skipped\n", dirLabel, dirName)
	}
	return nil
}

const topLevelInstallDir string = "/var/resultra"
const appDataDir string = "/var/resultra/appdata"
const attachmentDir string = "/var/resultra/appdata/attachments"
const databaseDir string = "/var/resultra/appdata/database"
const configDir string = "/var/resultra/config"

func createInstallDirs() error {

	topLevelDirErr := makeInstallDir("top level directory", topLevelInstallDir, 0755)
	if topLevelDirErr != nil {
		return topLevelDirErr
	}

	appDataDirErr := makeInstallDir("application data directory", appDataDir, 0755)
	if appDataDirErr != nil {
		return appDataDirErr
	}

	attachmentDirErr := makeInstallDir("attachment data directory", attachmentDir, 0755)
	if appDataDirErr != nil {
		return attachmentDirErr
	}

	databaseDirErr := makeInstallDir("database data directory", databaseDir, 0755)
	if databaseDirErr != nil {
		return databaseDirErr
	}

	configDirErr := makeInstallDir("configuration directory", configDir, 0755)
	if configDirErr != nil {
		return configDirErr
	}

	return nil
}

// By default, Postgres runs on port 5432. If this setup program is running on a system where
// Postgres is already running, then we need to map the default Postgres port to a higher
// numbered port for the setup itself.
const setupDBPort int = 43401

func connectToResultraDB() (*sql.DB, error) {

	databaseHost := "localhost"
	databaseUserName := "postgres"
	databasePassword := "docker"

	connectStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		databaseHost, setupDBPort, databaseUserName, databasePassword, "resultra")

	dbHandle, openErr := sql.Open("postgres", connectStr)
	if openErr != nil {
		return nil, fmt.Errorf("can't establish connection to database: %v", openErr)
	}

	// Open doesn't directly open the database connection. To verify the connection, the Ping() function
	// is needed.
	if pingErr := dbHandle.Ping(); pingErr != nil {
		return nil, fmt.Errorf("can't establish connection to database (ping failed): %v", pingErr)
	}

	// TODO - Add more robustness to this function. Ping will establish the connection, but doesn't actually
	// test it. To test it, running an end-to-end query like "SELECT 1 as testcol" would be helpful.

	return dbHandle, nil
}

func initDatabaseTables(dbHandle *sql.DB) error {

	initErr := databaseWrapper.InitNewTrackerDatabaseToDest(dbHandle)
	if initErr != nil {
		return fmt.Errorf("can't initialize tracker database: %v", initErr)
	}

	return nil
}

func grantWebUserPerms(dbHandle *sql.DB) error {

	permsSchema := `REVOKE ALL ON DATABASE resultra FROM public;` +
		`GRANT CONNECT ON DATABASE resultra TO resultra_web_user;` +
		`REVOKE ALL ON SCHEMA public FROM public;` +
		`GRANT USAGE ON SCHEMA public TO resultra_web_user;` +
		`REVOKE ALL ON ALL TABLES IN SCHEMA public FROM PUBLIC;` +
		`GRANT SELECT,UPDATE,INSERT,DELETE ON ALL TABLES in SCHEMA public to resultra_web_user;` +
		`ALTER DEFAULT PRIVILEGES FOR ROLE resultra_web_user IN SCHEMA public GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO resultra_web_user;`

	if _, permsErr := dbHandle.Exec(permsSchema); permsErr != nil {
		return fmt.Errorf("can't initialize tracker database permissions: %v", permsErr)
	}

	return nil
}

func connectAdminPostgresUser() (*sql.DB, error) {
	databaseHost := "localhost"
	databaseUserName := "postgres"
	databasePassword := "docker"

	connectStr := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		databaseHost, setupDBPort, databaseUserName, databasePassword)

	dbHandle, openErr := sql.Open("postgres", connectStr)
	if openErr != nil {
		return nil, fmt.Errorf("can't establish connection to database: %v", openErr)
	}

	// Open doesn't directly open the database connection. To verify the connection, the Ping() function
	// is needed.
	if pingErr := dbHandle.Ping(); pingErr != nil {
		return nil, fmt.Errorf("can't establish connection to database (ping failed): %v", pingErr)
	}

	return dbHandle, nil

}

func createDatabaseAndWebUser(dbPassword string) error {

	databaseHost := "localhost"
	databaseUserName := "postgres"
	databasePassword := "docker"

	connectStr := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		databaseHost, setupDBPort, databaseUserName, databasePassword)

	dbHandle, openErr := sql.Open("postgres", connectStr)
	if openErr != nil {
		return fmt.Errorf("can't establish connection to database: %v", openErr)
	}
	defer dbHandle.Close()

	_, createDBErr := dbHandle.Exec(`CREATE DATABASE resultra`)
	if createDBErr != nil {
		return fmt.Errorf("can't create main resultra database: %v", createDBErr)
	}

	createUserQuery := fmt.Sprintf("CREATE USER resultra_web_user WITH NOSUPERUSER NOCREATEDB NOCREATEROLE PASSWORD '%v'", dbPassword)
	_, createUserErr := dbHandle.Exec(createUserQuery)
	if createUserErr != nil {
		return fmt.Errorf("can't create main resultra database: %v", createUserErr)
	}
	return nil
}

func waitPostgresPortReady() error {

	// When starting up the Postgres container, Docker will also try to connect to the Postgres
	// port for port mapping. To ensure there is not a conflict with the Postgres continer
	// start-up, wait a fixed period before trying to connect to the port.
	time.Sleep(10 * time.Second)

	timeout := time.Duration(60 * time.Second)
	waitPosgresSetupPortStr := fmt.Sprintf("localhost:%d", setupDBPort)
	conn, err := net.DialTimeout("tcp", waitPosgresSetupPortStr, timeout)
	if err != nil {
		return fmt.Errorf("error connecting to Postgres docker image for initialization: %v", err)
	}

	if conn != nil {
		conn.Close()
		log.Printf("Connection to Postgres docker image successfully established")
		return nil
	} else {
		return fmt.Errorf("error connecting to Postgres docker image for initialization: connection could not be establishd")
	}
}

func initDatabase(dbPassword string) error {

	if waitErr := waitPostgresPortReady(); waitErr != nil {
		return waitErr
	}

	log.Println("Starting database initialization ...")

	createErr := createDatabaseAndWebUser(dbPassword)
	if createErr != nil {
		return fmt.Errorf("Error setting up tracker database: %v", createErr)
	}

	resultraDBHandle, dbErr := connectToResultraDB()
	if dbErr != nil {
		log.Fatalf("Error setting up tracker database: %v", createErr)
	}
	defer resultraDBHandle.Close()

	initErr := initDatabaseTables(resultraDBHandle)
	if createErr != nil {
		log.Fatalf("Error initializing tracker database tables: %v", initErr)
	}

	grantErr := grantWebUserPerms(resultraDBHandle)
	if grantErr != nil {
		log.Fatalf("Error initializing tracker database tables: %v", grantErr)
	}

	log.Println("Done with database initialization.")

	return nil

}

func configureDatabase(dbPassword string) error {
	ctx := context.Background()

	//cli, err := client.NewEnvClient()
	// TBD: Without "pinning" the API client version, the daemon might return
	// an error message like: client version 1.40 is too new. Maximum supported API version is 1.39
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		panic(err)
	}

	reader, err := cli.ImagePull(ctx, "postgres:latest", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)

	exposedPorts := nat.PortSet{"5432/tcp": struct{}{}}
	containerConfig := container.Config{
		Image:        "postgres",
		ExposedPorts: exposedPorts,
		Tty:          true,
		AttachStdout: true,
		AttachStderr: true}

	volumeBind := databaseDir + ":" + "/var/lib/postgresql/data" + ":rw"
	containerPort := nat.Port("5432/tcp")
	hostSetupPortStr := fmt.Sprintf("%d", setupDBPort)
	hostPortBinding := []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: hostSetupPortStr}}
	portMap := nat.PortMap{containerPort: hostPortBinding}
	hostConfig := container.HostConfig{
		Binds:        []string{volumeBind},
		PortBindings: portMap}
	resp, err := cli.ContainerCreate(ctx, &containerConfig, &hostConfig, nil, "")
	if err != nil {
		panic(err)
	}
	log.Printf("Created container with ID = %v\n", resp.ID)

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	log.Printf("Container started: ID = %v: waiting for database initialization\n", resp.ID)

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, out)

	initDatabase(dbPassword)

	if err := cli.ContainerStop(ctx, resp.ID, nil); err != nil {
		panic(err)
	}
	log.Printf("Postgres Container: ID = %v: stopped\n", resp.ID)

	return nil
}

func setupEmailConfig(emailConfig *runtimeConfig.TransactionalEmailConfig) {

	emailConfig.FromEmailAddr = promptUserInputString("Return email address (e.g. noreply@yourdomain.com", "")
	emailConfig.SMTPServerAddress = promptUserInputString("SMTP email server address (e.g. smtp.mailgun.org)", "")
	emailConfig.SMTPUserName = promptUserInputString("SMTP user name (for login to SMTP server)", "")
	emailConfig.SMTPPassword = promptUserInputString("Email password (for SMTP server)", "")
}

func generateRandomKey(keyLen int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	randStrRunes := func(n int) string {
		b := make([]rune, n)
		for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}
		return string(b)
	}
	return randStrRunes(keyLen)
}

// In a production environment, the server defaults to running on the standard HTTP
// port. This can be mapped to a different using a reverse proxy setup, and Docker
// port mapping.
const defaultServerListenPort int = 80

func setupServerConfig(serverConfig *runtimeConfig.ServerConfig) {
	cookieKeyLen := 32
	serverConfig.CookieAuthKey = generateRandomKey(cookieKeyLen)
	serverConfig.CookieEncryptionKey = generateRandomKey(cookieKeyLen)
	serverConfig.ListenPortNumber = defaultServerListenPort
}

func setupTrackerConfig(trackerConfig *runtimeConfig.TrackerDatabaseConfig, dbPassword string) {

	trackerConfig.PostgresSingleAccountConfig = &databaseWrapper.PostgresSingleAccountDatabaseConfig{
		TrackerDBHostName: "resultra-postgres", // resultra-postgres is the network name given to the Docker database container.
		TrackerUserName:   "resultra_web_user",
		TrackerPassword:   dbPassword,
		TrackerDBName:     "resultra"}
	trackerConfig.LocalAttachmentConfig = &databaseWrapper.LocalAttachmentStorageConfig{
		AttachmentBasePath: "/var/resultra/appdata/attachments"}

}

func setupConfig(dbPassword string) error {

	config := runtimeConfig.RuntimeConfig{}

	emailConfig := runtimeConfig.TransactionalEmailConfig{}
	config.TransactionalEmailConfig = &emailConfig
	setupEmailConfig(config.TransactionalEmailConfig)

	// The factory templates are referenced from a fixed location inside the
	// docker image for the server.
	localFactoryTemplates := databaseWrapper.LocalSQLiteTrackerDatabaseConnectionConfig{
		DatabaseBasePath: "/usr/local/resultra/factoryTemplates"}
	factoryTemplateConfig := runtimeConfig.FactoryTemplateDatabaseConfig{LocalDatabaseConfig: &localFactoryTemplates}
	config.FactoryTemplateDatabaseConfig = &factoryTemplateConfig

	setupServerConfig(&config.ServerConfig)

	setupTrackerConfig(&config.TrackerDatabaseConfig, dbPassword)

	configFileName := configDir + "/" + "resultraConfig.yaml"

	saveErr := config.SaveConfigFile(configFileName)
	if saveErr != nil {
		return fmt.Errorf("Error saving config file: %v", saveErr)
	}

	return nil

}

func main() {
	dirsErr := createInstallDirs()
	if dirsErr != nil {
		log.Fatalf("ERROR: Setup failed: %v", dirsErr)
	}

	dbPassword := generateRandomKey(16)
	configErr := setupConfig(dbPassword)
	if configErr != nil {
		log.Fatalf("ERROR: Setup of configuration failed: %v", configErr)
	}

	dbConfigErr := configureDatabase(dbPassword)
	if dbConfigErr != nil {
		log.Fatalf("ERROR: Setup failed: %v", dbConfigErr)
	}
}
