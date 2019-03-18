package recordUpdate

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/generic/api"
	"resultra/tracker/server/common/userAuth"
	"resultra/tracker/server/record"
)

type DummyStructForInclude struct {
	Val int64
}

func init() {
	recordUpdateRouter := mux.NewRouter()

	recordUpdateRouter.HandleFunc("/api/recordUpdate/newRecord", newRecordAPI)

	recordUpdateRouter.HandleFunc("/api/recordUpdate/setBoolFieldValue", setBoolFieldValue)
	recordUpdateRouter.HandleFunc("/api/recordUpdate/setNumberFieldValue", setNumberFieldValue)
	recordUpdateRouter.HandleFunc("/api/recordUpdate/setTextFieldValue", setTextFieldValue)
	recordUpdateRouter.HandleFunc("/api/recordUpdate/setTimeFieldValue", setTimeFieldValue)
	recordUpdateRouter.HandleFunc("/api/recordUpdate/setLongTextFieldValue", setLongTextFieldValue)
	recordUpdateRouter.HandleFunc("/api/recordUpdate/setUserFieldValue", setUserFieldValue)
	recordUpdateRouter.HandleFunc("/api/recordUpdate/setUsersFieldValue", setUsersFieldValue)
	recordUpdateRouter.HandleFunc("/api/recordUpdate/setFileFieldValue", setFileFieldValue)
	recordUpdateRouter.HandleFunc("/api/recordUpdate/setImageFieldValue", setImageFieldValue)

	recordUpdateRouter.HandleFunc("/api/recordUpdate/setAttachmentFieldValue", setAttachmentFieldValue)
	recordUpdateRouter.HandleFunc("/api/recordUpdate/setCommentFieldValue", setCommentFieldValue)
	recordUpdateRouter.HandleFunc("/api/recordUpdate/setLabelFieldValue", setLabelFieldValue)
	recordUpdateRouter.HandleFunc("/api/recordUpdate/setEmailAddrFieldValue", setEmailAddrFieldValue)
	recordUpdateRouter.HandleFunc("/api/recordUpdate/setUrlLinkFieldValue", setUrlFieldValue)

	recordUpdateRouter.HandleFunc("/api/recordUpdate/commitChangeSet", commitChangeSetAPI)
	recordUpdateRouter.HandleFunc("/api/recordUpdate/setDefaultValues", setDefaultValuesAPI)

	http.Handle("/api/recordUpdate/", recordUpdateRouter)
}

func newRecordAPI(w http.ResponseWriter, r *http.Request) {

	params := record.NewRecordParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	currUserID, userErr := userAuth.GetCurrentUserID(r)
	if userErr != nil {
		api.WriteJSONResponse(w, fmt.Errorf("Can't verify user authentication"))
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	newRecordRef, err := newRecord(trackerDBHandle, currUserID, params)
	if err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *newRecordRef)
	}

}

