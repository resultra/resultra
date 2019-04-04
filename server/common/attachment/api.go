// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package attachment

import (
	"github.com/gorilla/mux"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/generic/api"
	"net/http"
)

type DummyStructForInclude struct{ Val int64 }

func init() {

	attachmentRouter := mux.NewRouter()

	attachmentRouter.HandleFunc("/api/attachment/upload", uploadAttachmentAPI)
	attachmentRouter.HandleFunc("/api/attachment/saveURL", saveURLAPI)
	attachmentRouter.HandleFunc("/api/attachment/get/{databaseID}/{cloudFileName}", getAttachmentAPI)
	attachmentRouter.HandleFunc("/api/attachment/getReferences", getAttachmentReferencesAPI)
	attachmentRouter.HandleFunc("/api/attachment/getReference", getAttachmentReferenceAPI)

	attachmentRouter.HandleFunc("/api/attachment/setCaption", setCaptionAPI)
	attachmentRouter.HandleFunc("/api/attachment/setTitle", setTitleAPI)

	http.Handle("/api/attachment/", attachmentRouter)
}

func uploadAttachmentAPI(w http.ResponseWriter, req *http.Request) {

	if uploadResponse, uploadErr := uploadAttachment(w, req); uploadErr != nil {
		api.WriteErrorResponse(w, uploadErr)
	} else {
		api.WriteJSONResponse(w, *uploadResponse)
	}

}

func saveURLAPI(w http.ResponseWriter, r *http.Request) {

	var params SaveURLParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	if attachInfo, saveErr := saveURL(r, params); saveErr != nil {
		api.WriteErrorResponse(w, saveErr)
	} else {
		api.WriteJSONResponse(w, *attachInfo)
	}

}

func getAttachmentAPI(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	databaseID := vars["databaseID"]
	cloudFileName := vars["cloudFileName"]

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	origFileName, err := getOrigFilenameFromCloudFileName(trackerDBHandle, cloudFileName)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	w.Header().Add("Content-Disposition", "attachment;filename="+origFileName)

	serveParams := databaseWrapper.ServeAttachmentParams{
		RespWriter:       w,
		HTTPReq:          r,
		ParentDatabaseID: databaseID,
		CloudFileName:    cloudFileName,
	}
	databaseWrapper.ServeAttachment(serveParams)

}

type GetAttachmentReferencesParams struct {
	AttachmentIDs []string `json:"attachmentIDs"`
}

func getAttachmentReferencesAPI(w http.ResponseWriter, r *http.Request) {

	params := GetAttachmentReferencesParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	// TODO - include database ID in request.
	if attachRefs, err := getAttachmentReferences(trackerDBHandle, params.AttachmentIDs); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, attachRefs)
	}

}

type GetAttachmentReferenceParams struct {
	AttachmentID string `json:"attachmentID"`
}

func getAttachmentReferenceAPI(w http.ResponseWriter, r *http.Request) {

	params := GetAttachmentReferenceParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	// TODO - include database ID in request.
	if attachRefs, err := GetAttachmentReference(trackerDBHandle, params.AttachmentID); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, attachRefs)
	}

}

func setCaptionAPI(w http.ResponseWriter, r *http.Request) {

	params := SetCaptionParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	// TODO - include database ID in request.

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if attachInfo, err := setCaption(trackerDBHandle, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, attachInfo)
	}

}

func setTitleAPI(w http.ResponseWriter, r *http.Request) {

	params := SetTitleParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	if attachInfo, err := setTitle(trackerDBHandle, params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, attachInfo)
	}

}
