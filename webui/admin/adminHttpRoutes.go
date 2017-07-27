package admin

import (
	"github.com/gorilla/mux"

	"resultra/datasheet/webui/admin/general"

	"resultra/datasheet/webui/admin/fields/fieldList"
	"resultra/datasheet/webui/admin/fields/fieldProps"

	"resultra/datasheet/webui/admin/formLink/formLinkList"
	"resultra/datasheet/webui/admin/formLink/formLinkProps"

	"resultra/datasheet/webui/admin/tables/colProps"
	"resultra/datasheet/webui/admin/tables/tableList"
	"resultra/datasheet/webui/admin/tables/tableProps"

	"resultra/datasheet/webui/admin/forms/design"
	"resultra/datasheet/webui/admin/forms/formList"

	"resultra/datasheet/webui/admin/itemList/itemListList"
	"resultra/datasheet/webui/admin/itemList/itemListProps"

	"resultra/datasheet/webui/admin/userRole/userRoleList"
	"resultra/datasheet/webui/admin/userRole/userRoleProps"

	"resultra/datasheet/webui/admin/valueLists/valueListList"
	"resultra/datasheet/webui/admin/valueLists/valueListProps"

	"resultra/datasheet/webui/admin/dashboards"

	"resultra/datasheet/webui/admin/globals"

	"resultra/datasheet/webui/admin/collaborators/collaboratorList"
	"resultra/datasheet/webui/admin/collaborators/collaboratorProps"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {

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
}
