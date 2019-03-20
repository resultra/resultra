// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package runtimeConfig

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"resultra/tracker/server/common/databaseWrapper"
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

type TransactionalEmailConfig struct {
	FromEmailAddr     string `json:"fromEmailAddress"`
	SMTPServerAddress string `json:"smtpServerAddress"`
	SMTPUserName      string `json:"smtpUserName"`
	SMTPPort          *int   `json:"smtpPort"`
	SMTPPassword      string `json:"smtpPassword"`
}

const defaultSMTPPort int = 587

type RuntimeConfig struct {
	FactoryTemplateDatabaseConfig *FactoryTemplateDatabaseConfig `json:"factoryTemplateDatabase"`
	TrackerDatabaseConfig         TrackerDatabaseConfig          `json:"trackerDatabase"`
	TransactionalEmailConfig      *TransactionalEmailConfig      `json:"transactionalEmail"`

	PortNumber            int  `json:"portNumber"`
	IsSingleUserWorkspace bool `json:"isSingleUserWorkspace"`
}

const permsOwnerReadWriteOnly os.FileMode = 0700

func NewDefaultRuntimeConfig() RuntimeConfig {
	config := RuntimeConfig{
		PortNumber:               defaultPortNum,
		IsSingleUserWorkspace:    false,
		TransactionalEmailConfig: nil}
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

	if config.TransactionalEmailConfig == nil {
		log.Println("WARNING: No configuration provided for transactional email. No email will be sent. Use for development only.")
	} else {
		emailConfig := config.TransactionalEmailConfig
		if len(emailConfig.FromEmailAddr) == 0 {
			// TODO - Do a full-blown validation of the email address format
			return fmt.Errorf("configuration missing return address for transactional emails")
		}
		if len(emailConfig.SMTPServerAddress) == 0 {
			return fmt.Errorf("configuration missing SMTP server address for transactionsl emails")
		}
		if len(emailConfig.SMTPPassword) == 0 {
			return fmt.Errorf("configuration missing SMTP password for transactionsl emails")
		}
		if len(emailConfig.SMTPUserName) == 0 {
			return fmt.Errorf("configuration missing SMTP user name for transactionsl emails")
		}

		if emailConfig.SMTPPort == nil {
			defaultPort := defaultSMTPPort
			config.TransactionalEmailConfig.SMTPPort = &defaultPort
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
