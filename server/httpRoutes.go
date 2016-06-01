package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/recordFilter"
	"resultra/datasheet/server/recordUpdate"
)

func RegisterAPIHTTPHandlers() {

	apiRouter := mux.NewRouter()

	recordUpdate.RegisterHTTPHandlers(apiRouter)
	recordFilter.RegisterHTTPHandlers(apiRouter)

	http.Handle("/api/", apiRouter)
}
