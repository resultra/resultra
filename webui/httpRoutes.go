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

	router.HandleFunc("/editRecord/{layoutID}", editRecord)

	router.HandleFunc("/tableProps", tableProps)
	router.HandleFunc("/designForm/{layoutID}", designForm)

	http.Handle("/", router)
}
