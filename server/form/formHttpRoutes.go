package form

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/form/textBox"
)

func RegisterHTTPHandlers(apiRouter *mux.Router) {

	textBox.RegisterHTTPHandlers(apiRouter)

	apiRouter.HandleFunc("/api/newLayout", newLayout)
	//	apiRouter.HandleFunc("/api/getLayoutEditInfo", getLayoutEditInfo)

}

func init() {

	formRouter := mux.NewRouter()

	formRouter.HandleFunc("/api/frm/getFormInfo", getFormInfo)

	http.Handle("/api/frm/", formRouter)
}
