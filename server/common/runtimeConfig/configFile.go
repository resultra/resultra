package runtimeConfig

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var defaultPortNum int = 43400

type RuntimeConfig struct {
	DatabaseBasePath *string `json:"databaseBasePath"`
	PortNumber       int     `json:"portNumber"`
}

const permsOwnerReadWriteOnly os.FileMode = 0700

func (config RuntimeConfig) validateWellFormedDatabaseBasePath() error {

	if config.DatabaseBasePath == nil {
		return fmt.Errorf("configuration file missing database path configuration")
	}
	if len(*config.DatabaseBasePath) == 0 {
		return fmt.Errorf("configuration file missing database path configuration")
	}
	return nil

}

func (config RuntimeConfig) AttachmentBasePath() string {
	if err := config.validateWellFormedDatabaseBasePath(); err != nil {
		panic(fmt.Sprintf("runtime config: tried to retrieve attachment path from invalid config: %v", err))
	}
	return (*config.DatabaseBasePath) + `/attachments`
}

func (config RuntimeConfig) TrackerDatabaseFileName() string {
	if err := config.validateWellFormedDatabaseBasePath(); err != nil {
		panic(fmt.Sprintf("runtime config: tried to database path from invalid config: %v", err))
	}
	return (*config.DatabaseBasePath) + `/trackers.db`
}

func newDefaultRuntimeConfig() RuntimeConfig {
	config := RuntimeConfig{
		PortNumber: defaultPortNum}
	return config
}

var CurrRuntimeConfig RuntimeConfig

func init() {
	CurrRuntimeConfig = newDefaultRuntimeConfig()
}

func (config RuntimeConfig) initDatabaseBasePath() error {

	if err := config.validateWellFormedDatabaseBasePath(); err != nil {
		return fmt.Errorf("runtime config: tried to database path from invalid config: %v", err)
	}

	err := os.MkdirAll(*config.DatabaseBasePath, permsOwnerReadWriteOnly)
	if err != nil {
		return fmt.Errorf("Error initializing tracker directory %v: %v",
			config.DatabaseBasePath, err)
	}
	return nil
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

	if err := config.validateWellFormedDatabaseBasePath(); err != nil {
		return fmt.Errorf("invalid runtime config: %v: %v", configFileName, err)
	}
	if err := config.initDatabaseBasePath(); err != nil {
		return fmt.Errorf("configuration error: unable to create path for tracker database: %v: %v", configFileName, err)
	}
	CurrRuntimeConfig = config

	return nil
}

func PrintCurrentConfig() {
	log.Printf("Runtime configuration: %+v", CurrRuntimeConfig)
}
