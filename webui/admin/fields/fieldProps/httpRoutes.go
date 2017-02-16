package fieldProps

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/field/{fieldID}", editFieldPropsPage)
}
