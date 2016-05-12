package record

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/file"
	"google.golang.org/cloud/storage"
	"io/ioutil"
	"log"
	"net/http"
)

type UploadFileResponse struct {
	UploadSuccess bool `json:"status"`
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

	client, clientErr := storage.NewClient(appEngContext)
	if clientErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to save to cloud storage: %v", clientErr)
	}
	defer client.Close()

	//	bucketName := "test-attachments"
	bucketName, bucketNameErr := file.DefaultBucketName(appEngContext)
	if bucketNameErr != nil {
		return nil, fmt.Errorf("uploadFile: failed to get default GCS bucket name: %v", bucketNameErr)
	}
	bucket := client.Bucket(bucketName)
	log.Printf("Uploading file: uploading to bucket: name = %v", bucketName)

	cloudWriter := bucket.Object(cloudFileName).NewWriter(appEngContext)
	cloudWriter.ContentType = "text/plain"

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

	return &UploadFileResponse{true}, nil
}
