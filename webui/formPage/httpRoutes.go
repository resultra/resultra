package formPage

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/submitForm/{sharedLinkID}", submitFormPage)
	mainRouter.HandleFunc("/newItem/{formLinkID}", newItemFormPage)
	mainRouter.HandleFunc("/viewItem/{formID}/{recordID}", viewFormPage)
}
