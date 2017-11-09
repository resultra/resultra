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

	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/common/runtimeConfig"

	"resultra/datasheet/server/userRole"
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
		"static/admin/tables/colProps/userTag.html",
		"static/admin/tables/colProps/formButton.html",
		"static/admin/tables/colProps/attachment.html",
		"static/admin/tables/colProps/note.html",
		"static/admin/tables/colProps/comment.html",
		"static/admin/tables/colProps/progress.html",
		"static/admin/tables/colProps/socialButton.html",
		"static/admin/tables/colProps/tag.html",
		"static/admin/tables/colProps/emailAddr.html",
		"static/admin/tables/colProps/urlLink.html",
		"static/admin/tables/colProps/file.html",
		"static/admin/tables/colProps/image.html"}

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
	CurrUserIsAdmin     bool
	NumberInputParams   NumberInputColPropsTemplateParams
	TextInputParams     TextInputColPropsTemplateParams
	DatePickerParams    DatePickerColPropsTemplateParams
	CheckBoxParams      CheckBoxColPropsTemplateParams
	RatingParams        RatingColPropsTemplateParams
	ToggleParams        ToggleColPropsTemplateParams
	UserSelectionParams UserSelectionColPropsTemplateParams
	UserTagParams       UserTagColPropsTemplateParams
	FormButtonParams    FormButtonColPropsTemplateParams
	AttachmentParams    AttachmentColPropsTemplateParams
	NoteParams          NoteColPropsTemplateParams
	CommentParams       CommentColPropsTemplateParams
	ProgressParams      ProgressColPropsTemplateParams
	SocialButtonParams  SocialButtonColPropsTemplateParams
	TagParams           TagColPropsTemplateParams
	EmailAddrParams     EmailAddrColPropsTemplateParams
	UrlLinkParams       UrlLinkColPropsTemplateParams
	FileParams          FileColPropsTemplateParams
	ImageParams         ImageColPropsTemplateParams
}

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/tablecol/{colID}", editPropsPage)
}

func editPropsPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	colID := vars["colID"]

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		http.Error(w, dbErr.Error(), http.StatusInternalServerError)
		return
	}

	colInfo, err := colCommon.GetTableColumnInfo(trackerDBHandle, colID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tableInfo, err := displayTable.GetTable(trackerDBHandle, colInfo.TableID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dbInfo, err := databaseController.GetDatabaseInfo(trackerDBHandle, tableInfo.ParentDatabaseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	elemPrefix := "colProps_"
	isAdmin := userRole.CurrUserIsDatabaseAdmin(r, dbInfo.DatabaseID)

	templParams := TemplParams{
		ElemPrefix:          elemPrefix,
		Title:               "Column properties",
		DatabaseID:          dbInfo.DatabaseID,
		DatabaseName:        dbInfo.DatabaseName,
		CurrUserIsAdmin:     isAdmin,
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
		UserTagParams:       newUserTagTemplateParams(),
		FormButtonParams:    newFormButtonTemplateParams(),
		AttachmentParams:    newAttachmentTemplateParams(),
		NoteParams:          newNoteTemplateParams(),
		CommentParams:       newCommentTemplateParams(),
		ProgressParams:      newProgressTemplateParams(),
		SocialButtonParams:  newSocialButtonTemplateParams(),
		TagParams:           newTagTemplateParams(),
		EmailAddrParams:     newEmailAddrTemplateParams(),
		UrlLinkParams:       newUrlLinkTemplateParams(),
		FileParams:          newFileTemplateParams(),
		ImageParams:         newImageTemplateParams()}

	if err := tablePropTemplates.ExecuteTemplate(w, "colPropsAdminPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
