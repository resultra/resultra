// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package comment

import (
	"github.com/gorilla/mux"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/generic/api"
	"net/http"
)

func init() {
	commentRouter := mux.NewRouter()

	commentRouter.HandleFunc("/api/frm/comment/new", newComment)
	commentRouter.HandleFunc("/api/frm/comment/resize", resizeComment)
	commentRouter.HandleFunc("/api/frm/comment/setLabelFormat", setLabelFormat)
	commentRouter.HandleFunc("/api/frm/comment/setVisibility", setVisibility)
	commentRouter.HandleFunc("/api/frm/comment/setPermissions", setPermissions)
	commentRouter.HandleFunc("/api/frm/comment/setHelpPopupMsg", setHelpPopupMsg)

	http.Handle("/api/frm/comment/", commentRouter)
}

func newComment(w http.ResponseWriter, r *http.Request) {

	commentParams := NewCommentParams{}
	if err := api.DecodeJSONRequest(r, &commentParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if commentRef, err := saveNewComment(trackerDBHandle, commentParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *commentRef)
	}

}

func processCommentPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater CommentPropUpdater) {
	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}
	if commentRef, err := updateCommentProps(trackerDBHandle, propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, commentRef)
	}
}

func resizeComment(w http.ResponseWriter, r *http.Request) {
	var resizeParams CommentResizeParams
	if err := api.DecodeJSONRequest(r, &resizeParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCommentPropUpdate(w, r, resizeParams)
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params CommentLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCommentPropUpdate(w, r, params)
}

func setVisibility(w http.ResponseWriter, r *http.Request) {
	var params CommentVisibilityParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCommentPropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params CommentPermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCommentPropUpdate(w, r, params)
}

func setHelpPopupMsg(w http.ResponseWriter, r *http.Request) {
	var params HelpPopupMsgParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processCommentPropUpdate(w, r, params)
}
