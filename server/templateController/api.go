package templateController

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/generic/userAuth"
)

type DummyStructForInclude struct {
	Val int64
}

func init() {
	templateRouter := mux.NewRouter()

	//	templateRouter.HandleFunc("/api/template/import", importTemplateAPI)

	templateRouter.HandleFunc("/api/template/save", saveTemplateAPI)

	http.Handle("/api/template/", templateRouter)
}

func saveTemplateAPI(w http.ResponseWriter, r *http.Request) {

	params := SaveTemplateParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	currUserID, userErr := userAuth.GetCurrentUserID(r)
	if userErr != nil {
		api.WriteJSONResponse(w, fmt.Errorf("Can't verify user authentication"))
		return
	}

	err := saveTemplate(currUserID, params)
	if err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, true)
	}

}
