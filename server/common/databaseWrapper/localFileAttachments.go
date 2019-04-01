// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package databaseWrapper

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const attachmentPermsOwnerReadWriteOnly os.FileMode = 0700

func fullyQualifiedLocalAttachmentFileName(attachmentBasePath string, databaseID string, fileName string) string {
	return attachmentBasePath + "/" + databaseID + "/" + fileName
}

func fullyQualifiedLocalAttachmentPath(attachmentBasePath string, databaseID string) string {
	return attachmentBasePath + "/" + databaseID
}

func initLocalAttachmentBasePath(attachmentBasePath string) error {

	err := os.MkdirAll(attachmentBasePath, permsOwnerReadWriteOnly)
	if err != nil {
		return fmt.Errorf("Error initializing attachment directory %v: %v",
			attachmentBasePath, err)
	}
	return nil
}

func saveLocalAttachmentFile(attachmentBasePath string, databaseID string, fileName string, fileData []byte) error {

	fullyQualifiedPath := fullyQualifiedLocalAttachmentPath(attachmentBasePath, databaseID)

	err := os.MkdirAll(fullyQualifiedPath, attachmentPermsOwnerReadWriteOnly)
	if err != nil {
		return fmt.Errorf("Error initializing attachment directory %v: %v",
			fullyQualifiedPath, err)
	}

	fullyQualifiedFile := fullyQualifiedLocalAttachmentFileName(attachmentBasePath, databaseID, fileName)

	writeErr := ioutil.WriteFile(fullyQualifiedFile, fileData, permsOwnerReadWriteOnly)
	if writeErr != nil {
		return fmt.Errorf("SaveCloudFile: Error writing file: filename = %v, error = %v", fileName, writeErr)
	}

	log.Printf("uploadFile: ... done uloading file to cloud: file name = %v", fileName)

	return nil

}

type LocalAttachmentStorageConfig struct {
	AttachmentBasePath string `yaml:"attachmentBasePath"`
}

func serveLocalFileAttachment(attachmentBasePath string, serveParams ServeAttachmentParams) {

	fullyQualifiedFile := fullyQualifiedLocalAttachmentFileName(attachmentBasePath,
		serveParams.ParentDatabaseID, serveParams.CloudFileName)

	// Serve file is a simple way to serve the attachment. Another way would be to use S3 or minio
	// as described here: https://github.com/minio/minio-go/issues/593

	http.ServeFile(serveParams.RespWriter, serveParams.HTTPReq, fullyQualifiedFile)

}

func (config LocalAttachmentStorageConfig) validateWellFormedAttachmentBasePath() error {

	if len(config.AttachmentBasePath) == 0 {
		return fmt.Errorf("configuration file missing database path configuration")
	}
	return nil

}

func (config LocalAttachmentStorageConfig) Init() error {
	if err := config.validateWellFormedAttachmentBasePath(); err != nil {
		return fmt.Errorf("LocalAttachmentStorageConfig.Init: %v", err)
	}

	if err := initLocalAttachmentBasePath(config.AttachmentBasePath); err != nil {
		return fmt.Errorf("LocalAttachmentStorageConfig.Init: %v", err)
	}

	log.Printf("Local directory initialized for attachments: base path = %v", config.AttachmentBasePath)
	return nil
}

func (config LocalAttachmentStorageConfig) GetAttachmentBasePath(r *http.Request) (string, error) {
	if err := config.validateWellFormedAttachmentBasePath(); err != nil {
		panic(fmt.Sprintf("runtime config: tried to retrieve attachment path from invalid config: %v", err))
	}
	return config.AttachmentBasePath, nil
}

func (config LocalAttachmentStorageConfig) SaveAttachment(saveParams SaveAttachmentParams) error {

	attachmentBasePath, err := config.GetAttachmentBasePath(saveParams.HTTPReq)
	if err != nil {
		return fmt.Errorf("LocalSQLiteTrackerDatabaseConnectionConfig: can't get base path: %v", err)
	}

	if err := saveLocalAttachmentFile(attachmentBasePath, saveParams.ParentDatabaseID,
		saveParams.CloudFileName, saveParams.FileData); err != nil {
		return fmt.Errorf("LocalSQLiteTrackerDatabaseConnectionConfig: can't save file: %v", err)
	}

	return nil

}

func (config LocalAttachmentStorageConfig) ServeAttachment(serveParams ServeAttachmentParams) {

	attachmentBasePath, err := config.GetAttachmentBasePath(serveParams.HTTPReq)
	if err != nil {
		http.Error(serveParams.RespWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	serveLocalFileAttachment(attachmentBasePath, serveParams)

}
