package runtimeConfig

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"resultra/datasheet/server/common/databaseWrapper"
)

var defaultPortNum int = 43400

type FactoryTemplateDatabaseConfig struct {
	LocalDatabaseConfig         *databaseWrapper.LocalSQLiteTrackerDatabaseConnectionConfig `json:"localSQLiteDatabase"`
	PostgresSingleAccountConfig *databaseWrapper.PostgresSingleAccountDatabaseConfig        `json:"postgresDatabase"`
}

type TrackerDatabaseConfig struct {
	LocalDatabaseConfig        *databaseWrapper.LocalSQLiteTrackerDatabaseConnectionConfig `json:"localSQLiteDatabase"`
	PostgresMultiAccountConfig *databaseWrapper.PostgresMultipleAccountDatabaseConfig      `json:"postgresMultiAccountDatabase"`
	LocalAttachmentConfig      *databaseWrapper.LocalAttachmentStorageConfig               `json:"localAttachmentStorage"`
}

type RuntimeConfig struct {
	FactoryTemplateDatabaseConfig *FactoryTemplateDatabaseConfig `json:"factoryTemplateDatabase"`
	TrackerDatabaseConfig         TrackerDatabaseConfig          `json:"trackerDatabase"`

	PortNumber            int  `json:"portNumber"`
	IsSingleUserWorkspace bool `json:"isSingleUserWorkspace"`
}

const permsOwnerReadWriteOnly os.FileMode = 0700

func NewDefaultRuntimeConfig() RuntimeConfig {
	config := RuntimeConfig{
		PortNumber:            defaultPortNum,
		IsSingleUserWorkspace: false}
	return config
}

var CurrRuntimeConfig RuntimeConfig

func init() {
	CurrRuntimeConfig = NewDefaultRuntimeConfig()
}

func InitRuntimeConfig(config RuntimeConfig) error {

	if config.TrackerDatabaseConfig.LocalDatabaseConfig != nil {
		log.Println("Initializing local database connection")
		if err := databaseWrapper.InitConnectionConfiguration(config.TrackerDatabaseConfig.LocalDatabaseConfig); err != nil {
			return err
		}
	} else if config.TrackerDatabaseConfig.PostgresMultiAccountConfig != nil {
		log.Println("Initialize Postgres multi-account database connection")
		if err := databaseWrapper.InitConnectionConfiguration(config.TrackerDatabaseConfig.PostgresMultiAccountConfig); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("runtime configuration missing database connection configuration")
	}

	if config.TrackerDatabaseConfig.LocalAttachmentConfig != nil {
		log.Println("Initializing local attachment storage")
		if err := databaseWrapper.InitAttachmentStorageConfiguration(config.TrackerDatabaseConfig.LocalAttachmentConfig); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("runtime configuration missing attachment storage configuration")
	}

	// Optional database configuration for templates
	if config.FactoryTemplateDatabaseConfig != nil {
		log.Println("Initializing factory templates database connnection")
		if config.FactoryTemplateDatabaseConfig.LocalDatabaseConfig != nil {
			log.Println("Initializing local database connection for factory templates")
			if err := databaseWrapper.InitFactoryTemplateConnectionConfiguration(
				config.FactoryTemplateDatabaseConfig.LocalDatabaseConfig); err != nil {
				return err
			}
		} else if config.FactoryTemplateDatabaseConfig.PostgresSingleAccountConfig != nil {
			log.Println("Initialize Postgres multi-account database connection for factory templates")
			if err := databaseWrapper.InitFactoryTemplateConnectionConfiguration(
				config.FactoryTemplateDatabaseConfig.PostgresSingleAccountConfig); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("runtime configuration missing database connection configuration for factory templates")
		}
	}

	CurrRuntimeConfig = config

	return nil

}

func InitConfigFromConfigFile(configFileName string) error {

	configFile, openErr := os.Open(configFileName)
	if openErr != nil {
		return fmt.Errorf("init runtime config: can't open config file %v: %v", configFileName, openErr)
	}

	decoder := json.NewDecoder(configFile)

	config := NewDefaultRuntimeConfig()
	if err := decoder.Decode(&config); err != nil {
		return fmt.Errorf("init runtime config: %v, %v", configFileName, err)
	}

	err := InitRuntimeConfig(config)
	if err != nil {
		return fmt.Errorf("ERROR: Can't initialize configuration: error = %v: config file = %v", configFileName)
	}

	return nil
}
