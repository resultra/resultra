package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterAPIHTTPHandlers() {

	apiRouter := mux.NewRouter()

	apiRouter.HandleFunc("/api/newLayout", newLayout)
	apiRouter.HandleFunc("/api/newLayoutContainer", newLayoutContainer)
	apiRouter.HandleFunc("/api/resizeLayoutContainer", resizeLayoutContainer)
	apiRouter.HandleFunc("/api/getLayoutContainers", getLayoutContainers)
	apiRouter.HandleFunc("/api/newField", newField)
	apiRouter.HandleFunc("/api/getFieldsByType", getFieldsByType)

	http.Handle("/api/", apiRouter)
}
