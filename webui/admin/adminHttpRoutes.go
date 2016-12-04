package admin

import (
	"github.com/gorilla/mux"

	"resultra/datasheet/webui/admin/itemList/itemListProps"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/{databaseID}", adminPage)

	itemListProps.RegisterHTTPHandlers(mainRouter)
}
