package comment

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/generic/api"
)

func init() {
	commentRouter := mux.NewRouter()

	commentRouter.HandleFunc("/api/tableView/comment/new", newComment)

	commentRouter.HandleFunc("/api/tableView/comment/get", getCommentAPI)

	commentRouter.HandleFunc("/api/tableView/comment/setLabelFormat", setLabelFormat)
	commentRouter.HandleFunc("/api/tableView/comment/setPermissions", setPermissions)
	commentRouter.HandleFunc("/api/tableView/comment/setHelpPopupMsg", setHelpPopupMsg)

	http.Handle("/api/tableView/comment/", commentRouter)
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

type GetCommentParams struct {
	ParentTableID string `json:"parentTableID"`
	CommentID     string `json:"commentID"`
}

func getCommentAPI(w http.ResponseWriter, r *http.Request) {

	var params GetCommentParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	numberInput, err := getComment(trackerDBHandle, params.ParentTableID, params.CommentID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	api.WriteJSONResponse(w, *numberInput)
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

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params CommentLabelFormatParams
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
