package textBox

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterHTTPHandlers(apiRouter *mux.Router) {

	/*	apiRouter.HandleFunc("/api/newLayoutContainer", newLayoutContainer)
		apiRouter.HandleFunc("/api/resizeLayoutContainer", resizeLayoutContainer)
		apiRouter.HandleFunc("/api/getLayoutContainers", getLayoutContainers)
	*/
}

func init() {
	textBoxRouter := mux.NewRouter()

	textBoxRouter.HandleFunc("/api/frm/textBox/new", newTextBox)
	textBoxRouter.HandleFunc("/api/frm/textBox/resize", resizeTextBox)

	http.Handle("/api/frm/textBox/", textBoxRouter)
}
