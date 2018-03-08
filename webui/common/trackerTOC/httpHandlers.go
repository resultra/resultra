package trackerTOC

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/common/trackerTOC/toc", tocPageContent)
	mainRouter.HandleFunc("/common/trackerTOC/headerButtons", headerButtonsContent)
}
