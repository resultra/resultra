package alert

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/datasheet/server/calcField"
	"resultra/datasheet/server/record"
	"resultra/datasheet/server/recordFilter"
	"resultra/datasheet/server/userRole"
	"sort"
	"time"
)

type AlertProcessingResult struct {
	Error         error
	RecordID      string
	Notifications []AlertNotification
}

type AlertGenerationContext struct {
	Alert                         Alert
	TriggerConditionFilterContext []recordFilter.FilterRuleContext
}

type RecordAlertProcessingConfig struct {
	TrackerDBHandle *sql.DB
	RecordID        string
	Record          record.Record
	CalcFieldConfig *calcField.CalcFieldUpdateConfig
	RecCellUpdates  *record.RecordCellUpdates
	AlertContexts   []AlertGenerationContext
}

// generateOneRecordAlertsFromConfig is the internal (lower level) implementation function for
// generating alerts for a single record.
func generateOneRecordAlertsFromConfig(recProcessingConfig RecordAlertProcessingConfig) ([]AlertNotification, error) {

	// Cell updates need to be in chronological order to be processed for alerts
	sort.Sort(record.CellUpdateByUpdateTime(recProcessingConfig.RecCellUpdates.CellUpdates))

	cellUpdateFieldValIndex, indexErr := record.NewUpdateFieldValueIndexForCellUpdates(
		recProcessingConfig.RecCellUpdates,
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
	if calcErr := calcField.UpdateCalcFieldValues(recProcessingConfig.CalcFieldConfig,
		recProcessingConfig.Record, latestFieldValues); calcErr != nil {
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
		if calcErr := calcField.UpdateCalcFieldValues(recProcessingConfig.CalcFieldConfig,
			recProcessingConfig.Record, &currFieldValues); calcErr != nil {
			return nil, fmt.Errorf("GenerateRecordAlerts: : err = %v", calcErr)
		}

		for _, currAlertContext := range recProcessingConfig.AlertContexts {

			currAlert := currAlertContext.Alert

			// Test for a match on the trigger conditions  record's field values "as of" the date of the
			// value update.
			recMatchesTriggerCond, condErr := recordFilter.MatchOneRecordFromFieldValues(
				currAlert.Properties.TriggerConditions.MatchLogic, currAlertContext.TriggerConditionFilterContext, currFieldValues)
			if condErr != nil {
				return nil, fmt.Errorf("GenerateRecordAlerts: : err = %v", condErr)
			}

			if recMatchesTriggerCond {
				context := AlertProcessingContext{
					CalcFieldConfig: recProcessingConfig.CalcFieldConfig,
					RecordID:        recProcessingConfig.RecordID,
					UpdateTimestamp: currCellUpdate.UpdateTimeStamp,
					PrevFieldVals:   prevFieldValues,
					CurrFieldVals:   currFieldValues,
					LatestFieldVals: *latestFieldValues,
					ProcessedAlert:  currAlert}

				alertNotification, processAlertErr := processAlert(context)
				if processAlertErr != nil {
					return nil, fmt.Errorf("generateOneRecordAlertsFromConfig: : err = %v", processAlertErr)
				} else if alertNotification != nil {
					alertNotifications = append(alertNotifications, *alertNotification)
				}

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
	AlertsByID       map[string]Alert    `json:"alertsByID"`
	Notifications    []AlertNotification `json:"notifications"`
	LatestAlertTime  time.Time           `json:"latestAlertTime"`
	NumAlertsNotSeen int                 `json:"numAlertsNotSeen"`
}

func getAlertsWithUserNotification(trackerDBHandle *sql.DB, databaseID string, userID string, userIsAdmin bool) ([]Alert, error) {

	allAlerts, alertErr := getAllAlerts(trackerDBHandle, databaseID)
	if alertErr != nil {
		return nil, fmt.Errorf("getAlertsWithUserNotification: Error getting alerts: %v", alertErr)
	}

	if userIsAdmin {
		return allAlerts, nil
	}

	userAlertByID, privsErr := userRole.GetAlertsWithUserPrivs(trackerDBHandle, databaseID, userID)
	if privsErr != nil {
		return nil, fmt.Errorf("getAlertsWithUserNotification: Error getting user alert notifications: %v", privsErr)
	}

	userAlerts := []Alert{}
	for _, currAlert := range allAlerts {
		_, userAlertFound := userAlertByID[currAlert.AlertID]
		if userAlertFound {
			userAlerts = append(userAlerts, currAlert)
		}
	}

	return userAlerts, nil

}

func createAlertGenerationContexts(trackerDBHandle *sql.DB, currUserID string, alerts []Alert) ([]AlertGenerationContext, error) {
	alertContexts := []AlertGenerationContext{}
	for _, currAlert := range alerts {
		triggerCondContext, condErr := recordFilter.CreateFilterRuleContexts(trackerDBHandle,
			currUserID, currAlert.Properties.TriggerConditions.FilterRules)
		if condErr != nil {
			return nil, fmt.Errorf("GenerateRecordAlerts: error setting up trigger condition filter contexts: %v", condErr)
		}
		alertContext := AlertGenerationContext{
			Alert: currAlert,
			TriggerConditionFilterContext: triggerCondContext}
		alertContexts = append(alertContexts, alertContext)
	}
	return alertContexts, nil
}

// If multiple alerts are generated for the same record under the same alert definition, only the latest alert notification is of interest.
// Otherwise, the user will be looking at "stale" information if they pull up the form linked to the alert definition. For example,
// an alert could be setup to trigger when a task is marked done.  The task could be marked done more than once, but only the latest
// is of interest to the user.
func pruneLatestUniqueAlertsByRecordAndAlert(unprunedNotifications []AlertNotification) []AlertNotification {

	pruningMap := map[string]AlertNotification{}
	for _, currNotification := range unprunedNotifications {

		alertRecKey := currNotification.AlertID + `-` + currNotification.RecordID

		existingNotification, foundNotification := pruningMap[alertRecKey]
		if foundNotification {
			if currNotification.Timestamp.After(existingNotification.Timestamp) {
				pruningMap[alertRecKey] = currNotification
			}
		} else {
			pruningMap[alertRecKey] = currNotification
		}

	}

	prunedNotifications := []AlertNotification{}
	for _, currNotification := range pruningMap {
		prunedNotifications = append(prunedNotifications, currNotification)
	}

	return prunedNotifications

}

// GenerateAllAlerts regenerates all alerts for all records. This is the top-level function to re-generate all the
// alert notifications at once.
func generateAllAlerts(trackerDBHandle *sql.DB, currUserID string, databaseID string, userID string, userIsAdmin bool) (*AlertGenerationResult, error) {

	// Create a config/context object used for calculating the calculated fields for alert generation. This same
	// config can be reused by the alert generation for all the records.
	calcFieldUpdateConfig, err := calcField.CreateCalcFieldUpdateConfig(trackerDBHandle, currUserID, databaseID)
	if err != nil {
		return nil, fmt.Errorf("MapAllRecordUpdatesToFieldValues: %v", err)
	}

	// Get all cell updates at once for the given tracking databse. This is much faster than doing a database call for each and every record.
	recordCellUpdateMap, err := record.GetAllNonDraftCellUpdates(trackerDBHandle, databaseID, record.FullyCommittedCellUpdatesChangeSetID)
	if err != nil {
		return nil, fmt.Errorf("MapAllRecordUpdatesToFieldValues: %v", err)
	}

	alerts, alertErr := getAlertsWithUserNotification(trackerDBHandle, databaseID, userID, userIsAdmin)
	if alertErr != nil {
		return nil, fmt.Errorf("GenerateRecordAlerts: Error getting alerts: %v", alertErr)
	}
	alertsByID := map[string]Alert{}
	for _, currAlert := range alerts {
		alertsByID[currAlert.AlertID] = currAlert
	}
	alertContexts, contextErr := createAlertGenerationContexts(trackerDBHandle, currUserID, alerts)
	if contextErr != nil {
		return nil, fmt.Errorf("GenerateRecordAlerts: Error setting up alert contexts: %v", contextErr)

	}

	recordIDRecordMap, err := record.GetNonDraftRecordIDRecordMap(trackerDBHandle, databaseID)
	if err != nil {
		return nil, fmt.Errorf("MapAllRecordUpdatesToFieldValues: %v", err)
	}

	alertNotifications := []AlertNotification{}
	resultsChan := make(chan AlertProcessingResult)

	for currRecordID, currRecCellUpdates := range recordCellUpdateMap {

		currRecord, recFound := recordIDRecordMap[currRecordID]
		if !recFound {
			return nil, fmt.Errorf("MapAllRecordUpdatesToFieldValues: %v", err)
		}

		alertProcessConfig := RecordAlertProcessingConfig{
			RecordID:        currRecordID,
			Record:          currRecord,
			CalcFieldConfig: calcFieldUpdateConfig,
			RecCellUpdates:  currRecCellUpdates,
			AlertContexts:   alertContexts}

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

	prunedNotifications := pruneLatestUniqueAlertsByRecordAndAlert(alertNotifications)

	// Sort in reverse chronological order
	sort.Sort(NotificationByTime(prunedNotifications))

	latestAlertTime := getLatestNotificationTime(trackerDBHandle, userID, databaseID)
	alertsNotSeen := 0
	for _, currNotification := range prunedNotifications {
		if latestAlertTime.Before(currNotification.Timestamp) {
			alertsNotSeen++
		}
	}

	alertGenResults := AlertGenerationResult{
		AlertsByID:       alertsByID,
		Notifications:    prunedNotifications,
		LatestAlertTime:  latestAlertTime,
		NumAlertsNotSeen: alertsNotSeen}

	return &alertGenResults, nil

}

// GenerateOneRecordAlerts is a top-level entry point for regenerating the alerts for an entire
// tracker, but a single recordID
func GenerateOneRecordAlerts(trackerDBHandle *sql.DB,
	currUserID string, databaseID string, recordID string, userID string, userIsAdmin bool) ([]AlertNotification, error) {

	log.Printf("Regenerating alerts ...")

	calcFieldUpdateConfig, err := calcField.CreateCalcFieldUpdateConfig(trackerDBHandle, currUserID, databaseID)
	if err != nil {
		return nil, fmt.Errorf("MapAllRecordUpdatesToFieldValues: %v", err)
	}

	// Retrieve a list of sorted record updates.
	recCellUpdates, getErr := record.GetRecordCellUpdates(trackerDBHandle, recordID, record.FullyCommittedCellUpdatesChangeSetID)
	if getErr != nil {
		return nil, fmt.Errorf("GenerateRecordAlerts: failure retrieving cell updates for record = %v: error = %v",
			recordID, getErr)
	}

	alerts, alertErr := getAlertsWithUserNotification(trackerDBHandle, databaseID, userID, userIsAdmin)
	if alertErr != nil {
		return nil, fmt.Errorf("GenerateRecordAlerts: Error getting alerts: %v", alertErr)
	}
	alertContexts, contextErr := createAlertGenerationContexts(trackerDBHandle, currUserID, alerts)
	if contextErr != nil {
		return nil, fmt.Errorf("GenerateRecordAlerts: Error setting up alert contexts: %v", contextErr)

	}

	alertProcessConfig := RecordAlertProcessingConfig{
		TrackerDBHandle: trackerDBHandle,
		RecordID:        recordID,
		CalcFieldConfig: calcFieldUpdateConfig,
		RecCellUpdates:  recCellUpdates,
		AlertContexts:   alertContexts}

	return generateOneRecordAlertsFromConfig(alertProcessConfig)
}
