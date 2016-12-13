package form

import (
	"github.com/gorilla/mux"
	"resultra/datasheet/webui/form/design"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/frm/{formID}", design.DesignForm)
}
