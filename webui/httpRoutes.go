package webui

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/webui/admin"
	"resultra/datasheet/webui/dashboard"
	"resultra/datasheet/webui/form"
	"resultra/datasheet/webui/homePage"
	"resultra/datasheet/webui/itemList"
)

func init() {

	router := mux.NewRouter()

	form.RegisterHTTPHandlers(router)
	itemList.RegisterHTTPHandlers(router)
	dashboard.RegisterHTTPHandlers(router)
	homePage.RegisterHTTPHandlers(router)
	admin.RegisterHTTPHandlers(router)

	http.Handle("/", router)
}
