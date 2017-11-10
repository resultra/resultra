package databaseWrapper

import (
	"database/sql"
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

// Interface for handling connections to different types of databases
type TrackerDatabaseConnection interface {
	InitConnection() error
	GetTrackerDBHandle(r *http.Request) (*sql.DB, error)
	GetAttachmentBasePath(r *http.Request) (string, error)
	SaveAttachment(saveParams SaveAttachmentParams) error
	ServeAttachment(serveParams ServeAttachmentParams)
}

var dbConnection TrackerDatabaseConnection

func GetTrackerDatabaseHandle(r *http.Request) (*sql.DB, error) {

	if dbConnection == nil {
		return nil, fmt.Errorf("GetTrackerDatabaseHandle: uninitialized database connection")
	}
	return dbConnection.GetTrackerDBHandle(r)

}

func GetTrackerAttachmentBasePath(r *http.Request) (string, error) {
	if dbConnection == nil {
		return "", fmt.Errorf("GetTrackerDatabaseHandle: uninitialized database connection")
	}
	return dbConnection.GetAttachmentBasePath(r)
}

func SaveAttachment(saveParams SaveAttachmentParams) error {

	if dbConnection == nil {
		return fmt.Errorf("GetTrackerDatabaseHandle: uninitialized database connection")
	}

	return dbConnection.SaveAttachment(saveParams)
}

func ServeAttachment(serveParams ServeAttachmentParams) {

	if dbConnection == nil {
		errorMsg := "GetTrackerDatabaseHandle: uninitialized database connection"
		http.Error(serveParams.RespWriter, errorMsg, http.StatusInternalServerError)
	}

	dbConnection.ServeAttachment(serveParams)
}

func InitConnectionConfiguration(conn TrackerDatabaseConnection) error {

	if err := conn.InitConnection(); err != nil {
		return fmt.Errorf("InitConnectionConfiguration: %v", err)
	}

	dbConnection = conn

	return nil

}
