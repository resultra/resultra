package main

import (
	"fmt"
	"io"
	"log"
	"os"

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

func configureDatabase() error {
	ctx := context.Background()

	//	cli, err := client.NewEnvClient()
	// TBD: Without "pinning" the API client version, the daemon might return
	// an error message like: client version 1.40 is too new. Maximum supported API version is 1.39
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		panic(err)
	}

	reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"echo", "hello world"},
		Tty:   true,
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, out)

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
