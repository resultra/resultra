package fieldList

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/fields/mainContent/{databaseID}", fieldListAdminPage)
	mainRouter.HandleFunc("/admin/fields/offPageContent", fieldListOffpageContent)

}
