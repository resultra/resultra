package attachment

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/generic/cloudStorageWrapper"
)

type DummyStructForInclude struct{ Val int64 }

func init() {

	attachmentRouter := mux.NewRouter()

	attachmentRouter.HandleFunc("/api/attachment/upload", uploadAttachmentAPI)
	attachmentRouter.HandleFunc("/api/attachment/get/{fileName}", getAttachmentAPI)
	attachmentRouter.HandleFunc("/api/attachment/getReferences", getAttachmentReferencesAPI)

	http.Handle("/api/attachment/", attachmentRouter)
}

func uploadAttachmentAPI(w http.ResponseWriter, req *http.Request) {

	if uploadResponse, uploadErr := uploadAttachment(req); uploadErr != nil {
		api.WriteErrorResponse(w, uploadErr)
	} else {
		api.WriteJSONResponse(w, *uploadResponse)
	}

}

func getAttachmentAPI(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	fileName := vars["fileName"]

	http.ServeFile(w, r, cloudStorageWrapper.LocalAttachmentFileUploadDir+fileName)

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

	// TODO - include database ID in request.
	if attachRefs, err := getAttachmentReferences(params.AttachmentIDs); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, attachRefs)
	}

}
