package form

import (
	"github.com/gorilla/mux"
	"resultra/datasheet/webui/form/design"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/designForm/{formID}", design.DesignForm)
}
