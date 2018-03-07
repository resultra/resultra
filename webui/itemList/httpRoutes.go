package itemList

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/itemList/contentLayout", listViewContentLayout)
	mainRouter.HandleFunc("/itemList/propertySidebarContent", propertySidebarContent)
	mainRouter.HandleFunc("/itemList/offPageContent", listViewOffPageContent)

	// TODO - Add handling of get item list info method.
	//	mainRouter.HandleFunc("/viewList/{listID}", ViewList)
}
