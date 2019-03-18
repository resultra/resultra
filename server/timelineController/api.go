// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package timelineController

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/tracker/server/generic/api"
)

type DummyStructForInclude struct {
	Val int64
}

func init() {
	timelineRouter := mux.NewRouter()

	timelineRouter.HandleFunc("/api/timeline/saveFieldComment", saveFieldCommentAPI)
	timelineRouter.HandleFunc("/api/timeline/getFieldComments", getFieldCommentsAPI)
	timelineRouter.HandleFunc("/api/timeline/getTimelineInfo", getFieldTimelineInfoAPI)

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

func getFieldTimelineInfoAPI(w http.ResponseWriter, r *http.Request) {

	var params GetFieldTimelineInfoParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	timelineInfo, err := getFieldTimelineInfo(r, params)
	if err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, timelineInfo)
	}

}
