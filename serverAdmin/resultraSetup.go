package main

import (
	"fmt"
	"log"
	"os"
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

func main() {
	dirsErr := createInstallDirs()
	if dirsErr != nil {
		log.Fatalf("ERROR: Setup failed: %v", dirsErr)
	}
}
