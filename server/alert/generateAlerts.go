package alert

import (
	"fmt"
	"log"
	"resultra/datasheet/server/calcField"
	"resultra/datasheet/server/record"
	"sort"
)

type AlertProcessingResult struct {
	Error         error
	RecordID      string
	Notifications []AlertNotification
}

type RecordAlertProcessingConfig struct {
	RecordID        string
	CalcFieldConfig *calcField.CalcFieldUpdateConfig
	RecCellUpdates  *record.RecordCellUpdates
	Alerts          []Alert
}

// generateOneRecordAlertsFromConfig is the internal (lower level) implementation function for
// generating alerts for a single record.
func generateOneRecordAlertsFromConfig(recProcessingConfig RecordAlertProcessingConfig) ([]AlertNotification, error) {

	// Cell updates need to be in chronological order to be processed for alerts
	sort.Sort(record.CellUpdateByUpdateTime(recProcessingConfig.RecCellUpdates.CellUpdates))

	cellUpdateFieldValIndex, indexErr := record.NewUpdateFieldValueIndexForCellUpdates(recProcessingConfig.RecCellUpdates,
		recProcessingConfig.CalcFieldConfig.FieldsByID)
	if indexErr != nil {
		return nil, fmt.Errorf("MapOneRecordUpdatesToFieldValues: %v", indexErr)
	}

	// For non-calculated fields, get the latest (most recent) field values.
	prevFieldValues := record.RecFieldValues{}

	alertNotifications := []AlertNotification{}

	// Build up a set of field values for current time (latest). This is needed for the summary
	// item/record information which is presented alongside the alert notification. In particular,
	// each alert has an associated summary field which is selected to identify a record/item
	// in the context of an alert notification.
	latestFieldValues := cellUpdateFieldValIndex.LatestNonCalcFieldValues()
	if calcErr := calcField.UpdateCalcFieldValues(recProcessingConfig.CalcFieldConfig, latestFieldValues); calcErr != nil {
		return nil, fmt.Errorf("GenerateRecordAlerts: : err = %v", calcErr)
	}

	// Iterate through the record's cell updates. This provides a "tick by tick" iteration of the changes
	// for the given record. Then, each time there is a cell update, recalculate the calculated field
	// values for the record. Finally, process each of the alerts and generate an alert notification if needed.
	for _, currCellUpdate := range recProcessingConfig.RecCellUpdates.CellUpdates {

		currFieldValues := cellUpdateFieldValIndex.NonCalcFieldValuesAsOf(currCellUpdate.UpdateTimeStamp)

		// Populate calculated field values into currFieldValues at the time of currCellUpdates's timestamp.
		// This allows alerts which trigger from calculated field values to be processed just like
		// non calculated fields.
		if calcErr := calcField.UpdateCalcFieldValues(recProcessingConfig.CalcFieldConfig, &currFieldValues); calcErr != nil {
			return nil, fmt.Errorf("GenerateRecordAlerts: : err = %v", calcErr)
		}

		for _, currAlert := range recProcessingConfig.Alerts {

			itemSummaryFieldVal, foundSummary := latestFieldValues.GetTextFieldValue(currAlert.Properties.SummaryFieldID)
			itemSummary := ""
			if foundSummary {
				itemSummary = itemSummaryFieldVal
			}

			context := AlertProcessingContext{
				CalcFieldConfig: recProcessingConfig.CalcFieldConfig,
				RecordID:        recProcessingConfig.RecordID,
				UpdateTimestamp: currCellUpdate.UpdateTimeStamp,
				PrevFieldVals:   prevFieldValues,
				CurrFieldVals:   currFieldValues,
				LatestFieldVals: *latestFieldValues,
				ItemSummary:     itemSummary,
				ProcessedAlert:  currAlert}

			alertNotification, processAlertErr := processAlert(context)
			if processAlertErr != nil {
				return nil, fmt.Errorf("generateOneRecordAlertsFromConfig: : err = %v", processAlertErr)
			} else if alertNotification != nil {
				alertNotifications = append(alertNotifications, *alertNotification)
			}
		}
		prevFieldValues = currFieldValues
	}

	return alertNotifications, nil
}

