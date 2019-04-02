// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package stringValidation

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/resultra/resultra/server/generic/api"
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
