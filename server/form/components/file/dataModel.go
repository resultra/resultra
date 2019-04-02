// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package file

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/resultra/resultra/server/common/componentLayout"
	"github.com/resultra/resultra/server/field"
	"github.com/resultra/resultra/server/form/components/common"
	"github.com/resultra/resultra/server/generic"
	"github.com/resultra/resultra/server/generic/uniqueID"
	"github.com/resultra/resultra/server/trackerDatabase"
)

const fileEntityKind string = "file"

type File struct {
	ParentFormID string         `json:"parentFormID"`
	FileID       string         `json:"fileID"`
	Properties   FileProperties `json:"properties"`
}

type NewFileParams struct {
	ParentFormID string                         `json:"parentFormID"`
	FieldID      string                         `json:"fieldID"`
	Geometry     componentLayout.LayoutGeometry `json:"geometry"`
}

func validFileFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeFile {
		return true
	} else {
		return false
	}
}

func saveFile(destDBHandle *sql.DB, newFile File) error {
	if saveErr := common.SaveNewFormComponent(destDBHandle, fileEntityKind,
		newFile.ParentFormID, newFile.FileID, newFile.Properties); saveErr != nil {
		return fmt.Errorf("saveFile: Unable to save text box: %v", saveErr)
	}
	return nil

}

func saveNewFile(trackerDBHandle *sql.DB, params NewFileParams) (*File, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if fieldErr := field.ValidateField(trackerDBHandle, params.FieldID, validFileFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewFile: %v", fieldErr)
	}

	properties := newDefaultFileProperties()
	properties.Geometry = params.Geometry
	properties.FieldID = params.FieldID

	newFile := File{ParentFormID: params.ParentFormID,
		FileID:     uniqueID.GenerateUniqueID(),
		Properties: properties}

	if err := saveFile(trackerDBHandle, newFile); err != nil {
		return nil, fmt.Errorf("saveNewFile: Unable to save text box with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: NewLayout: Created new Layout container: %+v", newFile)

	return &newFile, nil

}

func getFile(trackerDBHandle *sql.DB, parentFormID string, fileID string) (*File, error) {

	fileProps := newDefaultFileProperties()
	if getErr := common.GetFormComponent(trackerDBHandle, fileEntityKind, parentFormID, fileID, &fileProps); getErr != nil {
		return nil, fmt.Errorf("getCheckBox: Unable to retrieve text box: %v", getErr)
	}

	file := File{
		ParentFormID: parentFormID,
		FileID:       fileID,
		Properties:   fileProps}

	return &file, nil
}

func getFilesFromSrc(srcDBHandle *sql.DB, parentFormID string) ([]File, error) {

	filees := []File{}
	addFile := func(fileID string, encodedProps string) error {

		fileProps := newDefaultFileProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &fileProps); decodeErr != nil {
			return fmt.Errorf("GetFile: can't decode properties: %v", encodedProps)
		}

		currFile := File{
			ParentFormID: parentFormID,
			FileID:       fileID,
			Properties:   fileProps}
		filees = append(filees, currFile)

		return nil
	}
	if getErr := common.GetFormComponents(srcDBHandle, fileEntityKind, parentFormID, addFile); getErr != nil {
		return nil, fmt.Errorf("GetCheckBoxes: Can't get text boxes: %v")
	}

	return filees, nil

}

func GetFiles(trackerDBHandle *sql.DB, parentFormID string) ([]File, error) {
	return getFilesFromSrc(trackerDBHandle, parentFormID)
}

func CloneFiles(cloneParams *trackerDatabase.CloneDatabaseParams, parentFormID string) error {

	srcFile, err := getFilesFromSrc(cloneParams.SrcDBHandle, parentFormID)
	if err != nil {
		return fmt.Errorf("CloneFile: %v", err)
	}

	for _, srcFile := range srcFile {
		remappedFileID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcFile.FileID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcFile.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneFile: %v", err)
		}
		destProperties, err := srcFile.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneFile: %v", err)
		}
		destFile := File{
			ParentFormID: remappedFormID,
			FileID:       remappedFileID,
			Properties:   *destProperties}
		if err := saveFile(cloneParams.DestDBHandle, destFile); err != nil {
			return fmt.Errorf("CloneFile: %v", err)
		}
	}

	return nil
}

func updateExistingFile(trackerDBHandle *sql.DB, fileID string, updatedFile *File) (*File, error) {

	if updateErr := common.UpdateFormComponent(trackerDBHandle, fileEntityKind, updatedFile.ParentFormID,
		updatedFile.FileID, updatedFile.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingFile: error updating existing text box component: %v", updateErr)
	}

	return updatedFile, nil

}
