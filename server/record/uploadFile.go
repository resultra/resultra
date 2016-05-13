package record

import (
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/appengine"
	"google.golang.org/cloud"
	"google.golang.org/cloud/storage"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type UploadFile struct {
	Name  string `json:"name"`
	Size  int    `json:"size"`
	Error string `json"error,omitempty"`
	Url   string `json:"url"`
}

type UploadFileResponse struct {
	Files []UploadFile `json:"files"`
}

const cloudStorageBucketName string = "resultra-db-dev"
const cloudStorageJSONAuthInfoFile string = "/Users/sroehling/Development/Datasheet-Dev-60167588e163.json"

func readServiceAuthConfig(jsonKeyFileName string) (*jwt.Config, error) {
	jsonKey, err := ioutil.ReadFile(cloudStorageJSONAuthInfoFile)
	if err != nil {
		return nil, fmt.Errorf("getCloudContext: unable to read json auth info: %v", err)
	}
	conf, err := google.JWTConfigFromJSON(jsonKey, storage.ScopeFullControl)
	if err != nil {
		return nil, fmt.Errorf("getCloudContext: unable to create config from json auth info: %v", err)
	}
	return conf, nil
}

func getCloudClient(appEngContext context.Context, config *jwt.Config) (*storage.Client, error) {

	client, clientErr := storage.NewClient(appEngContext, cloud.WithTokenSource(config.TokenSource(appEngContext)))
	if clientErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to save to cloud storage: %v", clientErr)
	}

	return client, nil
}

func getSignedURL(bucketName string, fileName string, authConfig *jwt.Config, expirationSecs int) (string, error) {

	if expirationSecs <= 0 {
		return "", fmt.Errorf("getSignedURL: seconds to expiration must be > 0")
	}

	signedURL, urlErr := storage.SignedURL(bucketName, fileName, &storage.SignedURLOptions{
		GoogleAccessID: authConfig.Email,
		PrivateKey:     authConfig.PrivateKey,
		Method:         "GET",
		Expires:        time.Now().Add(time.Second * time.Duration(expirationSecs))})

	if urlErr != nil {
		return "", urlErr
	}
	return signedURL, nil
}

func uploadFile(req *http.Request) (*UploadFileResponse, error) {

	recordID := req.FormValue("recordID")
	fieldID := req.FormValue("fieldID")
	log.Printf("Uploading file: record ID = %v, field id = %v", recordID, fieldID)

	formFile, handler, formErr := req.FormFile("uploadFile")
	if formErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to upload file: invalid api/form input in request: %v req = %+v", formErr, req)
	}
	log.Printf("Uploading file: %v", handler.Filename)

	cloudFileName := handler.Filename

	fileContents, readErr := ioutil.ReadAll(formFile)
	if formErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to read file contents: %v", readErr)
	}

	fileLength := len(fileContents)
	log.Printf("Uploading file: got file contents: length (bytes) = %v", fileLength)

	authConfig, configErr := readServiceAuthConfig(cloudStorageJSONAuthInfoFile)
	if configErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to get credentials for cloud storage: %v", configErr)
	}

	appEngContext := appengine.NewContext(req)

	// Save the file to cloud storage
	client, clientErr := getCloudClient(appEngContext, authConfig)
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

	// Generate an URL for the newly saved file
	signedURL, urlErr := getSignedURL(cloudStorageBucketName, cloudFileName, authConfig, 60)
	if urlErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to create signed URL for newly uploaded file: %v", urlErr)
	}

	uploadFile := UploadFile{
		Name:  cloudFileName,
		Size:  fileLength,
		Error: "",
		Url:   signedURL}
	uploadResponse := UploadFileResponse{Files: []UploadFile{uploadFile}}

	return &uploadResponse, nil
}
