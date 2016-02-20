package webui

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type EditRecordTemplateParams struct {
	Title    string
	LayoutID string
}

func viewForm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// TODO - Verify layoutID and recordID are valid and have
	// been passed to this function.
	layoutID := vars["layoutID"]
	log.Println("editRecord: editing record: layout ID = %v", layoutID)

	templParams := EditRecordTemplateParams{"View Form", layoutID}
	err := htmlTemplates.ExecuteTemplate(w, "viewForm", templParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
