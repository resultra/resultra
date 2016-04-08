package webui

import (
	"appengine"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"resultra/datasheet/controller"
	"resultra/datasheet/datamodel"
)

type FormElemTemplateParams struct {
	ElemPrefix string
}

type DesignFormTemplateParams struct {
	Title          string
	LayoutID       string
	LayoutName     string
	CheckboxParams FormElemTemplateParams
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

	templParams := DesignFormTemplateParams{
		Title:          "Edit Layout",
		LayoutID:       layoutID,
		LayoutName:     layoutRef.Layout.Name,
		CheckboxParams: FormElemTemplateParams{ElemPrefix: "checkbox_"}}

	err := htmlTemplates.ExecuteTemplate(w, "designForm", templParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