func setTextFieldValue(w http.ResponseWriter, r *http.Request) {

	setValParams := record.SetRecordTextValueParams{}
	if err := api.DecodeJSONRequest(r, &setValParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	updatedRecordRef, setErr := updateRecordValue(r, setValParams)
	if setErr != nil {
		api.WriteErrorResponse(w, setErr)
		return
	} else {
		api.WriteJSONResponse(w, updatedRecordRef)
	}

}

func setLongTextFieldValue(w http.ResponseWriter, r *http.Request) {

	// Reuse same parameter struct as setting text.
	setValParams := record.SetRecordLongTextValueParams{}
	if err := api.DecodeJSONRequest(r, &setValParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	updatedRecordRef, setErr := updateRecordValue(r, setValParams)
	if setErr != nil {
		api.WriteErrorResponse(w, setErr)
		return
	} else {
		api.WriteJSONResponse(w, updatedRecordRef)
	}

}

func setNumberFieldValue(w http.ResponseWriter, r *http.Request) {

	setValParams := record.SetRecordNumberValueParams{}
	if err := api.DecodeJSONRequest(r, &setValParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	updatedRecordRef, setErr := updateRecordValue(r, setValParams)
	if setErr != nil {
		api.WriteErrorResponse(w, setErr)
		return
	} else {
		api.WriteJSONResponse(w, updatedRecordRef)
	}

}

func setUserFieldValue(w http.ResponseWriter, r *http.Request) {

	setValParams := record.SetRecordUserValueParams{}
	if err := api.DecodeJSONRequest(r, &setValParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	updatedRecordRef, setErr := updateRecordValue(r, setValParams)
	if setErr != nil {
		api.WriteErrorResponse(w, setErr)
		return
	} else {
		api.WriteJSONResponse(w, updatedRecordRef)
	}

}

func setUsersFieldValue(w http.ResponseWriter, r *http.Request) {

	setValParams := record.SetRecordUsersValueParams{}
	if err := api.DecodeJSONRequest(r, &setValParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	updatedRecordRef, setErr := updateRecordValue(r, setValParams)
	if setErr != nil {
		api.WriteErrorResponse(w, setErr)
		return
	} else {
		api.WriteJSONResponse(w, updatedRecordRef)
	}

}

func setLabelFieldValue(w http.ResponseWriter, r *http.Request) {

	setValParams := record.SetRecordLabelValueParams{}
	if err := api.DecodeJSONRequest(r, &setValParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	updatedRecordRef, setErr := updateRecordValue(r, setValParams)
	if setErr != nil {
		api.WriteErrorResponse(w, setErr)
		return
	} else {
		api.WriteJSONResponse(w, updatedRecordRef)
	}

}

func setEmailAddrFieldValue(w http.ResponseWriter, r *http.Request) {

	setValParams := record.SetRecordEmailAddrValueParams{}
	if err := api.DecodeJSONRequest(r, &setValParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	updatedRecordRef, setErr := updateRecordValue(r, setValParams)
	if setErr != nil {
		api.WriteErrorResponse(w, setErr)
		return
	} else {
		api.WriteJSONResponse(w, updatedRecordRef)
	}

}

func setUrlFieldValue(w http.ResponseWriter, r *http.Request) {

	setValParams := record.SetRecordUrlValueParams{}
	if err := api.DecodeJSONRequest(r, &setValParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	updatedRecordRef, setErr := updateRecordValue(r, setValParams)
	if setErr != nil {
		api.WriteErrorResponse(w, setErr)
		return
	} else {
		api.WriteJSONResponse(w, updatedRecordRef)
	}

}

func setBoolFieldValue(w http.ResponseWriter, r *http.Request) {

	setValParams := record.SetRecordBoolValueParams{}
	if err := api.DecodeJSONRequest(r, &setValParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	updatedRecordRef, setErr := updateRecordValue(r, setValParams)
	if setErr != nil {
		api.WriteErrorResponse(w, setErr)
		return
	} else {
		api.WriteJSONResponse(w, updatedRecordRef)
	}

}

func setTimeFieldValue(w http.ResponseWriter, r *http.Request) {

	setValParams := record.SetRecordTimeValueParams{}
	if err := api.DecodeJSONRequest(r, &setValParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	updatedRecordRef, setErr := updateRecordValue(r, setValParams)
	if setErr != nil {
		api.WriteErrorResponse(w, setErr)
		return
	} else {
		api.WriteJSONResponse(w, updatedRecordRef)
	}

}

func setCommentFieldValue(w http.ResponseWriter, r *http.Request) {

	setValParams := record.SetRecordCommentValueParams{}
	if err := api.DecodeJSONRequest(r, &setValParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	updatedRecordRef, setErr := updateRecordValue(r, setValParams)
	if setErr != nil {
		api.WriteErrorResponse(w, setErr)
		return
	} else {
		api.WriteJSONResponse(w, updatedRecordRef)
	}

}

func setAttachmentFieldValue(w http.ResponseWriter, r *http.Request) {

	setValParams := record.SetRecordAttachmentValueParams{}
	if err := api.DecodeJSONRequest(r, &setValParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	updatedRecordRef, setErr := updateRecordValue(r, setValParams)
	if setErr != nil {
		api.WriteErrorResponse(w, setErr)
		return
	} else {
		api.WriteJSONResponse(w, updatedRecordRef)
	}

}

func setFileFieldValue(w http.ResponseWriter, r *http.Request) {

	setValParams := record.SetRecordFileAddrValueParams{}
	if err := api.DecodeJSONRequest(r, &setValParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	updatedRecordRef, setErr := updateRecordValue(r, setValParams)
	if setErr != nil {
		api.WriteErrorResponse(w, setErr)
		return
	} else {
		api.WriteJSONResponse(w, updatedRecordRef)
	}

}

func setImageFieldValue(w http.ResponseWriter, r *http.Request) {

	setValParams := record.SetRecordImageValueParams{}
	if err := api.DecodeJSONRequest(r, &setValParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	updatedRecordRef, setErr := updateRecordValue(r, setValParams)
	if setErr != nil {
		api.WriteErrorResponse(w, setErr)
		return
	} else {
		api.WriteJSONResponse(w, updatedRecordRef)
	}

}

func commitChangeSetAPI(w http.ResponseWriter, r *http.Request) {
	var params CommitChangeSetParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	currUserID, userErr := userAuth.GetCurrentUserID(r)
	if userErr != nil {
		api.WriteJSONResponse(w, fmt.Errorf("Can't verify user authentication"))
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if updatedRecordRef, err := commitChangeSet(trackerDBHandle, currUserID, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, updatedRecordRef)
	}
}

func setDefaultValuesAPI(w http.ResponseWriter, r *http.Request) {
	var params record.SetDefaultValsParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if updatedRecordRef, err := setDefaultValues(r, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, updatedRecordRef)
	}
}
