package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"database/sql"
	_ "github.com/lib/pq"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
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

func initDatabase() error {
	databaseHost := "localhost"
	databaseUserName := "postgres"
	databasePassword := "docker"

	createDatabaseAndWebUser := func() error {
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

	time.Sleep(2 * time.Second)

	createErr := createDatabaseAndWebUser()
	if createErr != nil {
		return fmt.Errorf("Error setting up tracker database: %v", createErr)
	}
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

	reader, err := cli.ImagePull(ctx, "docker.io/library/postgres", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)

	containerConfig := container.Config{
		Image: "postgres",
		//	Cmd:   []string{"echo", "hello world"},
		Tty: true}
	volumeBind := databaseDir + ":" + "/var/lib/postgresql/data" + ":rw"
	hostConfig := container.HostConfig{Binds: []string{volumeBind}}
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

	/*	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
		select {
		case err := <-errCh:
			if err != nil {
				panic(err)
			}
		case <-statusCh:
		}
	*/
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
