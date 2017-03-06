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

func saveHeader(newHeader Header) error {

	if saveErr := common.SaveNewFormComponent(headerEntityKind,
		newHeader.ParentFormID, newHeader.HeaderID, newHeader.Properties); saveErr != nil {
		return fmt.Errorf("saveHeader: Unable to save header: error = %v", saveErr)
	}
	return nil

}

func saveNewHeader(params NewHeaderParams) (*Header, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid form component layout parameters: %+v", params)
	}

	properties := HeaderProperties{
		Geometry:   params.Geometry,
		Label:      params.Label,
		HeaderSize: headerSizeMedium}

	newHeader := Header{ParentFormID: params.ParentFormID,
		HeaderID:   uniqueID.GenerateSnowflakeID(),
		Properties: properties}

	if err := saveHeader(newHeader); err != nil {
		return nil, fmt.Errorf("saveNewHeader: Unable to save header with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: New form header: Created new form header: %+v", newHeader)

	return &newHeader, nil

}

func getHeader(parentFormID string, headerID string) (*Header, error) {

	headerProps := newDefaultHeaderProperties()
	if getErr := common.GetFormComponent(headerEntityKind, parentFormID, headerID, &headerProps); getErr != nil {
		return nil, fmt.Errorf("getHeader: Unable to retrieve header: %v", getErr)
	}

	header := Header{
		ParentFormID: parentFormID,
		HeaderID:     headerID,
		Properties:   headerProps}

	return &header, nil
}

func GetHeaders(parentFormID string) ([]Header, error) {

	headers := []Header{}
	addHeader := func(datePickerID string, encodedProps string) error {

		headerProps := newDefaultHeaderProperties()
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

func CloneHeaders(remappedIDs uniqueID.UniqueIDRemapper, parentFormID string) error {

	srcHeaders, err := GetHeaders(parentFormID)
	if err != nil {
		return fmt.Errorf("CloneHeaders: %v", err)
	}

	for _, srcHeader := range srcHeaders {
		remappedHeaderID := remappedIDs.AllocNewOrGetExistingRemappedID(srcHeader.HeaderID)
		remappedFormID, err := remappedIDs.GetExistingRemappedID(srcHeader.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneHeaders: %v", err)
		}
		destProperties, err := srcHeader.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneHeaders: %v", err)
		}
		destHeader := Header{
			ParentFormID: remappedFormID,
			HeaderID:     remappedHeaderID,
			Properties:   *destProperties}
		if err := saveHeader(destHeader); err != nil {
			return fmt.Errorf("CloneHeaders: %v", err)
		}
	}

	return nil
}

func updateExistingHeader(updatedHeader *Header) (*Header, error) {

	if updateErr := common.UpdateFormComponent(headerEntityKind, updatedHeader.ParentFormID,
		updatedHeader.HeaderID, updatedHeader.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingHeader: failure updating header: %v", updateErr)
	}
	return updatedHeader, nil

}
