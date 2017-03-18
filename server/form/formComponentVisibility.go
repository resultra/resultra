package form

import (
	"fmt"
	"log"
	"resultra/datasheet/server/recordFilter"
)

type RecordFilterRuleSet []recordFilter.RecordFilterRule

type FormComponentFilterMap map[string]RecordFilterRuleSet

// Build a map of component IDs to the filtering rules used to determine if the component is visible or not. This simplified
// structure is used to filter individual records to determine if the components are visible for each individual record.
func GetDatabaseFormComponentFilterMap(parentDatabaseID string) (FormComponentFilterMap, error) {

	forms, err := GetAllForms(parentDatabaseID)
	if err != nil {
		return nil, fmt.Errorf("GetDatabaseFormComponentFilterMap: Error getting forms for parent database ID = %v: %v",
			parentDatabaseID, err)
	}

	compFilterMap := FormComponentFilterMap{}

	for _, currForm := range forms {

		formInfo, formInfoErr := GetFormInfo(currForm.FormID)
		if formInfoErr != nil {
			return nil, fmt.Errorf("GetDatabaseFormComponentFilterMap: Error getting form info for form ID = %v: %v",
				currForm.FormID, err)
		}
		for _, currCaption := range formInfo.Captions {
			if len(currCaption.Properties.VisibilityConditions) > 0 {
				compFilterMap[currCaption.CaptionID] = currCaption.Properties.VisibilityConditions
				log.Printf("Adding visibility filter to filter map for component ID = %v", currCaption.CaptionID)
			}
		}
		for _, currEditor := range formInfo.HtmlEditors {
			if len(currEditor.Properties.VisibilityConditions) > 0 {
				compFilterMap[currEditor.HtmlEditorID] = currEditor.Properties.VisibilityConditions
				log.Printf("Adding visibility filter to filter map for component ID = %v", currEditor.HtmlEditorID)
			}
		}
	}

	return compFilterMap, nil
}
