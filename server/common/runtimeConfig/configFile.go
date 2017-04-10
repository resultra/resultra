package runtimeConfig

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type RuntimeConfig struct {
	AttachmentBasePath string `json:"attachmentBasePath"`
}

func newDefaultRuntimeConfig() RuntimeConfig {
	config := RuntimeConfig{
		AttachmentBasePath: `/private/tmp/`}
	return config
}

var CurrRuntimeConfig RuntimeConfig

func init() {
	CurrRuntimeConfig = newDefaultRuntimeConfig()
}

func InitConfig(configFileName string) error {

	configFile, openErr := os.Open(configFileName)
	if openErr != nil {
		return fmt.Errorf("Init runtime config: can't open config file %v: %v", configFileName, openErr)
	}

	decoder := json.NewDecoder(configFile)

	config := newDefaultRuntimeConfig()
	if err := decoder.Decode(&config); err != nil {
		return fmt.Errorf("Init runtime config: %v, %v", configFileName, err)
	}

	CurrRuntimeConfig = config

	return nil
}

func PrintCurrentConfig() {
	log.Printf("Runtime configuration: %+v", CurrRuntimeConfig)
}
