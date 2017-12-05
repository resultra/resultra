package webui

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/webui/admin"
	"resultra/datasheet/webui/dashboard"
	"resultra/datasheet/webui/formPage"
	"resultra/datasheet/webui/homePage"
	"resultra/datasheet/webui/itemList"
	"resultra/datasheet/webui/mainWindow"
	"resultra/datasheet/webui/templatePage"
	"resultra/datasheet/webui/userAdmin"
	"resultra/datasheet/webui/workspaceAdmin"
)

func init() {

	router := mux.NewRouter()

	itemList.RegisterHTTPHandlers(router)
	dashboard.RegisterHTTPHandlers(router)
	homePage.RegisterHTTPHandlers(router)
	admin.RegisterHTTPHandlers(router)
	formPage.RegisterHTTPHandlers(router)
	mainWindow.RegisterHTTPHandlers(router)
	templatePage.RegisterHTTPHandlers(router)
	userAdmin.RegisterHTTPHandlers(router)
	workspaceAdmin.RegisterHTTPHandlers(router)

	http.Handle("/", router)
}
