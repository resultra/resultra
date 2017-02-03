package formLink

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

type DummyStructForInclude struct{ Val int64 }

func init() {
	formLinkRouter := mux.NewRouter()

	formLinkRouter.HandleFunc("/api/formLink/newPreset", newPresetAPI)
	formLinkRouter.HandleFunc("/api/formLink/getPresets", getPresetsAPI)

	http.Handle("/api/formLink/", formLinkRouter)
}

func newPresetAPI(w http.ResponseWriter, r *http.Request) {

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

func getPresetsAPI(w http.ResponseWriter, r *http.Request) {

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
