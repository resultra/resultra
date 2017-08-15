package alert

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/userRole"
)

type DummyStructForInclude struct {
	Val int64
}

func init() {

	alertRouter := mux.NewRouter()

	alertRouter.HandleFunc("/api/alert/new", newAlertAPI)
	alertRouter.HandleFunc("/api/alert/get", getAlertAPI)
	alertRouter.HandleFunc("/api/alert/list", getAlertListAPI)

	alertRouter.HandleFunc("/api/alert/getNotificationList", getAlertNotificationListAPI)

	alertRouter.HandleFunc("/api/alert/setName", setAlertName)
	alertRouter.HandleFunc("/api/alert/setForm", setForm)

	alertRouter.HandleFunc("/api/alert/setConditions", setConditions)

	alertRouter.HandleFunc("/api/alert/validateFormName", validateAlertNameAPI)
	alertRouter.HandleFunc("/api/alert/validateNewAlertName", validateNewAlertNameAPI)

	//	alertRouter.HandleFunc("/api/alert/delete", deleteAlertAPI)

	http.Handle("/api/alert/", alertRouter)
}

func newAlertAPI(w http.ResponseWriter, r *http.Request) {

	var params NewAlertParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdmin(
		r, params.ParentDatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if formRef, err := newAlert(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *formRef)
	}

}

type GetAlertParams struct {
	AlertID string `json:"alertID"`
}

func getAlertAPI(w http.ResponseWriter, r *http.Request) {

	var params GetAlertParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if alert, err := GetAlert(params.AlertID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *alert)
	}

}

func getAlertListAPI(w http.ResponseWriter, r *http.Request) {

	var params GetAlertListParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if alerts, err := getAllAlerts(params.ParentDatabaseID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, alerts)
	}

}

func getAlertNotificationListAPI(w http.ResponseWriter, r *http.Request) {

	var params GetAlertListParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if notifications, err := generateAllAlerts(params.ParentDatabaseID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, notifications)
	}

}

func processAlertPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater AlertPropUpdater) {
	if updatedAlert, err := updateAlertProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, updatedAlert)
	}
}

func setAlertName(w http.ResponseWriter, r *http.Request) {
	var params SetAlertNameParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processAlertPropUpdate(w, r, params)
}

func setConditions(w http.ResponseWriter, r *http.Request) {
	var params SetConditionsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processAlertPropUpdate(w, r, params)
}

func setForm(w http.ResponseWriter, r *http.Request) {
	var params SetFormParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processAlertPropUpdate(w, r, params)
}

func validateAlertNameAPI(w http.ResponseWriter, r *http.Request) {

	alertName := r.FormValue("alertName")
	alertID := r.FormValue("alertID")

	if err := validateAlertName(alertID, alertName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}

func validateNewAlertNameAPI(w http.ResponseWriter, r *http.Request) {

	alertName := r.FormValue("alertName")
	databaseID := r.FormValue("databaseID")

	if err := validateNewFormName(databaseID, alertName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}
