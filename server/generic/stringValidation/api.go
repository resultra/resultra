package stringValidation

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"resultra/tracker/server/generic/api"
)

type DummyStructForInclude struct {
	Val int64
}

func init() {

	validationRouter := mux.NewRouter()

	validationRouter.HandleFunc("/api/generic/stringValidation/validateItemLabel", validateItemLabelAPI)
	validationRouter.HandleFunc("/api/generic/stringValidation/validateOptionalItemLabel", validateOptionalItemLabelAPI)

	http.Handle("/api/generic/stringValidation/", validationRouter)
}

func validateItemLabelAPI(w http.ResponseWriter, r *http.Request) {

	label := r.FormValue("label")

	if err := validateItemLabel(label); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}

func validateOptionalItemLabelAPI(w http.ResponseWriter, r *http.Request) {

	label := r.FormValue("label")

	if err := ValidateOptionalItemLabel(label); err != nil {
		api.WriteJSONResponse(w, fmt.Sprintf("%v", err))
		return
	}

	response := true
	api.WriteJSONResponse(w, response)

}
