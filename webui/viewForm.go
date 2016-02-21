package webui

import (
	"appengine"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"resultra/datasheet/controller"
	"resultra/datasheet/datamodel"
)

type ViewFormTemplateParams struct {
	Title      string
	LayoutID   string
	LayoutName string
}

func viewForm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// TODO - Verify layoutID and recordID are valid and have
	// been passed to this function.
	layoutID := vars["layoutID"]
	log.Println("editRecord: editing record: layout ID = %v", layoutID)

	appEngContext := appengine.NewContext(r)
	layoutRef, getErr := datamodel.GetLayoutRef(appEngContext, datamodel.GetLayoutParams{layoutID})
	if getErr != nil {
		controller.WriteErrorResponse(w, getErr)
		return
	}

	templParams := ViewFormTemplateParams{"View Form", layoutID, layoutRef.Layout.Name}
	err := htmlTemplates.ExecuteTemplate(w, "viewForm", templParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
