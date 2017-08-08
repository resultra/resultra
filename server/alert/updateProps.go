package alert

import (
	"fmt"
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
	alertForUpdate, getErr := getAlert(propUpdater.getAlertID())
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
