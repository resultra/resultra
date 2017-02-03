package newItem

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

type DummyStructForInclude struct{ Val int64 }

func init() {
	newItemRouter := mux.NewRouter()

	newItemRouter.HandleFunc("/api/newItem/newPreset", newPresetAPI)
	newItemRouter.HandleFunc("/api/newItem/getPresets", getPresetsAPI)

	http.Handle("/api/newItem/", newItemRouter)
}

func newPresetAPI(w http.ResponseWriter, r *http.Request) {

	params := NewNewItemPresetParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	newPreset, err := newNewItemPreset(params)
	if err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *newPreset)
	}

}

func getPresetsAPI(w http.ResponseWriter, r *http.Request) {

	params := GetPresetListParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	presets, err := getAllPresets(params.ParentDatabaseID)
	if err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, presets)
	}

}
