package formPage

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {

	mainRouter.HandleFunc("/submitForm/{sharedLinkID}", submitFormPage)
	mainRouter.HandleFunc("/newItem/{formLinkID}", newItemFormPage)

	mainRouter.Path("/viewItem/{formID}/{recordID}").HandlerFunc(viewFormPage)

	mainRouter.Path("/viewItem/{formID}/{recordID}").Queries("col", "{col}").HandlerFunc(viewFormPage)
	mainRouter.Path("/viewItem/{formID}/{recordID}").Queries("frm", "{frm}").HandlerFunc(viewFormPage)

}
