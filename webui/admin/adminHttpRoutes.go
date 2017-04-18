package admin

import (
	"github.com/gorilla/mux"

	"resultra/datasheet/webui/admin/fields/fieldProps"
	"resultra/datasheet/webui/admin/formLink/formLinkProps"
	"resultra/datasheet/webui/admin/forms/design"
	"resultra/datasheet/webui/admin/itemList/itemListProps"
	"resultra/datasheet/webui/admin/userRole/userRoleProps"
	"resultra/datasheet/webui/admin/valueLists/valueListProps"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/{databaseID}", adminPage)

	design.RegisterHTTPHandlers(mainRouter)

	itemListProps.RegisterHTTPHandlers(mainRouter)
	formLinkProps.RegisterHTTPHandlers(mainRouter)
	userRoleProps.RegisterHTTPHandlers(mainRouter)
	fieldProps.RegisterHTTPHandlers(mainRouter)
	valueListProps.RegisterHTTPHandlers(mainRouter)
}
