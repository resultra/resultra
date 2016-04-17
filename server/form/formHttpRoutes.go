package form

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterHTTPHandlers(apiRouter *mux.Router) {

	apiRouter.HandleFunc("/api/newLayout", newLayout)
	//	apiRouter.HandleFunc("/api/getLayoutEditInfo", getLayoutEditInfo)

}

func init() {

	formRouter := mux.NewRouter()

	formRouter.HandleFunc("/api/frm/getFormInfo", getFormInfo)

	http.Handle("/api/frm/", formRouter)
}
