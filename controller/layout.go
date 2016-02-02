package controller

import (
	"appengine"
	"encoding/json"
	"log"
	"net/http"
	"resultra/datasheet/datamodel"
)

func newLayout(w http.ResponseWriter, r *http.Request) {

	log.Println("newLayout method:", r.Method) //get request method

	var layoutParam map[string]string
	json.NewDecoder(r.Body).Decode(&layoutParam)
	log.Println("newLayout: New layout parameters:", layoutParam)

	appEngCntxt := appengine.NewContext(r)

	layoutID, err := datamodel.NewLayout(appEngCntxt, layoutParam["name"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"layoutID": layoutID})
}
