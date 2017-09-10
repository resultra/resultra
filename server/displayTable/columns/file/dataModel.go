package file

import (
	"fmt"
	"log"
	"resultra/datasheet/server/displayTable/columns/common"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
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

func saveFile(newFile File) error {
	if saveErr := common.SaveNewTableColumn(fileEntityKind,
		newFile.ParentTableID, newFile.FileID, newFile.Properties); saveErr != nil {
		return fmt.Errorf("saveFile: Unable to save text box: %v", saveErr)
	}
	return nil

}

func saveNewFile(params NewFileParams) (*File, error) {

	if fieldErr := field.ValidateField(params.FieldID, validFileFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewFile: %v", fieldErr)
	}

	properties := newDefaultFileProperties()
	properties.FieldID = params.FieldID

	fileID := uniqueID.GenerateSnowflakeID()
	newFile := File{ParentTableID: params.ParentTableID,
		FileID:     fileID,
		ColumnID:   fileID,
		Properties: properties,
		ColType:    fileEntityKind}

	if err := saveFile(newFile); err != nil {
		return nil, fmt.Errorf("saveNewFile: Unable to save text box with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: NewLayout: Created new Layout container: %+v", newFile)

	return &newFile, nil

}

func getFile(parentTableID string, fileID string) (*File, error) {

	fileProps := newDefaultFileProperties()
	if getErr := common.GetTableColumn(fileEntityKind, parentTableID, fileID, &fileProps); getErr != nil {
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

func GetFiles(parentTableID string) ([]File, error) {

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
	if getErr := common.GetTableColumns(fileEntityKind, parentTableID, addFile); getErr != nil {
		return nil, fmt.Errorf("GetFiles: Can't get text boxes: %v")
	}

	return files, nil

}

func CloneFiles(remappedIDs uniqueID.UniqueIDRemapper, parentFormID string) error {

	srcFilees, err := GetFiles(parentFormID)
	if err != nil {
		return fmt.Errorf("CloneFilees: %v", err)
	}

	for _, srcFile := range srcFilees {
		remappedFileID := remappedIDs.AllocNewOrGetExistingRemappedID(srcFile.FileID)
		remappedFormID, err := remappedIDs.GetExistingRemappedID(srcFile.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneFiles: %v", err)
		}
		destProperties, err := srcFile.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneFiles: %v", err)
		}
		destFile := File{
			ParentTableID: remappedFormID,
			FileID:        remappedFileID,
			ColumnID:      remappedFileID,
			Properties:    *destProperties,
			ColType:       fileEntityKind}
		if err := saveFile(destFile); err != nil {
			return fmt.Errorf("CloneFiles: %v", err)
		}
	}

	return nil
}

func updateExistingFile(fileID string, updatedFile *File) (*File, error) {

	if updateErr := common.UpdateTableColumn(fileEntityKind, updatedFile.ParentTableID,
		updatedFile.FileID, updatedFile.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingFile: error updating existing text box component: %v", updateErr)
	}

	return updatedFile, nil

}
