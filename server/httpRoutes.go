package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/dashboard"
	"resultra/datasheet/server/recordFilter"
	"resultra/datasheet/server/recordUpdate"
)

func RegisterAPIHTTPHandlers() {

	apiRouter := mux.NewRouter()

	dashboard.RegisterHTTPHandlers(apiRouter)
	recordUpdate.RegisterHTTPHandlers(apiRouter)
	recordFilter.RegisterHTTPHandlers(apiRouter)

	http.Handle("/api/", apiRouter)
}
