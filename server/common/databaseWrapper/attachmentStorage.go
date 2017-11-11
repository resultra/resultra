package databaseWrapper

import (
	"fmt"
	"net/http"
)

type SaveAttachmentParams struct {
	CloudFileName    string
	ParentDatabaseID string
	HTTPReq          *http.Request
	FileData         []byte
}

type ServeAttachmentParams struct {
	RespWriter       http.ResponseWriter
	HTTPReq          *http.Request
	ParentDatabaseID string
	CloudFileName    string
}

type TrackerAttachmentStorageConnection interface {
	Init() error
	GetAttachmentBasePath(r *http.Request) (string, error)
	SaveAttachment(saveParams SaveAttachmentParams) error
	ServeAttachment(serveParams ServeAttachmentParams)
}

var attachmentStorage TrackerAttachmentStorageConnection

func GetTrackerAttachmentBasePath(r *http.Request) (string, error) {
	if attachmentStorage == nil {
		return "", fmt.Errorf("ServeAttachment: uninitialized attachment storage")
	}
	return attachmentStorage.GetAttachmentBasePath(r)
}

func SaveAttachment(saveParams SaveAttachmentParams) error {

	if attachmentStorage == nil {
		return fmt.Errorf("ServeAttachment: uninitialized attachment storage")
	}

	return attachmentStorage.SaveAttachment(saveParams)
}

func ServeAttachment(serveParams ServeAttachmentParams) {

	if attachmentStorage == nil {
		errorMsg := "ServeAttachment: uninitialized attachment storage"
		http.Error(serveParams.RespWriter, errorMsg, http.StatusInternalServerError)
	}

	attachmentStorage.ServeAttachment(serveParams)
}

func InitAttachmentStorageConfiguration(attachStorage TrackerAttachmentStorageConnection) error {
	if err := attachStorage.Init(); err != nil {
		return fmt.Errorf("InitAttachmentStorageConfiguration: %v", err)
	}

	attachmentStorage = attachStorage

	return nil
}
