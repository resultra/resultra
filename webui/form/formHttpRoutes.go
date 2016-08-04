package form

import (
	"github.com/gorilla/mux"
	"resultra/datasheet/webui/form/design"
	"resultra/datasheet/webui/form/view"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/viewForm/{formID}", view.ViewForm)
	mainRouter.HandleFunc("/designForm/{formID}", design.DesignForm)
}
