package admin

import (
	"github.com/gorilla/mux"

	"resultra/tracker/webui/admin/common"

	"resultra/tracker/webui/admin/mainAdminPage"

	"resultra/tracker/webui/admin/general"

	"resultra/tracker/webui/admin/fields/fieldList"
	"resultra/tracker/webui/admin/fields/fieldProps"

	"resultra/tracker/webui/admin/formLink/formLinkList"
	"resultra/tracker/webui/admin/formLink/formLinkProps"

	"resultra/tracker/webui/admin/tables/colProps"
	"resultra/tracker/webui/admin/tables/tableList"
	"resultra/tracker/webui/admin/tables/tableProps"

	"resultra/tracker/webui/admin/forms/design"
	"resultra/tracker/webui/admin/forms/formList"

	"resultra/tracker/webui/admin/itemList/itemListList"
	"resultra/tracker/webui/admin/itemList/itemListProps"

	"resultra/tracker/webui/admin/userRole/userRoleList"
	"resultra/tracker/webui/admin/userRole/userRoleProps"

	"resultra/tracker/webui/admin/valueLists/valueListList"
	"resultra/tracker/webui/admin/valueLists/valueListProps"

	"resultra/tracker/webui/admin/dashboards"

	"resultra/tracker/webui/admin/globals"

	"resultra/tracker/webui/admin/collaborators/collaboratorList"
	"resultra/tracker/webui/admin/collaborators/collaboratorProps"

	"resultra/tracker/webui/admin/alerts/alertList"
	"resultra/tracker/webui/admin/alerts/alertProps"
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
