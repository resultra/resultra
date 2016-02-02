package controller

import (
	"net/http"
)

func RegisterAPIHTTPHandlers() {
	http.HandleFunc("/api/newLayout", newLayout)
}
