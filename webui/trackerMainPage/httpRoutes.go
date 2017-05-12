package trackerMainPage

import "github.com/gorilla/mux"

func RegisterHTTPHandlers(mainRouter *mux.Router) {

	mainRouter.HandleFunc("/main/{databaseID}", trackerMainPage)

}
