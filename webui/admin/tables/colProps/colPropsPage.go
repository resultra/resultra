package colProps

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/server/displayTable"
	colCommon "resultra/datasheet/server/displayTable/columns/common"
	adminCommon "resultra/datasheet/webui/admin/common"
	"resultra/datasheet/webui/admin/common/inputProperties"

	"resultra/datasheet/server/common/runtimeConfig"

	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/thirdParty"
)

var tablePropTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/admin/tables/colProps/colPropsPage.html",
		"static/admin/tables/colProps/numberInput.html",
		"static/admin/tables/colProps/textInput.html",
		"static/admin/tables/colProps/datePicker.html",
		"static/admin/tables/colProps/checkBox.html",
		"static/admin/tables/colProps/rating.html",
		"static/admin/tables/colProps/toggle.html",
		"static/admin/tables/colProps/userSelection.html",
		"static/admin/tables/colProps/formButton.html",
		"static/admin/tables/colProps/attachment.html",
		"static/admin/tables/colProps/note.html",
		"static/admin/tables/colProps/comment.html",
		"static/admin/tables/colProps/progress.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		adminCommon.TemplateFileList,
		common.TemplateFileList,
		inputProperties.TemplateFileList}

	tablePropTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type TemplParams struct {
	ElemPrefix          string
	Title               string
	DatabaseID          string
	DatabaseName        string
	TableID             string
	TableName           string
	ColID               string
	ColType             string
	ColName             string
	SiteBaseURL         string
	NumberInputParams   NumberInputColPropsTemplateParams
	TextInputParams     TextInputColPropsTemplateParams
	DatePickerParams    DatePickerColPropsTemplateParams
	CheckBoxParams      CheckBoxColPropsTemplateParams
	RatingParams        RatingColPropsTemplateParams
	ToggleParams        ToggleColPropsTemplateParams
	UserSelectionParams UserSelectionColPropsTemplateParams
	FormButtonParams    FormButtonColPropsTemplateParams
	AttachmentParams    AttachmentColPropsTemplateParams
	NoteParams          NoteColPropsTemplateParams
	CommentParams       CommentColPropsTemplateParams
	ProgressParams      ProgressColPropsTemplateParams
}

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/tablecol/{colID}", editPropsPage)
}

func editPropsPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	colID := vars["colID"]

	colInfo, err := colCommon.GetTableColumnInfo(colID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tableInfo, err := displayTable.GetTable(colInfo.TableID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	dbInfo, err := databaseController.GetDatabaseInfo(tableInfo.ParentDatabaseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	elemPrefix := "colProps_"

	templParams := TemplParams{
		ElemPrefix:          elemPrefix,
		Title:               "Column properties",
		DatabaseID:          dbInfo.DatabaseID,
		DatabaseName:        dbInfo.DatabaseName,
		TableID:             colInfo.TableID,
		TableName:           tableInfo.Name,
		ColID:               colID,
		ColType:             colInfo.ColType,
		ColName:             "TBD",
		SiteBaseURL:         runtimeConfig.GetSiteBaseURL(),
		NumberInputParams:   newNumberInputTemplateParams(),
		TextInputParams:     newTextInputTemplateParams(),
		DatePickerParams:    newDatePickerTemplateParams(),
		CheckBoxParams:      newCheckBoxTemplateParams(),
		RatingParams:        newRatingTemplateParams(),
		ToggleParams:        newToggleTemplateParams(),
		UserSelectionParams: newUserSelectionTemplateParams(),
		FormButtonParams:    newFormButtonTemplateParams(),
		AttachmentParams:    newAttachmentTemplateParams(),
		NoteParams:          newNoteTemplateParams(),
		CommentParams:       newCommentTemplateParams(),
		ProgressParams:      newProgressTemplateParams()}

	if err := tablePropTemplates.ExecuteTemplate(w, "colPropsAdminPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
