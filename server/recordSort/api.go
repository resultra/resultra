package recordSort

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {

	sortRouter := mux.NewRouter()

	sortRouter.HandleFunc("/api/recordSort/saveSortRule", newSortRuleAPI)

	http.Handle("/api/recordSort/", sortRouter)
}

func newSortRuleAPI(w http.ResponseWriter, r *http.Request) {

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
