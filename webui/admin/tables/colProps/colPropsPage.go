// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package colProps

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"resultra/tracker/server/databaseController"
	"resultra/tracker/server/displayTable"
	colCommon "resultra/tracker/server/displayTable/columns/common"
	adminCommon "resultra/tracker/webui/admin/common"
	"resultra/tracker/webui/admin/common/inputProperties"

	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/common/runtimeConfig"
	"resultra/tracker/server/workspace"

	"resultra/tracker/server/common/userAuth"
	"resultra/tracker/server/userRole"
	"resultra/tracker/webui/common"
	"resultra/tracker/webui/generic"
	"resultra/tracker/webui/thirdParty"
)

var tablePropTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/admin/tables/colProps/colPropsPage.html",
		"static/admin/tables/colProps/numberInput.html",
		"static/admin/tables/colProps/textInput.html",
		"static/admin/tables/colProps/textSelection.html",
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
	ElemPrefix            string
	Title                 string
	DatabaseID            string
	DatabaseName          string
	WorkspaceName         string
	TableID               string
	TableName             string
	ColID                 string
	ColType               string
	ColName               string
	SiteBaseURL           string
	CurrUserIsAdmin       bool
	IsSingleUserWorkspace bool
	NumberInputParams     NumberInputColPropsTemplateParams
	TextInputParams       TextInputColPropsTemplateParams
	TextSelectionParams   TextSelectionColPropsTemplateParams
	DatePickerParams      DatePickerColPropsTemplateParams
	CheckBoxParams        CheckBoxColPropsTemplateParams
	RatingParams          RatingColPropsTemplateParams
	ToggleParams          ToggleColPropsTemplateParams
	UserSelectionParams   UserSelectionColPropsTemplateParams
	UserTagParams         UserTagColPropsTemplateParams
	FormButtonParams      FormButtonColPropsTemplateParams
	AttachmentParams      AttachmentColPropsTemplateParams
	NoteParams            NoteColPropsTemplateParams
	CommentParams         CommentColPropsTemplateParams
	ProgressParams        ProgressColPropsTemplateParams
	SocialButtonParams    SocialButtonColPropsTemplateParams
	TagParams             TagColPropsTemplateParams
	EmailAddrParams       EmailAddrColPropsTemplateParams
	UrlLinkParams         UrlLinkColPropsTemplateParams
	FileParams            FileColPropsTemplateParams
	ImageParams           ImageColPropsTemplateParams
}

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/tablecol/{colID}", editPropsPage)
}

func editPropsPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	colID := vars["colID"]

	_, authErr := userAuth.GetCurrentUserInfo(r)
	if authErr != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		http.Error(w, dbErr.Error(), http.StatusInternalServerError)
		return
	}

	workspaceName, workspaceErr := workspace.GetWorkspaceName(trackerDBHandle)
	if workspaceErr != nil {
		http.Error(w, workspaceErr.Error(), http.StatusInternalServerError)
		return
	}

	colInfo, err := colCommon.GetTableColumnInfo(trackerDBHandle, colID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
		ElemPrefix:            elemPrefix,
		Title:                 "Column properties",
		DatabaseID:            dbInfo.DatabaseID,
		DatabaseName:          dbInfo.DatabaseName,
		WorkspaceName:         workspaceName,
		IsSingleUserWorkspace: runtimeConfig.CurrRuntimeConfig.IsSingleUserWorkspace,
		CurrUserIsAdmin:       isAdmin,
		TableID:               colInfo.TableID,
		TableName:             tableInfo.Name,
		ColID:                 colID,
		ColType:               colInfo.ColType,
		ColName:               "TBD",
		SiteBaseURL:           runtimeConfig.GetSiteBaseURL(),
		NumberInputParams:     newNumberInputTemplateParams(),
		TextInputParams:       newTextInputTemplateParams(),
		TextSelectionParams:   newTextSelectionTemplateParams(),
		DatePickerParams:      newDatePickerTemplateParams(),
		CheckBoxParams:        newCheckBoxTemplateParams(),
		RatingParams:          newRatingTemplateParams(),
		ToggleParams:          newToggleTemplateParams(),
		UserSelectionParams:   newUserSelectionTemplateParams(),
		UserTagParams:         newUserTagTemplateParams(),
		FormButtonParams:      newFormButtonTemplateParams(),
		AttachmentParams:      newAttachmentTemplateParams(),
		NoteParams:            newNoteTemplateParams(),
		CommentParams:         newCommentTemplateParams(),
		ProgressParams:        newProgressTemplateParams(),
		SocialButtonParams:    newSocialButtonTemplateParams(),
		TagParams:             newTagTemplateParams(),
		EmailAddrParams:       newEmailAddrTemplateParams(),
		UrlLinkParams:         newUrlLinkTemplateParams(),
		FileParams:            newFileTemplateParams(),
		ImageParams:           newImageTemplateParams()}

	if err := tablePropTemplates.ExecuteTemplate(w, "colPropsAdminPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
