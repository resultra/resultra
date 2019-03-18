package webui

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/tracker/webui/admin"
	"resultra/tracker/webui/alertListView"
	designDashboard "resultra/tracker/webui/dashboard/design"
	dashboardView "resultra/tracker/webui/dashboard/view"
	"resultra/tracker/webui/formPage"
	"resultra/tracker/webui/homePage"
	"resultra/tracker/webui/itemList"
	"resultra/tracker/webui/itemView"
	"resultra/tracker/webui/mainWindow"
	"resultra/tracker/webui/templatePage"
	"resultra/tracker/webui/userAdmin"

	"resultra/tracker/webui/common/trackerPageContent"

	"resultra/tracker/webui/workspaceAdmin/general"
	"resultra/tracker/webui/workspaceAdmin/users"
)

func init() {

	router := mux.NewRouter()

	itemList.RegisterHTTPHandlers(router)
	itemView.RegisterHTTPHandlers(router)
	designDashboard.RegisterHTTPHandlers(router)
	homePage.RegisterHTTPHandlers(router)
	templatePage.RegisterHTTPHandlers(router)
	admin.RegisterHTTPHandlers(router)
	formPage.RegisterHTTPHandlers(router)
	mainWindow.RegisterHTTPHandlers(router)
	templatePage.RegisterHTTPHandlers(router)
	userAdmin.RegisterHTTPHandlers(router)
	dashboardView.RegisterHTTPHandlers(router)
	alertListView.RegisterHTTPHandlers(router)

	trackerPageContent.RegisterHTTPHandlers(router)

	general.RegisterHTTPHandlers(router)
	users.RegisterHTTPHandlers(router)

	http.Handle("/", router)
}
