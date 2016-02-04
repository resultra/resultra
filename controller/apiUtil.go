package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

func decodeJSONRequest(r *http.Request, decodedVal interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(decodedVal); err != nil {
		return err
	} else {
		return nil
	}
}

func writeJSONResponse(w http.ResponseWriter, responseVals interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encodeErr := json.NewEncoder(w).Encode(responseVals)
	if encodeErr != nil {
		writeErrorResponse(w, encodeErr)
	}
}

func writeErrorResponse(w http.ResponseWriter, err error) {
	// TBD - Also log the error somewhere
	log.Printf("ERROR processing API request: %v", err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
