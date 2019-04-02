// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package admin

import (
	"github.com/gorilla/mux"

	"github.com/resultra/resultra/webui/admin/common"

	"github.com/resultra/resultra/webui/admin/mainAdminPage"

	"github.com/resultra/resultra/webui/admin/general"

	"github.com/resultra/resultra/webui/admin/fields/fieldList"
	"github.com/resultra/resultra/webui/admin/fields/fieldProps"

	"github.com/resultra/resultra/webui/admin/formLink/formLinkList"
	"github.com/resultra/resultra/webui/admin/formLink/formLinkProps"

	"github.com/resultra/resultra/webui/admin/tables/colProps"
	"github.com/resultra/resultra/webui/admin/tables/tableList"
	"github.com/resultra/resultra/webui/admin/tables/tableProps"

	"github.com/resultra/resultra/webui/admin/forms/design"
	"github.com/resultra/resultra/webui/admin/forms/formList"

	"github.com/resultra/resultra/webui/admin/itemList/itemListList"
	"github.com/resultra/resultra/webui/admin/itemList/itemListProps"

	"github.com/resultra/resultra/webui/admin/userRole/userRoleList"
	"github.com/resultra/resultra/webui/admin/userRole/userRoleProps"

	"github.com/resultra/resultra/webui/admin/valueLists/valueListList"
	"github.com/resultra/resultra/webui/admin/valueLists/valueListProps"

	"github.com/resultra/resultra/webui/admin/dashboards"

	"github.com/resultra/resultra/webui/admin/globals"

	"github.com/resultra/resultra/webui/admin/collaborators/collaboratorList"
	"github.com/resultra/resultra/webui/admin/collaborators/collaboratorProps"

	"github.com/resultra/resultra/webui/admin/alerts/alertList"
	"github.com/resultra/resultra/webui/admin/alerts/alertProps"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {

	common.RegisterHTTPHandlers(mainRouter)
	mainAdminPage.RegisterHTTPHandlers(mainRouter)

	general.RegisterHTTPHandlers(mainRouter)

	design.RegisterHTTPHandlers(mainRouter)
	formList.RegisterHTTPHandlers(mainRouter)

	tableList.RegisterHTTPHandlers(mainRouter)
	tableProps.RegisterHTTPHandlers(mainRouter)
	colProps.RegisterHTTPHandlers(mainRouter)

	itemListProps.RegisterHTTPHandlers(mainRouter)
	itemListList.RegisterHTTPHandlers(mainRouter)

	formLinkProps.RegisterHTTPHandlers(mainRouter)
	formLinkList.RegisterHTTPHandlers(mainRouter)

	userRoleProps.RegisterHTTPHandlers(mainRouter)
	userRoleList.RegisterHTTPHandlers(mainRouter)

	fieldProps.RegisterHTTPHandlers(mainRouter)
	fieldList.RegisterHTTPHandlers(mainRouter)

	valueListProps.RegisterHTTPHandlers(mainRouter)
	valueListList.RegisterHTTPHandlers(mainRouter)

	dashboards.RegisterHTTPHandlers(mainRouter)

	globals.RegisterHTTPHandlers(mainRouter)

	collaboratorList.RegisterHTTPHandlers(mainRouter)
	collaboratorProps.RegisterHTTPHandlers(mainRouter)

	alertList.RegisterHTTPHandlers(mainRouter)
	alertProps.RegisterHTTPHandlers(mainRouter)
}
