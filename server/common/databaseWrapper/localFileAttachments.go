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

func serveLocalFileAttachment(attachmentBasePath string, serveParams ServeAttachmentParams) {

	fullyQualifiedFile := fullyQualifiedLocalAttachmentFileName(attachmentBasePath,
		serveParams.ParentDatabaseID, serveParams.CloudFileName)

	// Serve file is a simple way to serve the attachment. Another way would be to use S3 or minio
	// as described here: https://github.com/minio/minio-go/issues/593

	http.ServeFile(serveParams.RespWriter, serveParams.HTTPReq, fullyQualifiedFile)

}
