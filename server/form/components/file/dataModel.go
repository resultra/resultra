package file

import (
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
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

func saveFile(newFile File) error {
	if saveErr := common.SaveNewFormComponent(fileEntityKind,
		newFile.ParentFormID, newFile.FileID, newFile.Properties); saveErr != nil {
		return fmt.Errorf("saveFile: Unable to save text box: %v", saveErr)
	}
	return nil

}

func saveNewFile(params NewFileParams) (*File, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if fieldErr := field.ValidateField(params.FieldID, validFileFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewFile: %v", fieldErr)
	}

	properties := newDefaultFileProperties()
	properties.Geometry = params.Geometry
	properties.FieldID = params.FieldID

	newFile := File{ParentFormID: params.ParentFormID,
		FileID:     uniqueID.GenerateSnowflakeID(),
		Properties: properties}

	if err := saveFile(newFile); err != nil {
		return nil, fmt.Errorf("saveNewFile: Unable to save text box with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: NewLayout: Created new Layout container: %+v", newFile)

	return &newFile, nil

}

func getFile(parentFormID string, fileID string) (*File, error) {

	fileProps := newDefaultFileProperties()
	if getErr := common.GetFormComponent(fileEntityKind, parentFormID, fileID, &fileProps); getErr != nil {
		return nil, fmt.Errorf("getCheckBox: Unable to retrieve text box: %v", getErr)
	}

	file := File{
		ParentFormID: parentFormID,
		FileID:       fileID,
		Properties:   fileProps}

	return &file, nil
}

func GetFiles(parentFormID string) ([]File, error) {

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
	if getErr := common.GetFormComponents(fileEntityKind, parentFormID, addFile); getErr != nil {
		return nil, fmt.Errorf("GetCheckBoxes: Can't get text boxes: %v")
	}

	return filees, nil

}

func CloneFiles(remappedIDs uniqueID.UniqueIDRemapper, parentFormID string) error {

	srcFile, err := GetFiles(parentFormID)
	if err != nil {
		return fmt.Errorf("CloneFile: %v", err)
	}

	for _, srcFile := range srcFile {
		remappedFileID := remappedIDs.AllocNewOrGetExistingRemappedID(srcFile.FileID)
		remappedFormID, err := remappedIDs.GetExistingRemappedID(srcFile.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneFile: %v", err)
		}
		destProperties, err := srcFile.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneFile: %v", err)
		}
		destFile := File{
			ParentFormID: remappedFormID,
			FileID:       remappedFileID,
			Properties:   *destProperties}
		if err := saveFile(destFile); err != nil {
			return fmt.Errorf("CloneFile: %v", err)
		}
	}

	return nil
}

func updateExistingFile(fileID string, updatedFile *File) (*File, error) {

	if updateErr := common.UpdateFormComponent(fileEntityKind, updatedFile.ParentFormID,
		updatedFile.FileID, updatedFile.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingFile: error updating existing text box component: %v", updateErr)
	}

	return updatedFile, nil

}
