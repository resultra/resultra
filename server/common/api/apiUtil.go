package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type JSONParams map[string]string

func DecodeJSONRequest(r *http.Request, decodedVal interface{}) error {

	if err := json.NewDecoder(r.Body).Decode(decodedVal); err != nil {
		return fmt.Errorf("DecodeJSONRequest:Error decoding server JSON request: decode error = %v", err)
	} else {
		log.Printf("INFO: API: Decoded JSON: %+v", decodedVal)
		return nil
	}
}

func WriteJSONResponse(w http.ResponseWriter, responseVals interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encodeErr := json.NewEncoder(w).Encode(responseVals)
	if encodeErr != nil {
		WriteErrorResponse(w, encodeErr)
	}
}
