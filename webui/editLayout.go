package webui

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var editLayoutTemplates = template.Must(template.ParseFiles("template/editLayout.html"))

type LayoutPageInfo struct {
	Title    string
	LayoutID string
}

func editLayout(w http.ResponseWriter, r *http.Request) {
	//	c := appengine.NewContext(r)

	vars := mux.Vars(r)
	layoutID := vars["layoutID"]
	log.Println("editLayout: editing for layout with ID = ", layoutID)

	p := LayoutPageInfo{"Edit Layout", layoutID}
	err := editLayoutTemplates.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
