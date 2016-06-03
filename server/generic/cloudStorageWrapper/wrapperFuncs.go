package cloudStorageWrapper

import (
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"

	// There are 2 different "app engine" packages: One used for cloud storage and another
	// used for processing Google datastore requests. TODO - Clarify and address this discrepency
	// further.
	cloudStorageAppEngine "google.golang.org/appengine"

	"google.golang.org/cloud"
	"google.golang.org/cloud/storage"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"resultra/datasheet/server/generic/uniqueID"
	"time"
)

func NewCloudStorageContext(req *http.Request) context.Context {
	return cloudStorageAppEngine.NewContext(req)
}

func ReadServiceAuthConfig(jsonKeyFileName string) (*jwt.Config, error) {
	jsonKey, err := ioutil.ReadFile(jsonKeyFileName)
	if err != nil {
		return nil, fmt.Errorf("getCloudContext: unable to read json auth info: %v", err)
	}
	conf, err := google.JWTConfigFromJSON(jsonKey, storage.ScopeFullControl)
	if err != nil {
		return nil, fmt.Errorf("getCloudContext: unable to create config from json auth info: %v", err)
	}
	return conf, nil
}

func GetCloudClient(appEngContext context.Context, config *jwt.Config) (*storage.Client, error) {

	client, clientErr := storage.NewClient(appEngContext, cloud.WithTokenSource(config.TokenSource(appEngContext)))
	if clientErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to save to cloud storage: %v", clientErr)
	}

	return client, nil
}

func GetSignedURL(bucketName string, fileName string, authConfig *jwt.Config, expirationSecs int) (string, error) {

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

// cloudFileNameFromUserFileName takes a file name and generates a unique
// file name for storage in the cloud. This filename is prefixed with a time-stamp
// so there is a human-readable portion which also makes the file name sortable. The
// second component of the filename is a unique "version 4" uuid, based upon
// random numbers. The uuid (along with the timestamp) ensures the filename
// is unique versus other files stored in the same cloud bucket.
func UniqueCloudFileNameFromUserFileName(userFileName string) string {

	uniqueIDStr := uniqueID.GenerateUniqueID()

	fileExt := path.Ext(userFileName)

	cloudFileName := uniqueIDStr + fileExt

	return cloudFileName
}

func SaveCloudFile(cloudContext context.Context, authConfig *jwt.Config, bucketName string, fileName string,
	fileData []byte) error {
	// Save the file to cloud storage
	client, clientErr := GetCloudClient(cloudContext, authConfig)
	if clientErr != nil {
		return fmt.Errorf("uploadFile: Unable to save to cloud storage: %v", clientErr)
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)
	log.Printf("uploadFile: start uploading file to cloud: bucket=%v, file name = %v ...", bucketName, fileName)

	cloudWriter := bucket.Object(fileName).NewWriter(cloudContext)

	fileLength := len(fileData)

	bytesWritten, writeErr := cloudWriter.Write(fileData)
	if writeErr != nil {
		return fmt.Errorf("SaveCloudFile: Unable to save to cloud storage: %v", writeErr)
	} else if bytesWritten != fileLength {
		return fmt.Errorf("SaveCloudFile: Unable to save to cloud storage: only %v of %v bytes written", bytesWritten, fileLength)
	}

	// TODO (Important) - Also set some metadata on the cloud object to identify the owner of the file. Otherwise,
	// someone could potentiall spoof the system into downloading other users' files (even though the uuid v4 is
	// generally unguessable). Or, another way to consider authentication of file attachments is to cross-check the
	// file ID within the record itself with the one being attempted to download (or obatin an URL); since only
	// someone authenticated to change the record will could have changed the file name/ID, by cross-checking the
	// file name to the one actually stored in the record should suffice to prevent unwanted downloads.

	// The file isn't actually written until Close() is called.
	if closeErr := cloudWriter.Close(); closeErr != nil {
		return fmt.Errorf("SaveCloudFile: Unable to save: failure closing cloud writer: %v", closeErr)
	}
	log.Printf("uploadFile: ... done uloading file to cloud: bucket=%v, file name = %v", bucketName, fileName)

	return nil

}
