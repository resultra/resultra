package progress

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/datasheet/server/displayTable/columns/common"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/trackerDatabase"
)

const progressEntityKind string = "progress"

type Progress struct {
	ParentTableID string             `json:"parentTableID"`
	ProgressID    string             `json:"progressID"`
	ColumnID      string             `json:"columnID"`
	ColType       string             `json:"colType"`
	Properties    ProgressProperties `json:"properties"`
}

type NewProgressParams struct {
	ParentTableID string `json:"parentTableID"`
	FieldID       string `json:"fieldID"`
}

func validProgressFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeNumber {
		return true
	} else {
		return false
	}
}

func saveProgress(destDBHandle *sql.DB, newProgress Progress) error {
	if saveErr := common.SaveNewTableColumn(destDBHandle, progressEntityKind,
		newProgress.ParentTableID, newProgress.ProgressID, newProgress.Properties); saveErr != nil {
		return fmt.Errorf("saveProgress: Unable to save progress indicator with error = %v", saveErr)
	}
	return nil
}

func saveNewProgress(trackerDBHandle *sql.DB, params NewProgressParams) (*Progress, error) {

	if fieldErr := field.ValidateField(trackerDBHandle, params.FieldID, validProgressFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewProgress: %v", fieldErr)
	}

	properties := newDefaultProgressProperties()
	properties.FieldID = params.FieldID

	progressID := uniqueID.GenerateUniqueID()
	newProgress := Progress{ParentTableID: params.ParentTableID,
		ProgressID: progressID,
		ColumnID:   progressID,
		ColType:    progressEntityKind,
		Properties: properties}

	if err := saveProgress(trackerDBHandle, newProgress); err != nil {
		return nil, fmt.Errorf("saveNewProgress: Unable to save progress indicator with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: New progress indicator: Created progress indicator container:  %+v", newProgress)

	return &newProgress, nil

}

func getProgress(trackerDBHandle *sql.DB, parentTableID string, progressID string) (*Progress, error) {

	progressProps := newDefaultProgressProperties()
	if getErr := common.GetTableColumn(trackerDBHandle, progressEntityKind, parentTableID, progressID, &progressProps); getErr != nil {
		return nil, fmt.Errorf("getNumberInput: Unable to retrieve number input: %v", getErr)
	}

	progress := Progress{
		ParentTableID: parentTableID,
		ProgressID:    progressID,
		ColumnID:      progressID,
		ColType:       progressEntityKind,
		Properties:    progressProps}

	return &progress, nil
}

func getProgressIndicatorsFromSrc(srcDBHandle *sql.DB, parentTableID string) ([]Progress, error) {

	progressIndicators := []Progress{}
	addProgress := func(progressID string, encodedProps string) error {

		var progressProps = newDefaultProgressProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &progressProps); decodeErr != nil {
			return fmt.Errorf("GetProgressIndicators: can't decode properties: %v", encodedProps)
		}

		currProgress := Progress{
			ParentTableID: parentTableID,
			ProgressID:    progressID,
			ColumnID:      progressID,
			ColType:       progressEntityKind,
			Properties:    progressProps}
		progressIndicators = append(progressIndicators, currProgress)

		return nil
	}
	if getErr := common.GetTableColumns(srcDBHandle, progressEntityKind, parentTableID, addProgress); getErr != nil {
		return nil, fmt.Errorf("GetProgressIndicators: Can't get progress indicators: %v")
	}

	return progressIndicators, nil
}

func GetProgressIndicators(trackerDBHandle *sql.DB, parentTableID string) ([]Progress, error) {
	return getProgressIndicatorsFromSrc(trackerDBHandle, parentTableID)
}

func CloneProgressIndicators(cloneParams *trackerDatabase.CloneDatabaseParams, parentTableID string) error {

	srcProgressIndicators, err := getProgressIndicatorsFromSrc(cloneParams.SrcDBHandle, parentTableID)
	if err != nil {
		return fmt.Errorf("CloneProgressIndicators: %v", err)
	}

	for _, srcProgress := range srcProgressIndicators {
		remappedProgressID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcProgress.ProgressID)
		remappedTableID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcProgress.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneProgressIndicators: %v", err)
		}
		destProperties, err := srcProgress.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneProgressIndicators: %v", err)
		}
		destProgress := Progress{
			ParentTableID: remappedTableID,
			ProgressID:    remappedProgressID,
			ColumnID:      remappedProgressID,
			ColType:       progressEntityKind,
			Properties:    *destProperties}
		if err := saveProgress(cloneParams.DestDBHandle, destProgress); err != nil {
			return fmt.Errorf("CloneProgressIndicators: %v", err)
		}
	}

	return nil
}

func updateExistingProgress(trackerDBHandle *sql.DB, updatedProgress *Progress) (*Progress, error) {

	if updateErr := common.UpdateTableColumn(trackerDBHandle, progressEntityKind, updatedProgress.ParentTableID,
		updatedProgress.ProgressID, updatedProgress.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingProgress: failure updating progress indicator: %v", updateErr)
	}
	return updatedProgress, nil

}
