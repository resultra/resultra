// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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
	"resultra/tracker/webui/setupPage"
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
	setupPage.RegisterHTTPHandlers(router)
	dashboardView.RegisterHTTPHandlers(router)
	alertListView.RegisterHTTPHandlers(router)

	trackerPageContent.RegisterHTTPHandlers(router)

	general.RegisterHTTPHandlers(router)
	users.RegisterHTTPHandlers(router)

	http.Handle("/", router)
}
