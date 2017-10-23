package webui

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/webui/admin"
	"resultra/datasheet/webui/alertPage"
	"resultra/datasheet/webui/dashboard"
	"resultra/datasheet/webui/formPage"
	"resultra/datasheet/webui/homePage"
	"resultra/datasheet/webui/itemList"
	"resultra/datasheet/webui/mainWindow"
)

func init() {

	router := mux.NewRouter()

	itemList.RegisterHTTPHandlers(router)
	dashboard.RegisterHTTPHandlers(router)
	homePage.RegisterHTTPHandlers(router)
	admin.RegisterHTTPHandlers(router)
	formPage.RegisterHTTPHandlers(router)
	alertPage.RegisterHTTPHandlers(router)
	mainWindow.RegisterHTTPHandlers(router)

	http.Handle("/", router)
}
