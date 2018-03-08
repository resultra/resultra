package alertListView

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/alertListView/contentLayout", alertListContentLayout)
}
