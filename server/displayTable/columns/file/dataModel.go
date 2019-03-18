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
	"resultra/tracker/server/displayTable/columns/common"
	"resultra/tracker/server/field"
	"resultra/tracker/server/generic"
	"resultra/tracker/server/generic/uniqueID"
	"resultra/tracker/server/trackerDatabase"
)

const fileEntityKind string = "file"

type File struct {
	ParentTableID string         `json:"parentTableID"`
	FileID        string         `json:"fileID"`
	ColType       string         `json:"colType"`
	ColumnID      string         `json:"columnID"`
	Properties    FileProperties `json:"properties"`
}

type NewFileParams struct {
	ParentTableID string `json:"parentTableID"`
	FieldID       string `json:"fieldID"`
}

func validFileFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeFile {
		return true
	} else {
		return false
	}
}

func saveFile(destDBHandle *sql.DB, newFile File) error {
	if saveErr := common.SaveNewTableColumn(destDBHandle, fileEntityKind,
		newFile.ParentTableID, newFile.FileID, newFile.Properties); saveErr != nil {
		return fmt.Errorf("saveFile: Unable to save text box: %v", saveErr)
	}
	return nil

}

func saveNewFile(trackerDBHandle *sql.DB, params NewFileParams) (*File, error) {

	if fieldErr := field.ValidateField(trackerDBHandle, params.FieldID, validFileFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewFile: %v", fieldErr)
	}

	properties := newDefaultFileProperties()
	properties.FieldID = params.FieldID

	fileID := uniqueID.GenerateUniqueID()
	newFile := File{ParentTableID: params.ParentTableID,
		FileID:     fileID,
		ColumnID:   fileID,
		Properties: properties,
		ColType:    fileEntityKind}

	if err := saveFile(trackerDBHandle, newFile); err != nil {
		return nil, fmt.Errorf("saveNewFile: Unable to save text box with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: NewLayout: Created new Layout container: %+v", newFile)

	return &newFile, nil

}

func getFile(trackerDBHandle *sql.DB, parentTableID string, fileID string) (*File, error) {

	fileProps := newDefaultFileProperties()
	if getErr := common.GetTableColumn(trackerDBHandle, fileEntityKind, parentTableID, fileID, &fileProps); getErr != nil {
		return nil, fmt.Errorf("getCheckBox: Unable to retrieve text box: %v", getErr)
	}

	file := File{
		ParentTableID: parentTableID,
		FileID:        fileID,
		ColumnID:      fileID,
		Properties:    fileProps,
		ColType:       fileEntityKind}

	return &file, nil
}

func getFilesFromSrc(srcDBHandle *sql.DB, parentTableID string) ([]File, error) {

	files := []File{}
	addFile := func(fileID string, encodedProps string) error {

		fileProps := newDefaultFileProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &fileProps); decodeErr != nil {
			return fmt.Errorf("GetFiles: can't decode properties: %v", encodedProps)
		}

		currFile := File{
			ParentTableID: parentTableID,
			FileID:        fileID,
			ColumnID:      fileID,
			Properties:    fileProps,
			ColType:       fileEntityKind}
		files = append(files, currFile)

		return nil
	}
	if getErr := common.GetTableColumns(srcDBHandle, fileEntityKind, parentTableID, addFile); getErr != nil {
		return nil, fmt.Errorf("GetFiles: Can't get text boxes: %v")
	}

	return files, nil

}

func GetFiles(trackerDBHandle *sql.DB, parentTableID string) ([]File, error) {
	return getFilesFromSrc(trackerDBHandle, parentTableID)
}

func CloneFiles(cloneParams *trackerDatabase.CloneDatabaseParams, parentFormID string) error {

	srcFilees, err := getFilesFromSrc(cloneParams.SrcDBHandle, parentFormID)
	if err != nil {
		return fmt.Errorf("CloneFilees: %v", err)
	}

	for _, srcFile := range srcFilees {
		remappedFileID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcFile.FileID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcFile.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneFiles: %v", err)
		}
		destProperties, err := srcFile.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneFiles: %v", err)
		}
		destFile := File{
			ParentTableID: remappedFormID,
			FileID:        remappedFileID,
			ColumnID:      remappedFileID,
			Properties:    *destProperties,
			ColType:       fileEntityKind}
		if err := saveFile(cloneParams.DestDBHandle, destFile); err != nil {
			return fmt.Errorf("CloneFiles: %v", err)
		}
	}

	return nil
}

func updateExistingFile(trackerDBHandle *sql.DB, fileID string, updatedFile *File) (*File, error) {

	if updateErr := common.UpdateTableColumn(trackerDBHandle, fileEntityKind, updatedFile.ParentTableID,
		updatedFile.FileID, updatedFile.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingFile: error updating existing text box component: %v", updateErr)
	}

	return updatedFile, nil

}
