package alert

import (
	"fmt"
	"log"
	"resultra/datasheet/server/calcField"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/record"
	"sort"
)

type AlertProcessingContext struct {
	CalcFieldConfig *calcField.CalcFieldUpdateConfig
	PrevFieldVals   record.RecFieldValues
	CurrFieldVals   record.RecFieldValues
	ProcessedAlert  Alert
}

const alertCondChange string = "changed"

func processTimeFieldAlert(context AlertProcessingContext, cond AlertCondition) (bool, error) {

	log.Printf("Processing time field alert: %+v", cond)

	switch cond.ConditionID {
	case alertCondChange:
		valBefore, foundValBefore := context.PrevFieldVals.GetTimeFieldValue(cond.FieldID)
		valAfter, foundValAfter := context.CurrFieldVals.GetTimeFieldValue(cond.FieldID)
		if (foundValBefore == false) && (foundValAfter == false) {
			log.Printf("processTimeFieldAlert: no change - both vals undefined")
			return false, nil
		} else if (foundValBefore == false) && (foundValAfter == true) {
			log.Printf("processTimeFieldAlert: change - value defined: %v", valAfter)
			return true, nil
		} else if (foundValBefore == true) && (foundValAfter == false) {
			log.Printf("processTimeFieldAlert: change - time value cleared: cleared val = %v", valBefore)
			return true, nil
		} else { // value found bother before and after update => compare the actual dates
			if valBefore.Equal(valAfter) {
				log.Printf("processTimeFieldAlert: no change - values equal %v", valBefore)
				return false, nil // no change
			} else {
				log.Printf("processTimeFieldAlert: change - time values changed: %v -> %v", valBefore, valAfter)
				return true, nil
			}
		}
	default:
		return false, nil
	}

	return false, nil
}

func processAlert(context AlertProcessingContext) error {
	for _, currAlertCond := range context.ProcessedAlert.Properties.Conditions {

		condFieldID := currAlertCond.FieldID
		fieldInfo, fieldInfoFound := context.CalcFieldConfig.FieldsByID[condFieldID]
		if !fieldInfoFound {
			return fmt.Errorf("GenerateRecordAlerts: missing field info for field id = %v", condFieldID)
		}
		switch fieldInfo.Type {
		//			case field.FieldTypeText:
		//			case field.FieldTypeNumber:
		//			case field.FieldTypeBool:
		case field.FieldTypeTime:
			alertGenerated, genErr := processTimeFieldAlert(context, currAlertCond)
			if genErr != nil {
				return fmt.Errorf("processAlert:  %v", genErr)
			} else if alertGenerated == true {
				log.Printf("Alert generated: alert = %v, field = %v, condition = %v",
					context.ProcessedAlert.Name, fieldInfo.Name, currAlertCond.ConditionID)
				return nil // No need to process after matching the first condition
			}
		default:
			return fmt.Errorf("processAlert: Unsupported field type for alert: %v", fieldInfo.Type)
		} // switch

	}
	return nil
}

type AlertProcessingResult struct {
	Error    error
	RecordID string
}

type RecordAlertProcessingConfig struct {
	RecordID        string
	CalcFieldConfig *calcField.CalcFieldUpdateConfig
	RecCellUpdates  *record.RecordCellUpdates
	Alerts          []Alert
}

