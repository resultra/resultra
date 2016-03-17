package webui

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/controller"
)

func init() {

	controller.RegisterAPIHTTPHandlers()

	router := mux.NewRouter()

	router.HandleFunc("/", home)

	router.HandleFunc("/viewForm/{layoutID}", viewForm)

	router.HandleFunc("/tableProps", tableProps)
	router.HandleFunc("/designForm/{layoutID}", designForm)

	router.HandleFunc("/designDashboard/{dashboardID}", designDashboard)

	http.Handle("/", router)
}
