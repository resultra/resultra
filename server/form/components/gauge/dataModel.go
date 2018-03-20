package gauge

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

const gaugeEntityKind string = "gauge"

type Gauge struct {
	ParentFormID string          `json:"parentFormID"`
	GaugeID      string          `json:"gaugeID"`
	Properties   GaugeProperties `json:"properties"`
}

type NewGaugeParams struct {
	ParentFormID string                         `json:"parentFormID"`
	FieldID      string                         `json:"fieldID"`
	Geometry     componentLayout.LayoutGeometry `json:"geometry"`
}

func validGaugeFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeNumber {
		return true
	} else {
		return false
	}
}

func saveGauge(destDBHandle *sql.DB, newGauge Gauge) error {
	if saveErr := common.SaveNewFormComponent(destDBHandle, gaugeEntityKind,
		newGauge.ParentFormID, newGauge.GaugeID, newGauge.Properties); saveErr != nil {
		return fmt.Errorf("saveGauge: Unable to save gauge indicator with error = %v", saveErr)
	}
	return nil
}

func saveNewGauge(trackerDBHandle *sql.DB, params NewGaugeParams) (*Gauge, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if fieldErr := field.ValidateField(trackerDBHandle, params.FieldID, validGaugeFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewGauge: %v", fieldErr)
	}

	properties := newDefaultGaugeProperties()
	properties.FieldID = params.FieldID
	properties.Geometry = params.Geometry

	newGauge := Gauge{ParentFormID: params.ParentFormID,
		GaugeID:    uniqueID.GenerateUniqueID(),
		Properties: properties}

	if err := saveGauge(trackerDBHandle, newGauge); err != nil {
		return nil, fmt.Errorf("saveNewGauge: Unable to save gauge indicator with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: New gauge indicator: Created gauge indicator container:  %+v", newGauge)

	return &newGauge, nil

}

func getGauge(trackerDBHandle *sql.DB, parentFormID string, gaugeID string) (*Gauge, error) {

	gaugeProps := newDefaultGaugeProperties()
	if getErr := common.GetFormComponent(trackerDBHandle, gaugeEntityKind, parentFormID, gaugeID, &gaugeProps); getErr != nil {
		return nil, fmt.Errorf("getGauge: Unable to retrieve gauge: %v", getErr)
	}

	gauge := Gauge{
		ParentFormID: parentFormID,
		GaugeID:      gaugeID,
		Properties:   gaugeProps}

	return &gauge, nil
}

func getGaugesFromSrc(srcDBHandle *sql.DB, parentFormID string) ([]Gauge, error) {

	gaugeIndicators := []Gauge{}
	addGauge := func(gaugeID string, encodedProps string) error {

		var gaugeProps = newDefaultGaugeProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &gaugeProps); decodeErr != nil {
			return fmt.Errorf("GetGauges: can't decode properties: %v", encodedProps)
		}

		currGauge := Gauge{
			ParentFormID: parentFormID,
			GaugeID:      gaugeID,
			Properties:   gaugeProps}
		gaugeIndicators = append(gaugeIndicators, currGauge)

		return nil
	}
	if getErr := common.GetFormComponents(srcDBHandle, gaugeEntityKind, parentFormID, addGauge); getErr != nil {
		return nil, fmt.Errorf("GetGaugeIndicators: Can't get gauge indicators: %v")
	}

	return gaugeIndicators, nil
}

func GetGauges(trackerDBHandle *sql.DB, parentFormID string) ([]Gauge, error) {
	return getGaugesFromSrc(trackerDBHandle, parentFormID)
}

func CloneGauges(cloneParams *trackerDatabase.CloneDatabaseParams, parentFormID string) error {

	srcGaugeIndicators, err := GetGauges(cloneParams.SrcDBHandle, parentFormID)
	if err != nil {
		return fmt.Errorf("CloneGauges: %v", err)
	}

	for _, srcGauge := range srcGaugeIndicators {
		remappedGaugeID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcGauge.GaugeID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcGauge.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneGauges: %v", err)
		}
		destProperties, err := srcGauge.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneGauges: %v", err)
		}
		destGauge := Gauge{
			ParentFormID: remappedFormID,
			GaugeID:      remappedGaugeID,
			Properties:   *destProperties}
		if err := saveGauge(cloneParams.DestDBHandle, destGauge); err != nil {
			return fmt.Errorf("CloneGauges: %v", err)
		}
	}

	return nil
}

func updateExistingGauge(trackerDBHandle *sql.DB, updatedGauge *Gauge) (*Gauge, error) {

	if updateErr := common.UpdateFormComponent(trackerDBHandle, gaugeEntityKind, updatedGauge.ParentFormID,
		updatedGauge.GaugeID, updatedGauge.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingGauge: failure updating gauge indicator: %v", updateErr)
	}
	return updatedGauge, nil

}
