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
	router.HandleFunc("/editLayout/{layoutID}", editLayout)
	router.HandleFunc("/editRecord/{layoutID}/{recordID}", editRecord)

	http.Handle("/", router)
}
