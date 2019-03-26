package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"resultra/tracker/server/common/databaseWrapper"
	"time"

	"database/sql"

	_ "github.com/lib/pq"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"golang.org/x/net/context"
)

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

func connectToResultraDB() (*sql.DB, error) {

	databaseHost := "localhost"
	databaseUserName := "postgres"
	databasePassword := "docker"

	connectStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		databaseHost, databaseUserName, databasePassword, "resultra")

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

	connectStr := fmt.Sprintf("host=%s user=%s password=%s sslmode=disable",
		databaseHost, databaseUserName, databasePassword)

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

func createDatabaseAndWebUser() error {

	databaseHost := "localhost"
	databaseUserName := "postgres"
	databasePassword := "docker"

	connectStr := fmt.Sprintf("host=%s user=%s password=%s sslmode=disable",
		databaseHost, databaseUserName, databasePassword)

	dbHandle, openErr := sql.Open("postgres", connectStr)
	if openErr != nil {
		return fmt.Errorf("can't establish connection to database: %v", openErr)
	}
	defer dbHandle.Close()

	_, createDBErr := dbHandle.Exec(`CREATE DATABASE resultra`)
	if createDBErr != nil {
		return fmt.Errorf("can't create main resultra database: %v", createDBErr)
	}

	// TODO - don't hard-code the password
	_, createUserErr := dbHandle.Exec(`CREATE USER resultra_web_user WITH NOSUPERUSER NOCREATEDB NOCREATEROLE PASSWORD 'resultrawebpw'`)
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
	conn, err := net.DialTimeout("tcp", "localhost:5432", timeout)
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

func initDatabase() error {

	if waitErr := waitPostgresPortReady(); waitErr != nil {
		return waitErr
	}

	log.Println("Starting database initialization ...")

	createErr := createDatabaseAndWebUser()
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

func configureDatabase() error {
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
	hostPortBinding := []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: "5432"}}
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

	initDatabase()

	if err := cli.ContainerStop(ctx, resp.ID, nil); err != nil {
		panic(err)
	}

	log.Printf("Postgres Container: ID = %v: stopped\n", resp.ID)

	return nil
}

func main() {
	dirsErr := createInstallDirs()
	if dirsErr != nil {
		log.Fatalf("ERROR: Setup failed: %v", dirsErr)
	}

	dbConfigErr := configureDatabase()
	if dbConfigErr != nil {
		log.Fatalf("ERROR: Setup failed: %v", dbConfigErr)
	}
}
