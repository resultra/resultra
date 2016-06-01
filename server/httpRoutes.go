package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/recordFilter"
)

func RegisterAPIHTTPHandlers() {

	apiRouter := mux.NewRouter()

	recordFilter.RegisterHTTPHandlers(apiRouter)

	http.Handle("/api/", apiRouter)
}
