package webui

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/controller"
)

func init() {

	controller.RegisterAPIHTTPHandlers()

	router := mux.NewRouter()
	router.HandleFunc("/", root)
	router.HandleFunc("/pageinfo", pageinfo)
	router.HandleFunc("/dataTable", dataTable)
	router.HandleFunc("/addRow", addRow)
	router.HandleFunc("/editRecord/{layoutID}", editRecord)

	router.HandleFunc("/tableProps", tableProps)
	router.HandleFunc("/editLayout/{layoutID}", editLayout)

	http.Handle("/", router)
}
