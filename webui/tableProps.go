package webui

import (
	// TODO:	"github.com/gorilla/mux"
	"appengine"
	"net/http"
	"resultra/datasheet/controller"
	"resultra/datasheet/datamodel"
)

type TablePropsPageInfo struct {
	Title      string
	DatabaseID string
	FieldRefs  []datamodel.FieldRef
	LayoutRefs []datamodel.LayoutRef
}

func tableProps(w http.ResponseWriter, r *http.Request) {
	// TODO:
	//	databaseID := vars["databaseID"]

	appEngCntxt := appengine.NewContext(r)
	fieldRefs, fieldErr := datamodel.GetAllFieldRefs(appEngCntxt)
	if fieldErr != nil {
		controller.WriteErrorResponse(w, fieldErr)
		return
	}

	layoutRefs, layoutErr := datamodel.GetAllLayoutRefs(appEngCntxt)
	if layoutErr != nil {
		controller.WriteErrorResponse(w, layoutErr)
		return
	}

	p := TablePropsPageInfo{"Database Table Properties", "dummyDatabaseID", fieldRefs, layoutRefs}
	templErr := htmlTemplates.ExecuteTemplate(w, "tableProps", p)
	if templErr != nil {
		controller.WriteErrorResponse(w, templErr)
	}

}
