package webui

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type LayoutPageInfo struct {
	Title    string
	LayoutID string
}

func designForm(w http.ResponseWriter, r *http.Request) {
	//	c := appengine.NewContext(r)

	vars := mux.Vars(r)
	layoutID := vars["layoutID"]
	log.Println("Design Form: editing for layout with ID = ", layoutID)

	p := LayoutPageInfo{"Edit Layout", layoutID}
	err := htmlTemplates.ExecuteTemplate(w, "designForm", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
