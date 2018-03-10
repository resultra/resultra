package trackerPageContent

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/common/trackerPageContent/toc", tocPageContent)
	mainRouter.HandleFunc("/common/trackerPageContent/headerButtons/{databaseID}", headerButtonsContent)
}
