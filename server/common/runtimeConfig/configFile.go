package runtimeConfig

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"resultra/datasheet/server/common/databaseWrapper"
)

var defaultPortNum int = 43400

type RuntimeConfig struct {
	LocalDatabaseConfig        *databaseWrapper.LocalSQLiteTrackerDatabaseConnectionConfig `json:"localSQLiteDatabase"`
	PostgresMultiAccountConfig *databaseWrapper.PostgresMultipleAccountDatabaseConfig      `json:"postgresMultiAccountDatabase"`

	LocalAttachmentConfig *databaseWrapper.LocalAttachmentStorageConfig `json:"localAttachmentStorage"`

	PortNumber int `json:"portNumber"`
}

const permsOwnerReadWriteOnly os.FileMode = 0700

func newDefaultRuntimeConfig() RuntimeConfig {
	config := RuntimeConfig{
		PortNumber: defaultPortNum}
	return config
}

var CurrRuntimeConfig RuntimeConfig

func init() {
	CurrRuntimeConfig = newDefaultRuntimeConfig()
}

func InitConfig(configFileName string) error {

	configFile, openErr := os.Open(configFileName)
	if openErr != nil {
		return fmt.Errorf("init runtime config: can't open config file %v: %v", configFileName, openErr)
	}

	decoder := json.NewDecoder(configFile)

	config := newDefaultRuntimeConfig()
	if err := decoder.Decode(&config); err != nil {
		return fmt.Errorf("init runtime config: %v, %v", configFileName, err)
	}

	if config.LocalDatabaseConfig != nil {
		log.Println("Initializing local database connection")
		if err := databaseWrapper.InitConnectionConfiguration(config.LocalDatabaseConfig); err != nil {
			return err
		}
	} else if config.PostgresMultiAccountConfig != nil {
		log.Println("Initialize Postgres multi-account database connection")
		if err := databaseWrapper.InitConnectionConfiguration(config.PostgresMultiAccountConfig); err != nil {
			return err
		}

	} else {
		return fmt.Errorf("runtime configuration %v missing database connection configuration", configFileName)
	}

	if config.LocalAttachmentConfig != nil {
		log.Println("Initializing local attachment storage")
		if err := databaseWrapper.InitAttachmentStorageConfiguration(config.LocalAttachmentConfig); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("runtime configuration %v missing attachment storage configuration", configFileName)
	}

	CurrRuntimeConfig = config

	return nil
}
