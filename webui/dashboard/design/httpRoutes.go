// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package design

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/dashboard/designDashboardMainContent/{dashboardID}", designDashboardMainContent)
	mainRouter.HandleFunc("/admin/dashboard/designDashboardOffpageContent/{dashboardID}", designDashboardOffpageContent)
	mainRouter.HandleFunc("/admin/dashboard/designDashboardSidebarContent/{dashboardID}", designDashboardSidebarContent)

}
