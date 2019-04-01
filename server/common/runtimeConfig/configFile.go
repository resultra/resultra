// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package runtimeConfig

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
	"os"
	"resultra/tracker/server/common/databaseWrapper"
)

var defaultPortNum int = 43400

type FactoryTemplateDatabaseConfig struct {
	LocalDatabaseConfig         *databaseWrapper.LocalSQLiteTrackerDatabaseConnectionConfig `yaml:"localSQLiteDatabase,omitempty"`
	PostgresSingleAccountConfig *databaseWrapper.PostgresSingleAccountDatabaseConfig        `yaml:"postgresDatabase,omitempty"`
}

type TrackerDatabaseConfig struct {
	LocalDatabaseConfig         *databaseWrapper.LocalSQLiteTrackerDatabaseConnectionConfig `yaml:"localSQLiteDatabase,omitempty"`
	PostgresMultiAccountConfig  *databaseWrapper.PostgresMultipleAccountDatabaseConfig      `yaml:"postgresMultiAccountDatabase,omitempty"`
	PostgresSingleAccountConfig *databaseWrapper.PostgresSingleAccountDatabaseConfig        `yaml:"postgresDatabase,omitempty"`
	LocalAttachmentConfig       *databaseWrapper.LocalAttachmentStorageConfig               `yaml:"localAttachmentStorage,omitempty"`
}

type TransactionalEmailConfig struct {
	FromEmailAddr     string `yaml:"fromEmailAddress"`
	SMTPServerAddress string `yaml:"smtpServerAddress"`
	SMTPUserName      string `yaml:"smtpUserName"`
	SMTPPort          *int   `yaml:"smtpPort,omitempty"`
	SMTPPassword      string `yaml:"smtpPassword"`
}

const defaultSMTPPort int = 587

type RuntimeConfig struct {
	FactoryTemplateDatabaseConfig *FactoryTemplateDatabaseConfig `yaml:"factoryTemplateDatabase,omitempty"`
	TrackerDatabaseConfig         TrackerDatabaseConfig          `yaml:"trackerDatabase"`
	TransactionalEmailConfig      *TransactionalEmailConfig      `yaml:"transactionalEmail,omitempty"`

	ServerConfig          `yaml:"server"`
	IsSingleUserWorkspace bool `yaml:"isSingleUserWorkspace"`
}

const permsOwnerReadWriteOnly os.FileMode = 0700

type ServerConfig struct {
	ListenPortNumber    int     `yaml:"listenPortNumber"`
	SiteBaseURL         *string `yaml:"baseSiteURL,omitempty"`
	CookieAuthKey       string  `yaml:"cookieAuthenticationKey"`
	CookieEncryptionKey string  `yaml:"cookieEncryptionKey"`
}

func NewDefaultRuntimeConfig() RuntimeConfig {

	defaultServerConfig := ServerConfig{ListenPortNumber: defaultPortNum}

	config := RuntimeConfig{
		ServerConfig:             defaultServerConfig,
		IsSingleUserWorkspace:    false,
		TransactionalEmailConfig: nil}
	return config
}

var CurrRuntimeConfig RuntimeConfig

func (config *RuntimeConfig) SaveConfigFile(configFileName string) error {

	configData, marshalErr := yaml.Marshal(config)
	if marshalErr != nil {
		return fmt.Errorf("error saving config file %v: cannot marshal configuration: %v", configFileName, marshalErr)
	}

	writeConfigErr := ioutil.WriteFile(configFileName, configData, 0600)
	if writeConfigErr != nil {
		return fmt.Errorf("error saving config file %v: cannot save file: %v", configFileName, writeConfigErr)
	}
	return nil

}

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
	} else if config.TrackerDatabaseConfig.PostgresSingleAccountConfig != nil {
		log.Println("Initialize Postgres single database connection")
		if err := databaseWrapper.InitConnectionConfiguration(config.TrackerDatabaseConfig.PostgresSingleAccountConfig); err != nil {
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

	if config.ServerConfig.SiteBaseURL == nil {
		localHostURL := fmt.Sprintf("http://localhost:%v/", config.ServerConfig.ListenPortNumber)
		log.Printf("WARNING: No configuration provided for the base URL for the server, defaulting to %v. Use for development and testing only", localHostURL)
		config.ServerConfig.SiteBaseURL = &localHostURL
	}

	requiredCookieKeyLen := 32
	if len(config.ServerConfig.CookieAuthKey) != requiredCookieKeyLen {
		return fmt.Errorf("Missing/incorrect cookie authentication authentication key: %v: key must be 32 bytes")
	}
	if len(config.ServerConfig.CookieEncryptionKey) != requiredCookieKeyLen {
		return fmt.Errorf("Missing/incorrect cookie authentication encryption key: %v: key must be 32 bytes")
	}

	if config.TransactionalEmailConfig == nil {
		log.Println("WARNING: No configuration provided for transactional email. No email will be sent. Use for development and testing only.")
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

	configFile, readErr := ioutil.ReadFile(configFileName)
	if readErr != nil {
		return fmt.Errorf("unable to read configuration file %v: %v ", configFileName, readErr)
	}

	config := NewDefaultRuntimeConfig()
	unmarshalErr := yaml.Unmarshal(configFile, &config)
	if unmarshalErr != nil {
		return fmt.Errorf("error parsing config file %v: %v", configFileName, unmarshalErr)
	}

	err := InitRuntimeConfig(config)
	if err != nil {
		return fmt.Errorf("ERROR: Can't initialize configuration: error = %v: config file = %v", err, configFileName)
	}

	return nil
}
