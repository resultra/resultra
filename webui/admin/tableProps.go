package admin

import (
	"fmt"
	"html/template"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

var tablePropsTemplates = template.Must(template.ParseFiles("static/admin/tableProps.html"))

func tableProps(w http.ResponseWriter, r *http.Request) {
	// TODO:
	//	databaseID := vars["databaseID"]

	api.WriteErrorResponse(w, fmt.Errorf("Not implemented yet - Table properties page is not implemented"))

}
