package form

import (
	"github.com/gorilla/mux"
	"resultra/datasheet/server/form/textBox"
)

func RegisterHTTPHandlers(apiRouter *mux.Router) {

	textBox.RegisterHTTPHandlers(apiRouter)
	apiRouter.HandleFunc("/api/newLayout", newLayout)
	apiRouter.HandleFunc("/api/getLayoutEditInfo", getLayoutEditInfo)

}
