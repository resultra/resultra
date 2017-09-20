package recordValueMappingController

import (
	"fmt"
	"log"
	"resultra/datasheet/server/calcField"
	"resultra/datasheet/server/form"
	"resultra/datasheet/server/record"
	"resultra/datasheet/server/recordFilter"
	"resultra/datasheet/server/recordValue"
	"time"
)

func calculateHiddenFormComponents(currUserID string, parentDatabaseID string, componentFilterCondMap form.FormComponentFilterMap,
	recordVals record.RecFieldValues) ([]string, error) {

	hiddenComponents := []string{}

	for componentID, filterConds := range componentFilterCondMap {
		filterContext, contextErr := recordFilter.CreateFilterRuleContexts(currUserID, filterConds.FilterRules)
		if contextErr != nil {
			return nil, fmt.Errorf("CalculateHiddenFormComponents: %v", contextErr)
		}
		filterConditionsIndicateShowComponent, matchErr :=
			recordFilter.MatchOneRecordFromFieldValues(filterConds.MatchLogic, filterContext, recordVals)
		log.Printf("Calculating visibility filters for component = %v, filter conditions = %+v, filter result = %v",
			componentID, filterConds, filterConditionsIndicateShowComponent)
		if matchErr != nil {
			return nil, fmt.Errorf("CalculateHiddenFormComponents: %v", matchErr)
		}
		if !filterConditionsIndicateShowComponent {
			log.Printf("Calculating visibility filters: hiding component = %v", componentID)
			hiddenComponents = append(hiddenComponents, componentID)
		} else {
			log.Printf("Calculating visibility filters: showing component = %v", componentID)
		}
	}

	return hiddenComponents, nil
}

func mapOneRecordUpdatesWithCalcFieldConfig(config *calcField.CalcFieldUpdateConfig,
	componentFilterCondMap form.FormComponentFilterMap,
	mappedRecord record.Record,
	recCellUpdates *record.RecordCellUpdates, changeSetID string) (*recordValue.RecordValueResults, error) {

	cellUpdateFieldValIndex, indexErr := record.NewUpdateFieldValueIndexForCellUpdates(recCellUpdates, config.FieldsByID)
	if indexErr != nil {
		return nil, fmt.Errorf("MapOneRecordUpdatesToFieldValues: %v", indexErr)
	}

	// For non-calculated fields, get the latest (most recent) field values.
	latestFieldValues := cellUpdateFieldValIndex.LatestNonCalcFieldValues()

	// Now that all the non-calculated fields have been populated into latestFieldValues, all the calculated
	// fields also need to be populated. The formulas for calculated field by refer to the latest value of non-calculated
	// fields, so this set of values needs to be passed into UpdateCalcFieldValues as a starting point.
	if calcErr := calcField.UpdateCalcFieldValues(config, mappedRecord, latestFieldValues); calcErr != nil {
		return nil, fmt.Errorf("MapOneRecordUpdatesToFieldValues: Can't set value: Error calculating fields to reflect update: err = %v", calcErr)
	}

	hiddenComponents, hiddenCalcErr := calculateHiddenFormComponents(config.CurrUserID, config.ParentDatabaseID,
		componentFilterCondMap, *latestFieldValues)
	if hiddenCalcErr != nil {
		return nil, fmt.Errorf("MapOneRecordUpdatesToFieldValues: %v", hiddenCalcErr)
	}

	recValResults := recordValue.RecordValueResults{
		ParentDatabaseID:     config.ParentDatabaseID,
		RecordID:             recCellUpdates.RecordID,
		FieldValues:          *latestFieldValues,
		HiddenFormComponents: hiddenComponents}

	return &recValResults, nil

}

