package comment

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	commentRouter := mux.NewRouter()

	commentRouter.HandleFunc("/api/frm/comment/new", newComment)
	commentRouter.HandleFunc("/api/frm/comment/resize", resizeComment)

	http.Handle("/api/frm/comment/", commentRouter)
}

func newComment(w http.ResponseWriter, r *http.Request) {

	commentParams := NewCommentParams{}
	if err := api.DecodeJSONRequest(r, &commentParams); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if commentRef, err := saveNewComment(commentParams); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *commentRef)
	}

}

func processCommentPropUpdate(w http.ResponseWriter, r *http.Request, propUpdater CommentPropUpdater) {
	if commentRef, err := updateCommentProps(propUpdater); err != nil {
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
