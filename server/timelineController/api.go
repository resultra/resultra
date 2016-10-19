package timelineController

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

type DummyStructForInclude struct {
	Val int64
}

func init() {
	timelineRouter := mux.NewRouter()

	timelineRouter.HandleFunc("/api/timeline/saveFieldComment", saveFieldCommentAPI)
	timelineRouter.HandleFunc("/api/timeline/getFieldComments", getFieldCommentsAPI)

	http.Handle("/api/timeline/", timelineRouter)
}

func saveFieldCommentAPI(w http.ResponseWriter, r *http.Request) {

	params := SaveFieldCommentParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	newComment, err := saveTimelineComment(r, params)
	if err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *newComment)
	}

}

func getFieldCommentsAPI(w http.ResponseWriter, r *http.Request) {

	params := GetFieldRecordCommentInfoParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	comments, err := getFieldRecordTimelineCommentInfo(r, params)
	if err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, comments)
	}

}
