package form

import (
	"fmt"
	"log"
	"resultra/datasheet/server/recordFilter"
)

type FormComponentFilterMap map[string]recordFilter.RecordFilterRuleSet

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
			if !currCaption.Properties.VisibilityConditions.IsEmptyRuleSet() {
				compFilterMap[currCaption.CaptionID] = currCaption.Properties.VisibilityConditions
				log.Printf("Adding visibility filter to filter map for component ID = %v", currCaption.CaptionID)
			}
		}
		for _, currEditor := range formInfo.HtmlEditors {
			if !currEditor.Properties.VisibilityConditions.IsEmptyRuleSet() {
				compFilterMap[currEditor.HtmlEditorID] = currEditor.Properties.VisibilityConditions
				log.Printf("Adding visibility filter to filter map for component ID = %v", currEditor.HtmlEditorID)
			}
		}
		for _, currTextBox := range formInfo.TextBoxes {
			if !currTextBox.Properties.VisibilityConditions.IsEmptyRuleSet() {
				compFilterMap[currTextBox.TextBoxID] = currTextBox.Properties.VisibilityConditions
				log.Printf("Adding visibility filter to filter map for component ID = %v", currTextBox.TextBoxID)
			}
		}
		for _, currCheckBox := range formInfo.CheckBoxes {
			if !currCheckBox.Properties.VisibilityConditions.IsEmptyRuleSet() {
				compFilterMap[currCheckBox.CheckBoxID] = currCheckBox.Properties.VisibilityConditions
				log.Printf("Adding visibility filter to filter map for component ID = %v", currCheckBox.CheckBoxID)
			}
		}
		for _, currDatePicker := range formInfo.DatePickers {
			if !currDatePicker.Properties.VisibilityConditions.IsEmptyRuleSet() {
				compFilterMap[currDatePicker.DatePickerID] = currDatePicker.Properties.VisibilityConditions
				log.Printf("Adding visibility filter to filter map for component ID = %v", currDatePicker.DatePickerID)
			}
		}
		for _, currRating := range formInfo.Ratings {
			if !currRating.Properties.VisibilityConditions.IsEmptyRuleSet() {
				compFilterMap[currRating.RatingID] = currRating.Properties.VisibilityConditions
				log.Printf("Adding visibility filter to filter map for component ID = %v", currRating.RatingID)
			}
		}
		for _, currProgress := range formInfo.ProgressIndicators {
			if !currProgress.Properties.VisibilityConditions.IsEmptyRuleSet() {
				compFilterMap[currProgress.ProgressID] = currProgress.Properties.VisibilityConditions
				log.Printf("Adding visibility filter to filter map for component ID = %v", currProgress.ProgressID)
			}
		}
		for _, currButton := range formInfo.FormButtons {
			if !currButton.Properties.VisibilityConditions.IsEmptyRuleSet() {
				compFilterMap[currButton.ButtonID] = currButton.Properties.VisibilityConditions
				log.Printf("Adding visibility filter to filter map for component ID = %v", currButton.ButtonID)
			}
		}
		for _, currHeader := range formInfo.Headers {
			if !currHeader.Properties.VisibilityConditions.IsEmptyRuleSet() {
				compFilterMap[currHeader.HeaderID] = currHeader.Properties.VisibilityConditions
				log.Printf("Adding visibility filter to filter map for component ID = %v", currHeader.HeaderID)
			}
		}
		for _, currSelection := range formInfo.Selections {
			if !currSelection.Properties.VisibilityConditions.IsEmptyRuleSet() {
				compFilterMap[currSelection.SelectionID] = currSelection.Properties.VisibilityConditions
				log.Printf("Adding visibility filter to filter map for component ID = %v", currSelection.SelectionID)
			}
		}
		for _, currUserSelection := range formInfo.UserSelections {
			if !currUserSelection.Properties.VisibilityConditions.IsEmptyRuleSet() {
				compFilterMap[currUserSelection.UserSelectionID] = currUserSelection.Properties.VisibilityConditions
				log.Printf("Adding visibility filter to filter map for component ID = %v", currUserSelection.UserSelectionID)
			}
		}
		for _, currComment := range formInfo.Comments {
			if !currComment.Properties.VisibilityConditions.IsEmptyRuleSet() {
				compFilterMap[currComment.CommentID] = currComment.Properties.VisibilityConditions
				log.Printf("Adding visibility filter to filter map for component ID = %v", currComment.CommentID)
			}
		}
	}

	return compFilterMap, nil
}
