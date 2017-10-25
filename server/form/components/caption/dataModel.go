package caption

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/trackerDatabase"
)

const captionEntityKind string = "caption"

type Caption struct {
	ParentFormID string            `json:"parentFormID"`
	CaptionID    string            `json:"captionID"`
	Properties   CaptionProperties `json:"properties"`
}

type NewCaptionParams struct {
	ParentFormID string                         `json:"parentFormID"`
	Geometry     componentLayout.LayoutGeometry `json:"geometry"`
	Label        string                         `json:"label"`
}

func saveCaption(destDBHandle *sql.DB, newCaption Caption) error {

	if saveErr := common.SaveNewFormComponent(destDBHandle, captionEntityKind,
		newCaption.ParentFormID, newCaption.CaptionID, newCaption.Properties); saveErr != nil {
		return fmt.Errorf("saveCaption: Unable to save caption: error = %v", saveErr)
	}
	return nil

}

func saveNewCaption(params NewCaptionParams) (*Caption, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid form component layout parameters: %+v", params)
	}

	properties := CaptionProperties{
		Geometry:    params.Geometry,
		Label:       params.Label,
		ColorScheme: colorSchemeDefault}

	newCaption := Caption{ParentFormID: params.ParentFormID,
		CaptionID:  uniqueID.GenerateSnowflakeID(),
		Properties: properties}

	if err := saveCaption(databaseWrapper.DBHandle(), newCaption); err != nil {
		return nil, fmt.Errorf("saveNewCaption: Unable to save caption with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: New form caption: Created new form caption: %+v", newCaption)

	return &newCaption, nil

}

func getCaption(parentFormID string, captionID string) (*Caption, error) {

	captionProps := newDefaultCaptionProperties()
	if getErr := common.GetFormComponent(captionEntityKind, parentFormID, captionID, &captionProps); getErr != nil {
		return nil, fmt.Errorf("getCaption: Unable to retrieve caption: %v", getErr)
	}

	caption := Caption{
		ParentFormID: parentFormID,
		CaptionID:    captionID,
		Properties:   captionProps}

	return &caption, nil
}

func getCaptionsFromSrc(srcDBHandle *sql.DB, parentFormID string) ([]Caption, error) {

	captions := []Caption{}
	addCaption := func(datePickerID string, encodedProps string) error {

		captionProps := newDefaultCaptionProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &captionProps); decodeErr != nil {
			return fmt.Errorf("GetCaptions: can't decode properties: %v", encodedProps)
		}

		currCaption := Caption{
			ParentFormID: parentFormID,
			CaptionID:    datePickerID,
			Properties:   captionProps}
		captions = append(captions, currCaption)

		return nil
	}
	if getErr := common.GetFormComponents(srcDBHandle, captionEntityKind, parentFormID, addCaption); getErr != nil {
		return nil, fmt.Errorf("GetCaptions: Can't get captions: %v")
	}

	return captions, nil

}

func GetCaptions(parentFormID string) ([]Caption, error) {
	return getCaptionsFromSrc(databaseWrapper.DBHandle(), parentFormID)
}

func CloneCaptions(cloneParams *trackerDatabase.CloneDatabaseParams, parentFormID string) error {

	srcCaptions, err := getCaptionsFromSrc(cloneParams.SrcDBHandle, parentFormID)
	if err != nil {
		return fmt.Errorf("CloneCaptions: %v", err)
	}

	for _, srcCaption := range srcCaptions {
		remappedCaptionID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcCaption.CaptionID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcCaption.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneCaptions: %v", err)
		}
		destProperties, err := srcCaption.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneCaptions: %v", err)
		}
		destCaption := Caption{
			ParentFormID: remappedFormID,
			CaptionID:    remappedCaptionID,
			Properties:   *destProperties}
		if err := saveCaption(cloneParams.DestDBHandle, destCaption); err != nil {
			return fmt.Errorf("CloneCaptions: %v", err)
		}
	}

	return nil
}

func updateExistingCaption(updatedCaption *Caption) (*Caption, error) {

	if updateErr := common.UpdateFormComponent(captionEntityKind, updatedCaption.ParentFormID,
		updatedCaption.CaptionID, updatedCaption.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingCaption: failure updating caption: %v", updateErr)
	}
	return updatedCaption, nil

}
