package progress

import (
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
)

const progressEntityKind string = "progress"

type Progress struct {
	ParentFormID string             `json:"parentFormID"`
	ProgressID   string             `json:"progressID"`
	Properties   ProgressProperties `json:"properties"`
}

type NewProgressParams struct {
	ParentFormID string                         `json:"parentFormID"`
	FieldID      string                         `json:"fieldID"`
	Geometry     componentLayout.LayoutGeometry `json:"geometry"`
}

func validProgressFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeNumber {
		return true
	} else {
		return false
	}
}

func saveProgress(newProgress Progress) error {
	if saveErr := common.SaveNewFormComponent(progressEntityKind,
		newProgress.ParentFormID, newProgress.ProgressID, newProgress.Properties); saveErr != nil {
		return fmt.Errorf("saveProgress: Unable to save progress indicator with error = %v", saveErr)
	}
	return nil
}

func saveNewProgress(params NewProgressParams) (*Progress, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if fieldErr := field.ValidateField(params.FieldID, validProgressFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewProgress: %v", fieldErr)
	}

	properties := newDefaultProgressProperties()
	properties.FieldID = params.FieldID
	properties.Geometry = params.Geometry

	newProgress := Progress{ParentFormID: params.ParentFormID,
		ProgressID: uniqueID.GenerateSnowflakeID(),
		Properties: properties}

	if err := saveProgress(newProgress); err != nil {
		return nil, fmt.Errorf("saveNewProgress: Unable to save progress indicator with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: New progress indicator: Created progress indicator container:  %+v", newProgress)

	return &newProgress, nil

}

func getProgress(parentFormID string, progressID string) (*Progress, error) {

	progressProps := newDefaultProgressProperties()
	if getErr := common.GetFormComponent(progressEntityKind, parentFormID, progressID, &progressProps); getErr != nil {
		return nil, fmt.Errorf("getCheckBox: Unable to retrieve check box: %v", getErr)
	}

	progress := Progress{
		ParentFormID: parentFormID,
		ProgressID:   progressID,
		Properties:   progressProps}

	return &progress, nil
}

func GetProgressIndicators(parentFormID string) ([]Progress, error) {

	progressIndicators := []Progress{}
	addProgress := func(progressID string, encodedProps string) error {

		var progressProps = newDefaultProgressProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &progressProps); decodeErr != nil {
			return fmt.Errorf("GetProgressIndicators: can't decode properties: %v", encodedProps)
		}

		currProgress := Progress{
			ParentFormID: parentFormID,
			ProgressID:   progressID,
			Properties:   progressProps}
		progressIndicators = append(progressIndicators, currProgress)

		return nil
	}
	if getErr := common.GetFormComponents(progressEntityKind, parentFormID, addProgress); getErr != nil {
		return nil, fmt.Errorf("GetProgressIndicators: Can't get progress indicators: %v")
	}

	return progressIndicators, nil
}

func CloneProgressIndicators(remappedIDs uniqueID.UniqueIDRemapper, parentFormID string) error {

	srcProgressIndicators, err := GetProgressIndicators(parentFormID)
	if err != nil {
		return fmt.Errorf("CloneProgressIndicators: %v", err)
	}

	for _, srcProgress := range srcProgressIndicators {
		remappedProgressID := remappedIDs.AllocNewOrGetExistingRemappedID(srcProgress.ProgressID)
		remappedFormID, err := remappedIDs.GetExistingRemappedID(srcProgress.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneProgressIndicators: %v", err)
		}
		destProperties, err := srcProgress.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneProgressIndicators: %v", err)
		}
		destProgress := Progress{
			ParentFormID: remappedFormID,
			ProgressID:   remappedProgressID,
			Properties:   *destProperties}
		if err := saveProgress(destProgress); err != nil {
			return fmt.Errorf("CloneProgressIndicators: %v", err)
		}
	}

	return nil
}

func updateExistingProgress(updatedProgress *Progress) (*Progress, error) {

	if updateErr := common.UpdateFormComponent(progressEntityKind, updatedProgress.ParentFormID,
		updatedProgress.ProgressID, updatedProgress.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingProgress: failure updating progress indicator: %v", updateErr)
	}
	return updatedProgress, nil

}