func processOneRecordAlertsWorker(resultsChan chan AlertProcessingResult,
	recAlertProcessingConfig RecordAlertProcessingConfig) {

	alertNotifications, err := generateOneRecordAlertsFromConfig(recAlertProcessingConfig)

	result := AlertProcessingResult{
		Error:         err,
		RecordID:      recAlertProcessingConfig.RecordID,
		Notifications: alertNotifications}

	resultsChan <- result

}

type AlertGenerationResult struct {
	AlertsByID    map[string]Alert    `json:"alertsByID"`
	Notifications []AlertNotification `json:"notifications"`
}

// GenerateAllAlerts regenerates all alerts for all records. This is the top-level function to re-generate all the
// alert notifications at once.
func generateAllAlerts(databaseID string) (*AlertGenerationResult, error) {

	// Create a config/context object used for calculating the calculated fields for alert generation. This same
	// config can be reused by the alert generation for all the records.
	calcFieldUpdateConfig, err := calcField.CreateCalcFieldUpdateConfig(databaseID)
	if err != nil {
		return nil, fmt.Errorf("MapAllRecordUpdatesToFieldValues: %v", err)
	}

	// Get all cell updates at once for the given tracking databse. This is much faster than doing a database call for each and every record.
	recordCellUpdateMap, err := record.GetAllNonDraftCellUpdates(databaseID, record.FullyCommittedCellUpdatesChangeSetID)
	if err != nil {
		return nil, fmt.Errorf("MapAllRecordUpdatesToFieldValues: %v", err)
	}

	alerts, alertErr := getAllAlerts(databaseID)
	if alertErr != nil {
		return nil, fmt.Errorf("GenerateRecordAlerts: Error getting alerts: %v", alertErr)
	}
	alertsByID := map[string]Alert{}
	for _, currAlert := range alerts {
		alertsByID[currAlert.AlertID] = currAlert
	}

	alertNotifications := []AlertNotification{}
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
			return nil, fmt.Errorf("MapAllRecordUpdatesToFieldValues: %v", result.Error)
		} else if result.Notifications != nil {
			alertNotifications = append(alertNotifications, result.Notifications...)
		}
	}

	// Sort in reverse chronological order
	sort.Sort(NotificationByTime(alertNotifications))

	alertGenResults := AlertGenerationResult{
		AlertsByID:    alertsByID,
		Notifications: alertNotifications}

	return &alertGenResults, nil

}

// GenerateOneRecordAlerts is a top-level entry point for regenerating the alerts for an entire
// tracker, but a single recordID
func GenerateOneRecordAlerts(databaseID string, recordID string) ([]AlertNotification, error) {

	log.Printf("Regenerating alerts ...")

	calcFieldUpdateConfig, err := calcField.CreateCalcFieldUpdateConfig(databaseID)
	if err != nil {
		return nil, fmt.Errorf("MapAllRecordUpdatesToFieldValues: %v", err)
	}

	// Retrieve a list of sorted record updates.
	recCellUpdates, getErr := record.GetRecordCellUpdates(recordID, record.FullyCommittedCellUpdatesChangeSetID)
	if getErr != nil {
		return nil, fmt.Errorf("GenerateRecordAlerts: failure retrieving cell updates for record = %v: error = %v",
			recordID, getErr)
	}

	alerts, alertErr := getAllAlerts(databaseID)
	if alertErr != nil {
		return nil, fmt.Errorf("GenerateRecordAlerts: Error getting alerts: %v", alertErr)
	}

	alertProcessConfig := RecordAlertProcessingConfig{
		RecordID:        recordID,
		CalcFieldConfig: calcFieldUpdateConfig,
		RecCellUpdates:  recCellUpdates,
		Alerts:          alerts}

	return generateOneRecordAlertsFromConfig(alertProcessConfig)
}
