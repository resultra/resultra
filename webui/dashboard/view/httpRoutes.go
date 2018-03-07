package view

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/dashboard/view/contentLayout", dashboardContentLayout)
	mainRouter.HandleFunc("/dashboard/view/sidebarLayout", dashboardSidebarLayout)
}
