package label

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/trackerDatabase"
)

const labelEntityKind string = "tags"

type Label struct {
	ParentFormID string          `json:"parentFormID"`
	LabelID      string          `json:"labelID"`
	Properties   LabelProperties `json:"properties"`
}

type NewLabelParams struct {
	ParentFormID string                         `json:"parentFormID"`
	FieldID      string                         `json:"fieldID"`
	Geometry     componentLayout.LayoutGeometry `json:"geometry"`
}

func validLabelFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeLabel {
		return true
	} else {
		return false
	}
}

func saveLabel(destDBHandle *sql.DB, newLabel Label) error {
	if saveErr := common.SaveNewFormComponent(destDBHandle, labelEntityKind,
		newLabel.ParentFormID, newLabel.LabelID, newLabel.Properties); saveErr != nil {
		return fmt.Errorf("saveNewLabel: Unable to save label: error = %v", saveErr)
	}
	return nil
}

func saveNewLabel(trackerDBHandle *sql.DB, params NewLabelParams) (*Label, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if fieldErr := field.ValidateField(trackerDBHandle, params.FieldID, validLabelFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewLabel: %v", fieldErr)
	}

	properties := newDefaultLabelProperties()
	properties.FieldID = params.FieldID
	properties.Geometry = params.Geometry

	newLabel := Label{ParentFormID: params.ParentFormID,
		LabelID:    uniqueID.GenerateUniqueID(),
		Properties: properties}

	if saveErr := saveLabel(trackerDBHandle, newLabel); saveErr != nil {
		return nil, fmt.Errorf("saveNewLabel: Unable to save label with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: New Label: Created new label component:  %+v", newLabel)

	return &newLabel, nil

}

func getLabel(trackerDBHandle *sql.DB, parentFormID string, labelID string) (*Label, error) {

	labelProps := newDefaultLabelProperties()
	if getErr := common.GetFormComponent(trackerDBHandle, labelEntityKind, parentFormID,
		labelID, &labelProps); getErr != nil {
		return nil, fmt.Errorf("getLabel: Unable to retrieve label: %v", getErr)
	}

	label := Label{
		ParentFormID: parentFormID,
		LabelID:      labelID,
		Properties:   labelProps}

	return &label, nil
}

func getLabelsFromSrc(srcDBHandle *sql.DB, parentFormID string) ([]Label, error) {

	labels := []Label{}
	addLabel := func(labelID string, encodedProps string) error {

		labelProps := newDefaultLabelProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &labelProps); decodeErr != nil {
			return fmt.Errorf("GetLabels: can't decode properties: %v", encodedProps)
		}

		currLabel := Label{
			ParentFormID: parentFormID,
			LabelID:      labelID,
			Properties:   labelProps}
		labels = append(labels, currLabel)

		return nil
	}
	if getErr := common.GetFormComponents(srcDBHandle, labelEntityKind, parentFormID, addLabel); getErr != nil {
		return nil, fmt.Errorf("GetLabels: Can't get labels: %v")
	}

	return labels, nil
}

func GetLabels(trackerDBHandle *sql.DB, parentFormID string) ([]Label, error) {
	return getLabelsFromSrc(trackerDBHandle, parentFormID)
}

func CloneLabels(cloneParams *trackerDatabase.CloneDatabaseParams, parentFormID string) error {

	srcLabels, err := getLabelsFromSrc(cloneParams.SrcDBHandle, parentFormID)
	if err != nil {
		return fmt.Errorf("CloneLabels: %v", err)
	}

	for _, srcLabel := range srcLabels {
		remappedLabelID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcLabel.LabelID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcLabel.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneLabels: %v", err)
		}
		destProperties, err := srcLabel.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneLabels: %v", err)
		}
		destLabel := Label{
			ParentFormID: remappedFormID,
			LabelID:      remappedLabelID,
			Properties:   *destProperties}
		if err := saveLabel(cloneParams.DestDBHandle, destLabel); err != nil {
			return fmt.Errorf("CloneLabels: %v", err)
		}
	}

	return nil
}

func updateExistingLabel(trackerDBHandle *sql.DB, updatedLabel *Label) (*Label, error) {

	if updateErr := common.UpdateFormComponent(trackerDBHandle, labelEntityKind, updatedLabel.ParentFormID,
		updatedLabel.LabelID, updatedLabel.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingLabel: failure updating label: %v", updateErr)
	}
	return updatedLabel, nil

}
