package cloudStorageWrapper

import (
	"fmt"

	"io/ioutil"
	"log"
	"path"
	"resultra/datasheet/server/generic/uniqueID"
)

const LocalAttachmentFileUploadDir string = `/private/tmp/`

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

func SaveCloudFile(fileName string, fileData []byte) error {

	writeErr := ioutil.WriteFile(LocalAttachmentFileUploadDir+fileName, fileData, 0644)
	if writeErr != nil {
		return fmt.Errorf("aveCloudFile: Error writing file: filename = %v, error = %v", fileName, writeErr)
	}

	log.Printf("uploadFile: ... done uloading file to cloud: file name = %v", fileName)

	return nil

}
