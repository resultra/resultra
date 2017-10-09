package runtimeConfig

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type RuntimeConfig struct {
	DatabaseBasePath *string `json:"databaseBasePath"`
}

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

func newDefaultRuntimeConfig() RuntimeConfig {
	config := RuntimeConfig{}
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

	if err := config.validateWellFormedDatabaseBasePath(); err != nil {
		return fmt.Errorf("invalid runtime config: %v: %v", configFileName, err)
	}

	CurrRuntimeConfig = config

	return nil
}

func PrintCurrentConfig() {
	log.Printf("Runtime configuration: %+v", CurrRuntimeConfig)
}
