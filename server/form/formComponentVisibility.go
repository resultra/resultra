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
		for _, currTextBox := range formInfo.TextBoxes {
			if len(currTextBox.Properties.VisibilityConditions) > 0 {
				compFilterMap[currTextBox.TextBoxID] = currTextBox.Properties.VisibilityConditions
				log.Printf("Adding visibility filter to filter map for component ID = %v", currTextBox.TextBoxID)
			}
		}
		for _, currCheckBox := range formInfo.CheckBoxes {
			if len(currCheckBox.Properties.VisibilityConditions) > 0 {
				compFilterMap[currCheckBox.CheckBoxID] = currCheckBox.Properties.VisibilityConditions
				log.Printf("Adding visibility filter to filter map for component ID = %v", currCheckBox.CheckBoxID)
			}
		}
		for _, currDatePicker := range formInfo.DatePickers {
			if len(currDatePicker.Properties.VisibilityConditions) > 0 {
				compFilterMap[currDatePicker.DatePickerID] = currDatePicker.Properties.VisibilityConditions
				log.Printf("Adding visibility filter to filter map for component ID = %v", currDatePicker.DatePickerID)
			}
		}
		for _, currRating := range formInfo.Ratings {
			if len(currRating.Properties.VisibilityConditions) > 0 {
				compFilterMap[currRating.RatingID] = currRating.Properties.VisibilityConditions
				log.Printf("Adding visibility filter to filter map for component ID = %v", currRating.RatingID)
			}
		}
		for _, currProgress := range formInfo.ProgressIndicators {
			if len(currProgress.Properties.VisibilityConditions) > 0 {
				compFilterMap[currProgress.ProgressID] = currProgress.Properties.VisibilityConditions
				log.Printf("Adding visibility filter to filter map for component ID = %v", currProgress.ProgressID)
			}
		}
		for _, currButton := range formInfo.FormButtons {
			if len(currButton.Properties.VisibilityConditions) > 0 {
				compFilterMap[currButton.ButtonID] = currButton.Properties.VisibilityConditions
				log.Printf("Adding visibility filter to filter map for component ID = %v", currButton.ButtonID)
			}
		}
		for _, currHeader := range formInfo.Headers {
			if len(currHeader.Properties.VisibilityConditions) > 0 {
				compFilterMap[currHeader.HeaderID] = currHeader.Properties.VisibilityConditions
				log.Printf("Adding visibility filter to filter map for component ID = %v", currHeader.HeaderID)
			}
		}
	}

	return compFilterMap, nil
}
