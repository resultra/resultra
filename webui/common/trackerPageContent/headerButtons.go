package trackerPageContent

import (
	"github.com/gorilla/mux"

	"html/template"
	"net/http"
	"resultra/datasheet/server/common/userAuth"
	"resultra/datasheet/server/userRole"
	"resultra/datasheet/webui/common/alert"
	"resultra/datasheet/webui/generic"
)

var headerTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/common/trackerPageContent/headerButtons.html"}

	templateFileLists := [][]string{baseTemplateFiles,
		alert.TemplateFileList}

	headerTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type HeaderTemplParams struct {
	CurrUserIsAdmin bool
	DatabaseID      string
}

func headerButtonsContent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	databaseID := vars["databaseID"]

	_, authErr := userAuth.GetCurrentUserInfo(r)
	if authErr != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	isAdmin := userRole.CurrUserIsDatabaseAdmin(r, databaseID)

	templParams := HeaderTemplParams{
		CurrUserIsAdmin: isAdmin,
		DatabaseID:      databaseID}

	if err := headerTemplates.ExecuteTemplate(w, "trackerHeaderButtons", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
