package design

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/frm/designPageContent/{formID}", designFormPageContent)
	mainRouter.HandleFunc("/admin/frm/offPageContent/{formID}", designFormOffpageContent)
	mainRouter.HandleFunc("/admin/frm/sidebarContent/{formID}", designFormSidebarContent)
}
