package form

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

	formRouter := mux.NewRouter()

	formRouter.HandleFunc("/api/frm/getFormInfo", getFormInfoAPI)
	formRouter.HandleFunc("/api/frm/new", newFormAPI)
	formRouter.HandleFunc("/api/frm/get", getFormAPI)
	formRouter.HandleFunc("/api/frm/list", getFormListAPI)
	formRouter.HandleFunc("/api/frm/formsByID", getFormByIDAPI)

	formRouter.HandleFunc("/api/frm/setName", setFormName)
	formRouter.HandleFunc("/api/frm/setLayout", setLayout)
	formRouter.HandleFunc("/api/frm/deleteComponent", deleteComponentAPI)

	formRouter.HandleFunc("/api/frm/validateFormName", validateFormNameAPI)
	formRouter.HandleFunc("/api/frm/validateNewFormName", validateNewFormNameAPI)

	http.Handle("/api/frm/", formRouter)
}

func newFormAPI(w http.ResponseWriter, r *http.Request) {

	var params NewFormParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if verifyErr := userRole.VerifyCurrUserIsDatabaseAdmin(
		r, params.ParentDatabaseID); verifyErr != nil {
		api.WriteErrorResponse(w, verifyErr)
		return
	}

	if formRef, err := newForm(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *formRef)
	}

}

type GetFormParams struct {
	FormID string `json:"formID"`
}

func getFormAPI(w http.ResponseWriter, r *http.Request) {

	var params GetFormParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if theForm, err := GetForm(params.FormID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *theForm)
	}

}

func getFormListAPI(w http.ResponseWriter, r *http.Request) {

	var params GetFormListParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if forms, err := GetAllForms(params.ParentDatabaseID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, forms)
	}

}

func getFormByIDAPI(w http.ResponseWriter, r *http.Request) {

	var params GetFormListParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if forms, err := GetAllForms(params.ParentDatabaseID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {

		formsByID := map[string]Form{}
		for _, currFormInfo := range forms {
			formsByID[currFormInfo.FormID] = currFormInfo
		}

		api.WriteJSONResponse(w, formsByID)
	}

}

func getFormInfoAPI(w http.ResponseWriter, r *http.Request) {

	var params GetFormInfoParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if formInfo, err := GetFormInfo(params.FormID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *formInfo)
	}

}

func deleteComponentAPI(w http.ResponseWriter, r *http.Request) {

	var params DeleteComponentParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if err := deleteComponent(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, true)
	}

}

func processFormPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater FormPropUpdater) {
	if updatedForm, err := updateFormProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, updatedForm)
	}
}

func setFormName(w http.ResponseWriter, r *http.Request) {
	var params SetFormNameParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processFormPropUpdate(w, r, params)
}

func setLayout(w http.ResponseWriter, r *http.Request) {
	var params SetLayoutParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processFormPropUpdate(w, r, params)
}

func validateFormNameAPI(w http.ResponseWriter, r *http.Request) {

	formName := r.FormValue("formName")
	formID := r.FormValue("formID")

	if err := validateFormName(formID, formName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}

func validateNewFormNameAPI(w http.ResponseWriter, r *http.Request) {

	formName := r.FormValue("formName")
	databaseID := r.FormValue("databaseID")

	if err := validateNewFormName(databaseID, formName); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}
