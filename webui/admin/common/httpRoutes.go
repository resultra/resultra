package common

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/common/settingsTOC", tocContentLayout)
}
