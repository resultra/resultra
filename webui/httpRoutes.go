package webui

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server"
	"resultra/datasheet/webui/admin"
	"resultra/datasheet/webui/dashboard"
	"resultra/datasheet/webui/form"
	"resultra/datasheet/webui/homePage"
)

func init() {

	server.RegisterAPIHTTPHandlers()

	router := mux.NewRouter()

	form.RegisterHTTPHandlers(router)
	dashboard.RegisterHTTPHandlers(router)
	homePage.RegisterHTTPHandlers(router)
	admin.RegisterHTTPHandlers(router)

	http.Handle("/", router)
}
