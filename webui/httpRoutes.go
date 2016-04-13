package webui

import (
	"admin"
	"dashboard"
	"form"
	"github.com/gorilla/mux"
	"home"
	"net/http"
	"resultra/datasheet/server"
)

func init() {

	server.RegisterAPIHTTPHandlers()

	router := mux.NewRouter()

	form.RegisterHTTPHandlers(router)
	dashboard.RegisterHTTPHandlers(router)
	home.RegisterHTTPHandlers(router)
	admin.RegisterHTTPHandlers(router)

	http.Handle("/", router)
}
