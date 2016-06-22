package recordSort

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {

	sortRouter := mux.NewRouter()

	sortRouter.HandleFunc("/api/recordSort/saveFormSortRules", saveFormSortRulesAPI)
	sortRouter.HandleFunc("/api/recordSort/getFormSortRules", getFormSortRulesAPI)

	http.Handle("/api/recordSort/", sortRouter)
}

func saveFormSortRulesAPI(w http.ResponseWriter, r *http.Request) {

	sortRuleParams := NewSortRuleParams{}
	if err := api.DecodeJSONRequest(r, &sortRuleParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	newSortRule, newErr := saveFormSortRules(sortRuleParams)
	if newErr != nil {
		api.WriteErrorResponse(w, newErr)
		return
	} else {
		api.WriteJSONResponse(w, newSortRule)
	}

}

type GetFormSortRulesParams struct {
	ParentFormID string `json:"parentFormID"`
}

func getFormSortRulesAPI(w http.ResponseWriter, r *http.Request) {

	params := GetFormSortRulesParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	sortRules, getErr := GetFormSortRules(params.ParentFormID)
	if getErr != nil {
		api.WriteErrorResponse(w, getErr)
		return
	} else {
		api.WriteJSONResponse(w, sortRules)
	}

}