// Re-map the series of value updates to "flattened" current (most recent) values for both calculated
// and non-calculated fields.
func MapOneRecordUpdatesToFieldValues(currUserID string, parentDatabaseID string, recCellUpdates *record.RecordCellUpdates,
	changeSetID string) (*recordValue.RecordValueResults, error) {

	updateConfig, err := calcField.CreateCalcFieldUpdateConfig(currUserID, parentDatabaseID)
	if err != nil {
		return nil, fmt.Errorf("MapOneRecordUpdatesToFieldValues: %v", err)
	}
	componentFilterCondMap, err := form.GetDatabaseFormComponentFilterMap(parentDatabaseID)
	if err != nil {
		return nil, fmt.Errorf("MapOneRecordUpdatesToFieldValues: %v", err)
	}

	currRecord, err := record.GetRecord(recCellUpdates.RecordID)
	if err != nil {
		return nil, fmt.Errorf("MapOneRecordUpdatesToFieldValues: %v", err)
	}

	return mapOneRecordUpdatesWithCalcFieldConfig(updateConfig, componentFilterCondMap, *currRecord, recCellUpdates, changeSetID)
}

type RecordMappingResult struct {
	Error         error
	RecValResults *recordValue.RecordValueResults
}

func mapOneRecordWorker(resultsChan chan RecordMappingResult,
	config *calcField.CalcFieldUpdateConfig, componentFilterCondMap form.FormComponentFilterMap,
	mappedRecord record.Record,
	recCellUpdates *record.RecordCellUpdates) {

	recValResults, err := mapOneRecordUpdatesWithCalcFieldConfig(config, componentFilterCondMap,
		mappedRecord, recCellUpdates, record.FullyCommittedCellUpdatesChangeSetID)

	result := RecordMappingResult{
		Error:         err,
		RecValResults: recValResults}

	resultsChan <- result

}

func MapAllRecordUpdatesToFieldValues(currUserID string, parentDatabaseID string) ([]recordValue.RecordValueResults, error) {

	start := time.Now()

	updateConfig, err := calcField.CreateCalcFieldUpdateConfig(currUserID, parentDatabaseID)
	if err != nil {
		return nil, fmt.Errorf("MapAllRecordUpdatesToFieldValues: %v", err)
	}
	componentFilterCondMap, err := form.GetDatabaseFormComponentFilterMap(parentDatabaseID)
	if err != nil {
		return nil, fmt.Errorf("MapAllRecordUpdatesToFieldValues: %v", err)
	}

	recordCellUpdateMap, err := record.GetAllNonDraftCellUpdates(parentDatabaseID, record.FullyCommittedCellUpdatesChangeSetID)
	if err != nil {
		return nil, fmt.Errorf("MapAllRecordUpdatesToFieldValues: %v", err)
	}

	recordIDRecordMap, err := record.GetNonDraftRecordIDRecordMap(parentDatabaseID)
	if err != nil {
		return nil, fmt.Errorf("MapAllRecordUpdatesToFieldValues: %v", err)
	}

	resultsChan := make(chan RecordMappingResult)

	// Scatter: Map the results in goroutines
	for _, currRecCellUpdates := range recordCellUpdateMap {
		currRec, recFound := recordIDRecordMap[currRecCellUpdates.RecordID]
		if !recFound {
			return nil, fmt.Errorf("MapAllRecordUpdatesToFieldValues: %v", err)
		}
		go mapOneRecordWorker(resultsChan, updateConfig, componentFilterCondMap, currRec, currRecCellUpdates)
	}

	recValResults := []recordValue.RecordValueResults{}

	// Gather the results
	for range recordCellUpdateMap {
		result := <-resultsChan
		if result.Error != nil {
			return nil, fmt.Errorf("MapAllRecordUpdatesToFieldValues: %v", result.Error)
		}
		recValResults = append(recValResults, *result.RecValResults)
	}

	elapsed := time.Since(start)
	log.Printf("MapAllRecordUpdatesToFieldValues: elapsed time for %v records =  %s", len(recordCellUpdateMap), elapsed)

	return recValResults, nil

}
