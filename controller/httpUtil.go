package controller

import (
	"log"
	"net/http"
)

func WriteErrorResponse(w http.ResponseWriter, err error) {
	// TBD - Also log the error somewhere
	log.Printf("ERROR: Couldn't process server request: %v", err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
