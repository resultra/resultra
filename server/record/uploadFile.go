package record

import (
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/appengine"
	"google.golang.org/cloud"
	"google.golang.org/cloud/storage"
	"io/ioutil"
	"log"
	"net/http"
)

type UploadFileResponse struct {
	UploadSuccess bool `json:"status"`
}

const cloudStorageBucketName string = "resultra-db-dev"
const cloudStorageJSONAuthInfoFile string = "/Users/sroehling/Development/Datasheet-Dev-60167588e163.json"

func getCloudClient(appEngContext context.Context) (*storage.Client, error) {

	jsonKey, err := ioutil.ReadFile(cloudStorageJSONAuthInfoFile)
	if err != nil {
		return nil, fmt.Errorf("getCloudContext: unable to read json auth info: %v", err)
	}
	conf, err := google.JWTConfigFromJSON(jsonKey, storage.ScopeFullControl)
	if err != nil {
		return nil, fmt.Errorf("getCloudContext: unable to create config from json auth info: %v", err)
	}

	client, clientErr := storage.NewClient(appEngContext, cloud.WithTokenSource(conf.TokenSource(appEngContext)))
	if clientErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to save to cloud storage: %v", clientErr)
	}

	return client, nil

}

func uploadFile(req *http.Request) (*UploadFileResponse, error) {
	formFile, handler, formErr := req.FormFile("file")
	if formErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to upload file: invalid api/form input in request: %v", formErr)
	}
	log.Printf("Uploading file: %v", handler.Filename)

	cloudFileName := handler.Filename

	fileContents, readErr := ioutil.ReadAll(formFile)
	if formErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to read file contents: %v", readErr)
	}

	fileLength := len(fileContents)
	log.Printf("Uploading file: got file contents: length (bytes) = %v", fileLength)

	appEngContext := appengine.NewContext(req)

	client, clientErr := getCloudClient(appEngContext)
	if clientErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to save to cloud storage: %v", clientErr)
	}
	defer client.Close()

	bucket := client.Bucket(cloudStorageBucketName)
	log.Printf("uploadFile: start uploading file to cloud: bucket=%v, file name = %v ...", cloudStorageBucketName, cloudFileName)

	cloudWriter := bucket.Object(cloudFileName).NewWriter(appEngContext)

	bytesWritten, writeErr := cloudWriter.Write(fileContents)
	if writeErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to save to cloud storage: %v", writeErr)
	} else if bytesWritten != fileLength {
		return nil, fmt.Errorf("uploadFile: Unable to save to cloud storage: only %v of %v bytes written", bytesWritten, fileLength)
	}

	// The file isn't actually written until Close() is called.
	if closeErr := cloudWriter.Close(); closeErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to save: failure closing cloud writer: %v", closeErr)
	}
	log.Printf("uploadFile: ... done uloading file to cloud: bucket=%v, file name = %v", cloudStorageBucketName, cloudFileName)

	return &UploadFileResponse{true}, nil
}
