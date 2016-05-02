package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/dashboard"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/record"
	"resultra/datasheet/server/recordFilter"
	"resultra/datasheet/server/recordUpdate"
)

func RegisterAPIHTTPHandlers() {

	apiRouter := mux.NewRouter()

	field.RegisterHTTPHandlers(apiRouter)
	dashboard.RegisterHTTPHandlers(apiRouter)
	record.RegisterHTTPHandlers(apiRouter)
	recordUpdate.RegisterHTTPHandlers(apiRouter)
	recordFilter.RegisterHTTPHandlers(apiRouter)

	http.Handle("/api/", apiRouter)
}
