package webui

import (
	"appengine"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"resultra/datasheet/controller"
	"resultra/datasheet/datamodel"
)

type LayoutPageInfo struct {
	Title      string
	LayoutID   string
	LayoutName string
}

func designForm(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	layoutID := vars["layoutID"]
	log.Println("Design Form: editing for layout with ID = ", layoutID)

	appEngContext := appengine.NewContext(r)
	layoutRef, getErr := datamodel.GetLayoutRef(appEngContext, datamodel.GetLayoutParams{layoutID})
	if getErr != nil {
		controller.WriteErrorResponse(w, getErr)
		return
	}

	p := LayoutPageInfo{"Edit Layout", layoutID, layoutRef.Layout.Name}
	err := htmlTemplates.ExecuteTemplate(w, "designForm", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
