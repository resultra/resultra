package form

import (
	"appengine"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/datasheet/server/common/api"
	"resultra/datasheet/server/form"
)

var viewFormTemplates = template.Must(template.ParseFiles(
	"static/common/common.html",
	"static/filter/filterRecords.html",
	"static/form/viewForm.html"))

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
	layoutRef, getErr := form.GetLayoutRef(appEngContext, form.GetLayoutParams{layoutID})
	if getErr != nil {
		api.WriteErrorResponse(w, getErr)
		return
	}

	templParams := ViewFormTemplateParams{"View Form", layoutID, layoutRef.Layout.Name}
	err := viewFormTemplates.ExecuteTemplate(w, "viewForm", templParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
