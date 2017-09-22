package label

import (
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
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

func saveLabel(newLabel Label) error {
	if saveErr := common.SaveNewFormComponent(labelEntityKind,
		newLabel.ParentFormID, newLabel.LabelID, newLabel.Properties); saveErr != nil {
		return fmt.Errorf("saveNewLabel: Unable to save label: error = %v", saveErr)
	}
	return nil
}

func saveNewLabel(params NewLabelParams) (*Label, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if fieldErr := field.ValidateField(params.FieldID, validLabelFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewLabel: %v", fieldErr)
	}

	properties := newDefaultLabelProperties()
	properties.FieldID = params.FieldID
	properties.Geometry = params.Geometry

	newLabel := Label{ParentFormID: params.ParentFormID,
		LabelID:    uniqueID.GenerateSnowflakeID(),
		Properties: properties}

	if saveErr := saveLabel(newLabel); saveErr != nil {
		return nil, fmt.Errorf("saveNewLabel: Unable to save label with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: New Label: Created new label component:  %+v", newLabel)

	return &newLabel, nil

}

func getLabel(parentFormID string, labelID string) (*Label, error) {

	labelProps := newDefaultLabelProperties()
	if getErr := common.GetFormComponent(labelEntityKind, parentFormID,
		labelID, &labelProps); getErr != nil {
		return nil, fmt.Errorf("getLabel: Unable to retrieve label: %v", getErr)
	}

	label := Label{
		ParentFormID: parentFormID,
		LabelID:      labelID,
		Properties:   labelProps}

	return &label, nil
}

func GetLabels(parentFormID string) ([]Label, error) {

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
	if getErr := common.GetFormComponents(labelEntityKind, parentFormID, addLabel); getErr != nil {
		return nil, fmt.Errorf("GetLabels: Can't get labels: %v")
	}

	return labels, nil
}

func CloneLabels(remappedIDs uniqueID.UniqueIDRemapper, parentFormID string) error {

	srcLabels, err := GetLabels(parentFormID)
	if err != nil {
		return fmt.Errorf("CloneLabels: %v", err)
	}

	for _, srcLabel := range srcLabels {
		remappedLabelID := remappedIDs.AllocNewOrGetExistingRemappedID(srcLabel.LabelID)
		remappedFormID, err := remappedIDs.GetExistingRemappedID(srcLabel.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneLabels: %v", err)
		}
		destProperties, err := srcLabel.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneLabels: %v", err)
		}
		destLabel := Label{
			ParentFormID: remappedFormID,
			LabelID:      remappedLabelID,
			Properties:   *destProperties}
		if err := saveLabel(destLabel); err != nil {
			return fmt.Errorf("CloneLabels: %v", err)
		}
	}

	return nil
}

func updateExistingLabel(updatedLabel *Label) (*Label, error) {

	if updateErr := common.UpdateFormComponent(labelEntityKind, updatedLabel.ParentFormID,
		updatedLabel.LabelID, updatedLabel.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingLabel: failure updating label: %v", updateErr)
	}
	return updatedLabel, nil

}
