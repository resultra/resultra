package webui

import (
	"net/http"
	"resultra/datasheet/controller"
)

func init() {

	controller.RegisterAPIHTTPHandlers()

	http.HandleFunc("/", root)
	http.HandleFunc("/pageinfo", pageinfo)
	http.HandleFunc("/dataTable", dataTable)
	http.HandleFunc("/addRow", addRow)
	http.HandleFunc("/editLayout", editLayout)
}
