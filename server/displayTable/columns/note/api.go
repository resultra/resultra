package note

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	noteRouter := mux.NewRouter()

	noteRouter.HandleFunc("/api/tableView/note/new", newNote)

	noteRouter.HandleFunc("/api/tableView/note/get", getNoteAPI)

	noteRouter.HandleFunc("/api/tableView/note/setLabelFormat", setLabelFormat)
	noteRouter.HandleFunc("/api/tableView/note/setPermissions", setPermissions)
	noteRouter.HandleFunc("/api/tableView/note/setValidation", setValidation)
	noteRouter.HandleFunc("/api/tableView/note/validateInput", validateInputAPI)

	http.Handle("/api/tableView/note/", noteRouter)
}

func newNote(w http.ResponseWriter, r *http.Request) {

	params := NewNoteParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if editorRef, err := saveNewNote(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *editorRef)
	}

}

type GetNoteParams struct {
	ParentTableID string `json:"parentTableID"`
	NoteID        string `json:"noteID"`
}

func getNoteAPI(w http.ResponseWriter, r *http.Request) {

	var params GetNoteParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	numberInput, err := getNote(params.ParentTableID, params.NoteID)
	if err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	api.WriteJSONResponse(w, *numberInput)
}

func validateInputAPI(w http.ResponseWriter, r *http.Request) {

	var params ValidateInputParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	validationResp := validateInput(params)
	api.WriteJSONResponse(w, validationResp)
}

func processNotePropUpdate(w http.ResponseWriter, r *http.Request, propUpdater NotePropUpdater) {
	if checkBoxRef, err := updateNoteProps(propUpdater); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, checkBoxRef)
	}
}

func setLabelFormat(w http.ResponseWriter, r *http.Request) {
	var params EditorLabelFormatParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processNotePropUpdate(w, r, params)
}

func setPermissions(w http.ResponseWriter, r *http.Request) {
	var params EditorPermissionParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processNotePropUpdate(w, r, params)
}

func setValidation(w http.ResponseWriter, r *http.Request) {
	var params EditorValidationParams
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}
	processNotePropUpdate(w, r, params)
}
