package webui

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type EditRecordTemplateParams struct {
	Title    string
	LayoutID string
	RecordID string
}

func editRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// TODO - Verify layoutID and recordID are valid and have
	// been passed to this function.
	layoutID := vars["layoutID"]
	recordID := vars["recordID"]
	log.Println("editRecord: editing record: layout ID = %v, record ID = %v", layoutID, recordID)

	templParams := EditRecordTemplateParams{"Edit Record", layoutID, recordID}
	err := htmlTemplates.ExecuteTemplate(w, "editRecord", templParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
