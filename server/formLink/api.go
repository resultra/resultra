package formLink

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

type DummyStructForInclude struct{ Val int64 }

func init() {
	formLinkRouter := mux.NewRouter()

	formLinkRouter.HandleFunc("/api/formLink/new", newFormLinkAPI)
	formLinkRouter.HandleFunc("/api/formLink/get", getFormLinkAPI)
	formLinkRouter.HandleFunc("/api/formLink/getList", getFormLinksAPI)
	formLinkRouter.HandleFunc("/api/formLink/setDefaultVals", setDefaultVals)

	http.Handle("/api/formLink/", formLinkRouter)
}

func newFormLinkAPI(w http.ResponseWriter, r *http.Request) {

	params := NewFormLinkParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	newPreset, err := newFormLink(params)
	if err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *newPreset)
	}

}

func getFormLinkAPI(w http.ResponseWriter, r *http.Request) {

	params := GetFormLinkParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	presets, err := GetFormLink(params.FormLinkID)
	if err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, presets)
	}

}

func getFormLinksAPI(w http.ResponseWriter, r *http.Request) {

	params := GetFormLinkListParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	presets, err := getAllFormLinks(params.ParentDatabaseID)
	if err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, presets)
	}

}

func processFormLinkPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater FormLinkPropUpdater) {
	if headerRef, err := updateFormLinkProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, headerRef)
	}
}

func setDefaultVals(w http.ResponseWriter, r *http.Request) {
	var params FormLinkDefaultValParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processFormLinkPropUpdate(w, r, params)
}
