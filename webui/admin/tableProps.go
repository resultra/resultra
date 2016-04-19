package admin

import (
	"appengine"
	"html/template"
	"net/http"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form"
	"resultra/datasheet/server/generic/api"
)

type TablePropsPageInfo struct {
	Title      string
	DatabaseID string
	FieldRefs  []field.FieldRef
	LayoutRefs []form.LayoutRef
}

var tablePropsTemplates = template.Must(template.ParseFiles("static/admin/tableProps.html"))

func tableProps(w http.ResponseWriter, r *http.Request) {
	// TODO:
	//	databaseID := vars["databaseID"]

	appEngCntxt := appengine.NewContext(r)
	fieldRefs, fieldErr := field.GetAllFieldRefs(appEngCntxt)
	if fieldErr != nil {
		api.WriteErrorResponse(w, fieldErr)
		return
	}

	layoutRefs, layoutErr := form.GetAllLayoutRefs(appEngCntxt)
	if layoutErr != nil {
		api.WriteErrorResponse(w, layoutErr)
		return
	}

	p := TablePropsPageInfo{"Database Table Properties", "dummyDatabaseID", fieldRefs, layoutRefs}
	templErr := tablePropsTemplates.ExecuteTemplate(w, "tableProps", p)
	if templErr != nil {
		api.WriteErrorResponse(w, templErr)
	}

}
