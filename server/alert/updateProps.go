package alert

import (
	"fmt"
	"resultra/datasheet/server/recordFilter"
)

type AlertIDInterface interface {
	getAlertID() string
}

type AlertIDHeader struct {
	AlertID string `json:"alertID"`
}

func (idHeader AlertIDHeader) getAlertID() string {
	return idHeader.AlertID
}

type AlertPropUpdater interface {
	AlertIDInterface
	updateProps(alert *Alert) error
}

func updateAlertProps(propUpdater AlertPropUpdater) (*Alert, error) {

	// Retrieve the bar chart from the data store
	alertForUpdate, getErr := GetAlert(propUpdater.getAlertID())
	if getErr != nil {
		return nil, fmt.Errorf("updateAlertProps: Unable to get existing alert: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(alertForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateFormProps: Unable to update existing form properties: %v", propUpdateErr)
	}

	alert, updateErr := updateExistingAlert(propUpdater.getAlertID(), alertForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateAlertProps: Unable to update existing form properties: datastore update error =  %v", updateErr)
	}

	return alert, nil
}

type SetAlertNameParams struct {
	AlertIDHeader
	AlertName string `json:"alertName"`
}

func (updateParams SetAlertNameParams) updateProps(alert *Alert) error {

	// TODO - Validate name

	alert.Name = updateParams.AlertName

	return nil
}

type SetConditionParams struct {
	AlertIDHeader
	Condition *AlertCondition `json:"conditions"`
}

func (updateParams SetConditionParams) updateProps(alert *Alert) error {

	// TODO - Validate conditions

	alert.Properties.Condition = updateParams.Condition

	return nil
}

type SetFormParams struct {
	AlertIDHeader
	FormID string `json:"formID"`
}

func (updateParams SetFormParams) updateProps(alert *Alert) error {

	// TODO - Validate conditions

	alert.Properties.FormID = updateParams.FormID

	return nil
}

type SetSummaryFieldParams struct {
	AlertIDHeader
	SummaryFieldID string `json:"summaryFieldID"`
}

func (updateParams SetSummaryFieldParams) updateProps(alert *Alert) error {

	// TODO - Validate field ID

	alert.Properties.SummaryFieldID = updateParams.SummaryFieldID

	return nil
}

type SetTriggerConditionsParams struct {
	AlertIDHeader
	TriggerConditions recordFilter.RecordFilterRuleSet `json:"triggerConditions"`
}

func (updateParams SetTriggerConditionsParams) updateProps(alert *Alert) error {

	// TODO - Validate filter rules before saving
	alert.Properties.TriggerConditions = updateParams.TriggerConditions

	return nil
}

type SetCaptionMessageParams struct {
	AlertIDHeader
	CaptionMessage string `json:"captionMessage"`
}

func (updateParams SetCaptionMessageParams) updateProps(alert *Alert) error {

	// For internal storage, replace occurences of field references with the immutable field ID.
	// This allows the field reference name to change without affecting the caption template.
	encodedCaptionMsg, err := replaceFieldRefWithFieldID(updateParams.CaptionMessage, alert.ParentDatabaseID)
	if err != nil {
		return fmt.Errorf("SetCaptionMessageParams.updateProps: %v", err)
	}

	alert.Properties.CaptionMessage = encodedCaptionMsg

	return nil
}
