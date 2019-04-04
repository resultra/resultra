// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package webui

import (
	"github.com/gorilla/mux"
	"github.com/resultra/resultra/webui/admin"
	"github.com/resultra/resultra/webui/alertListView"
	designDashboard "github.com/resultra/resultra/webui/dashboard/design"
	dashboardView "github.com/resultra/resultra/webui/dashboard/view"
	"github.com/resultra/resultra/webui/formPage"
	"github.com/resultra/resultra/webui/homePage"
	"github.com/resultra/resultra/webui/itemList"
	"github.com/resultra/resultra/webui/itemView"
	"github.com/resultra/resultra/webui/mainWindow"
	"github.com/resultra/resultra/webui/setupPage"
	"github.com/resultra/resultra/webui/templatePage"
	"github.com/resultra/resultra/webui/userAdmin"
	"net/http"

	"github.com/resultra/resultra/webui/common/trackerPageContent"

	"github.com/resultra/resultra/webui/workspaceAdmin/general"
	"github.com/resultra/resultra/webui/workspaceAdmin/users"
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