func generateOneRecordAlertsFromConfig(recProcessingConfig RecordAlertProcessingConfig) error {

	// Cell updates need to be in chronological order to be processed for alerts
	sort.Sort(record.CellUpdateByUpdateTime(recProcessingConfig.RecCellUpdates.CellUpdates))

	cellUpdateFieldValIndex, indexErr := record.NewUpdateFieldValueIndexForCellUpdates(recProcessingConfig.RecCellUpdates,
		recProcessingConfig.CalcFieldConfig.FieldsByID)
	if indexErr != nil {
		return fmt.Errorf("MapOneRecordUpdatesToFieldValues: %v", indexErr)
	}

	// For non-calculated fields, get the latest (most recent) field values.
	prevFieldValues := record.RecFieldValues{}

	// Iterate through the record's cell updates. This provides a "tick by tick" iteration of the changes
	// for the given record. Then, each time there is a cell update, recalculate the calculated field
	// values for the record. Finally, process each of the alerts and generate an alert notification if needed.
	for _, currCellUpdate := range recProcessingConfig.RecCellUpdates.CellUpdates {

		currFieldValues := cellUpdateFieldValIndex.NonCalcFieldValuesAsOf(currCellUpdate.UpdateTimeStamp)

		// Populate calculated field values into currFieldValues at the time of currCellUpdates's timestamp.
		// This allows alerts which trigger from calculated field values to be processed just like
		// non calculated fields.
		if calcErr := calcField.UpdateCalcFieldValues(recProcessingConfig.CalcFieldConfig, &currFieldValues); calcErr != nil {
			return fmt.Errorf("GenerateRecordAlerts: : err = %v", calcErr)
		}

		for _, currAlert := range recProcessingConfig.Alerts {

			context := AlertProcessingContext{
				CalcFieldConfig: recProcessingConfig.CalcFieldConfig,
				PrevFieldVals:   prevFieldValues,
				CurrFieldVals:   currFieldValues,
				ProcessedAlert:  currAlert}

			processAlert(context)
		}
		prevFieldValues = currFieldValues
	}

	return nil
}

func processOneRecordAlertsWorker(resultsChan chan AlertProcessingResult,
	recAlertProcessingConfig RecordAlertProcessingConfig) {

	err := generateOneRecordAlertsFromConfig(recAlertProcessingConfig)

	result := AlertProcessingResult{
		Error:    err,
		RecordID: recAlertProcessingConfig.RecordID}

	resultsChan <- result

}

func GenerateAllAlerts(databaseID string) error {

	// Create a config/context object used for calculating the calculated fields for alert generation. This same
	// config can be reused by the alert generation for all the records.
	calcFieldUpdateConfig, err := calcField.CreateCalcFieldUpdateConfig(databaseID)
	if err != nil {
		return fmt.Errorf("MapAllRecordUpdatesToFieldValues: %v", err)
	}

	// Get all cell updates at once for the given tracking databse. This is much faster than doing a database call for each and every record.
	recordCellUpdateMap, err := record.GetAllNonDraftCellUpdates(databaseID, record.FullyCommittedCellUpdatesChangeSetID)
	if err != nil {
		return fmt.Errorf("MapAllRecordUpdatesToFieldValues: %v", err)
	}

	alerts, alertErr := getAllAlerts(databaseID)
	if alertErr != nil {
		return fmt.Errorf("GenerateRecordAlerts: Error getting alerts: %v", alertErr)
	}

	resultsChan := make(chan AlertProcessingResult)

	for currRecordID, currRecCellUpdates := range recordCellUpdateMap {

		alertProcessConfig := RecordAlertProcessingConfig{
			RecordID:        currRecordID,
			CalcFieldConfig: calcFieldUpdateConfig,
			RecCellUpdates:  currRecCellUpdates,
			Alerts:          alerts}

		go processOneRecordAlertsWorker(resultsChan, alertProcessConfig)
	}

	// Gather the results
	for range recordCellUpdateMap {
		result := <-resultsChan
		if result.Error != nil {
			return fmt.Errorf("MapAllRecordUpdatesToFieldValues: %v", result.Error)
		}
	}

	return nil

}

func GenerateOneRecordAlerts(databaseID string, recordID string) error {

	log.Printf("Regenerating alerts ...")

	calcFieldUpdateConfig, err := calcField.CreateCalcFieldUpdateConfig(databaseID)
	if err != nil {
		return fmt.Errorf("MapAllRecordUpdatesToFieldValues: %v", err)
	}

	// Retrieve a list of sorted record updates.
	recCellUpdates, getErr := record.GetRecordCellUpdates(recordID, record.FullyCommittedCellUpdatesChangeSetID)
	if getErr != nil {
		return fmt.Errorf("GenerateRecordAlerts: failure retrieving cell updates for record = %v: error = %v",
			recordID, getErr)
	}

	alerts, alertErr := getAllAlerts(databaseID)
	if alertErr != nil {
		return fmt.Errorf("GenerateRecordAlerts: Error getting alerts: %v", alertErr)
	}

	alertProcessConfig := RecordAlertProcessingConfig{
		RecordID:        recordID,
		CalcFieldConfig: calcFieldUpdateConfig,
		RecCellUpdates:  recCellUpdates,
		Alerts:          alerts}

	return generateOneRecordAlertsFromConfig(alertProcessConfig)
}
