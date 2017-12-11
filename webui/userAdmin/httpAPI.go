package userAdmin

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/resetPassword/{resetID}", resetPassword)
	mainRouter.HandleFunc("/register/{inviteID}", registerNewUser)

}
