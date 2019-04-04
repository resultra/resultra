// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package global

import (
	"fmt"
	"github.com/resultra/resultra/server/common/attachment"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/generic/api"
	"log"
	"net/http"
)

type UploadFile struct {
	Name            string          `json:"name"`
	Size            int             `json:"size"`
	Error           string          `json"error,omitempty"`
	Url             string          `json:"url"`
	GlobalValUpdate GlobalValUpdate `json:"globalValUpdate"`
}

type UploadFileResponse struct {
	Files []UploadFile `json:"files"`
}

func uploadFile(w http.ResponseWriter, req *http.Request) (*UploadFileResponse, error) {

	// The string "uploadFile" matches the parameter name used in clients.
	uploadInfo, uploadErr := api.ReadUploadFile(w, req, "uploadFile")
	if uploadErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to read file contents: %v", uploadErr)
	}

	// These calls to req.FormValue needs to come after the reading the uploaded file, since
	// req.FormValue will in turn call http.ParseMultipartForm. However, api.ReadUploadFile
	// needs to do error checking on the size of the requested upload through a call
	// to http.ParseMultipartForm.
	globalID := req.FormValue("globalID")
	parentDatabaseID := req.FormValue("parentDatabaseID")

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(req)
	if dbErr != nil {
		return nil, dbErr
	}

	global, getErr := getGlobal(trackerDBHandle, globalID)
	if getErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to retrieve global info for global ID = %v: %v",
			globalID, getErr)
	}
	if global.ParentDatabaseID != parentDatabaseID {
		return nil, fmt.Errorf("uploadFile: database and global info mismatch (database IDs don't match) info = %+v: database ID %v",
			global, parentDatabaseID)

	}

	cloudFileName := attachment.UniqueAttachmentFileNameFromUserFileName(uploadInfo.FileName)

	saveParams := databaseWrapper.SaveAttachmentParams{
		CloudFileName:    cloudFileName,
		ParentDatabaseID: global.ParentDatabaseID,
		HTTPReq:          req,
		FileData:         uploadInfo.FileData}

	if saveErr := databaseWrapper.SaveAttachment(saveParams); saveErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to save file to cloud storage: %v", saveErr)
	}

	// Generate an URL for the newly saved file
	fileURL := GetFileURL(trackerDBHandle, cloudFileName)

	// setRecordFileNameFieldValue. Although the parameters for a record update with a filename aren't passed through the http request,
	// the standard record updating mechanism can be used to update the field with the filename.
	updateGlobalHeader := GlobalValUpdateHeader{
		ParentDatabaseID: parentDatabaseID,
		GlobalID:         globalID}
	updateGlobalParams := SetImageGlobalValueParams{
		GlobalValUpdateHeader: updateGlobalHeader,
		OrigFileName:          uploadInfo.FileName,
		CloudFileName:         cloudFileName}

	valUpdate, updateErr := updateGlobalValue(trackerDBHandle, updateGlobalParams)
	if updateErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to update record for newly uploaded file: %v", updateErr)
	}

	log.Printf("uploadFile: Done uploading file: updated global = %+v", valUpdate)

	uploadFile := UploadFile{
		Name:            cloudFileName,
		Size:            uploadInfo.FileLength,
		Error:           "",
		Url:             fileURL,
		GlobalValUpdate: *valUpdate}
	uploadResponse := UploadFileResponse{Files: []UploadFile{uploadFile}}

	return &uploadResponse, nil
}
