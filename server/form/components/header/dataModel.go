package header

import (
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
)

const headerEntityKind string = "header"

type HeaderProperties struct {
	Label    string                         `json:"label"`
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

type Header struct {
	ParentFormID string           `json:"parentFormID"`
	HeaderID     string           `json:"headerID"`
	Properties   HeaderProperties `json:"properties"`
}

type NewHeaderParams struct {
	ParentFormID string                         `json:"parentFormID"`
	Geometry     componentLayout.LayoutGeometry `json:"geometry"`
	Label        string                         `json:"label"`
}

func saveNewHeader(params NewHeaderParams) (*Header, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid form component layout parameters: %+v", params)
	}

	properties := HeaderProperties{
		Geometry: params.Geometry,
		Label:    params.Label}

	newHeader := Header{ParentFormID: params.ParentFormID,
		HeaderID:   uniqueID.GenerateSnowflakeID(),
		Properties: properties}

	if saveErr := common.SaveNewFormComponent(headerEntityKind,
		newHeader.ParentFormID, newHeader.HeaderID, newHeader.Properties); saveErr != nil {
		return nil, fmt.Errorf("saveNewHeader: Unable to save header with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: New form header: Created new form header: %+v", newHeader)

	return &newHeader, nil

}

func GetHeaders(parentFormID string) ([]Header, error) {

	headers := []Header{}
	addHeader := func(datePickerID string, encodedProps string) error {

		var headerProps HeaderProperties
		if decodeErr := generic.DecodeJSONString(encodedProps, &headerProps); decodeErr != nil {
			return fmt.Errorf("GetHeaders: can't decode properties: %v", encodedProps)
		}

		currHeader := Header{
			ParentFormID: parentFormID,
			HeaderID:     datePickerID,
			Properties:   headerProps}
		headers = append(headers, currHeader)

		return nil
	}
	if getErr := common.GetFormComponents(headerEntityKind, parentFormID, addHeader); getErr != nil {
		return nil, fmt.Errorf("GetHeaders: Can't get headers: %v")
	}

	return headers, nil

}
