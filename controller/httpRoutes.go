package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterAPIHTTPHandlers() {

	apiRouter := mux.NewRouter()

	apiRouter.HandleFunc("/api/newLayout", newLayout)
	apiRouter.HandleFunc("/api/newLayoutContainer", newLayoutContainer)

	http.Handle("/api/", apiRouter)
}
