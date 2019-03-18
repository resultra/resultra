package note

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/tracker/server/displayTable/columns/common"
	"resultra/tracker/server/field"
	"resultra/tracker/server/generic"
	"resultra/tracker/server/generic/uniqueID"
	"resultra/tracker/server/trackerDatabase"
)

const noteEntityKind string = "note"

type Note struct {
	ParentTableID string         `json:"parentTableID"`
	NoteID        string         `json:"noteID"`
	ColumnID      string         `json:"columnID"`
	ColType       string         `json:"colType"`
	Properties    NoteProperties `json:"properties"`
}

type NewNoteParams struct {
	ParentTableID string `json:"parentTableID"`
	FieldID       string `json:"fieldID"`
}

func validNoteFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeLongText {
		return true
	} else {
		return false
	}
}

func saveNote(destDBHandle *sql.DB, newNote Note) error {

	if saveErr := common.SaveNewTableColumn(destDBHandle, noteEntityKind,
		newNote.ParentTableID, newNote.NoteID, newNote.Properties); saveErr != nil {
		return fmt.Errorf("saveNewNote: Unable to save html editor: error = %v", saveErr)
	}
	return nil

}

func saveNewNote(trackerDBHandle *sql.DB, params NewNoteParams) (*Note, error) {

	if fieldErr := field.ValidateField(trackerDBHandle, params.FieldID, validNoteFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewCheckBox: %v", fieldErr)
	}

	properties := newDefaultEditorProperties()
	properties.FieldID = params.FieldID

	noteID := uniqueID.GenerateUniqueID()
	newNote := Note{ParentTableID: params.ParentTableID,
		NoteID:     noteID,
		ColumnID:   noteID,
		ColType:    noteEntityKind,
		Properties: properties}

	if err := saveNote(trackerDBHandle, newNote); err != nil {
		return nil, fmt.Errorf("saveNewNote: Unable to save html editor with params=%+v: error = %v", params, err)
	}

	log.Printf("INFO: API: New Note: Created new html editor container: %+v", newNote)

	return &newNote, nil

}

func getNote(trackerDBHandle *sql.DB, parentTableID string, noteID string) (*Note, error) {

	editorProps := newDefaultEditorProperties()
	if getErr := common.GetTableColumn(trackerDBHandle, noteEntityKind, parentTableID, noteID, &editorProps); getErr != nil {
		return nil, fmt.Errorf("getNote: Unable to retrieve html editor: %v", getErr)
	}

	note := Note{
		ParentTableID: parentTableID,
		NoteID:        noteID,
		ColumnID:      noteID,
		ColType:       noteEntityKind,
		Properties:    editorProps}

	return &note, nil
}

func getNotesFromSrc(srcDBHandle *sql.DB, parentTableID string) ([]Note, error) {

	notes := []Note{}

	addEditor := func(editorID string, encodedProps string) error {
		editorProps := newDefaultEditorProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &editorProps); decodeErr != nil {
			return fmt.Errorf("GetNotes: can't decode properties: %v", encodedProps)
		}

		currEditor := Note{
			ParentTableID: parentTableID,
			NoteID:        editorID,
			ColumnID:      editorID,
			ColType:       noteEntityKind,
			Properties:    editorProps}

		notes = append(notes, currEditor)

		return nil
	}
	if getErr := common.GetTableColumns(srcDBHandle, noteEntityKind, parentTableID, addEditor); getErr != nil {
		return nil, fmt.Errorf("GetNotes: Can't get html editors: %v")
	}

	return notes, nil

}

func GetNotes(trackerDBHandle *sql.DB, parentTableID string) ([]Note, error) {
	return getNotesFromSrc(trackerDBHandle, parentTableID)
}

func CloneNotes(cloneParams *trackerDatabase.CloneDatabaseParams, parentTableID string) error {

	srcNotes, err := getNotesFromSrc(cloneParams.SrcDBHandle, parentTableID)
	if err != nil {
		return fmt.Errorf("CloneHTMLEditors: %v", err)
	}

	for _, srcNote := range srcNotes {
		remappedNoteID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcNote.NoteID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcNote.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneHTMLEditors: %v", err)
		}
		destProperties, err := srcNote.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneHTMLEditors: %v", err)
		}
		destNote := Note{
			ParentTableID: remappedFormID,
			NoteID:        remappedNoteID,
			ColumnID:      remappedNoteID,
			ColType:       noteEntityKind,
			Properties:    *destProperties}
		if err := saveNote(cloneParams.DestDBHandle, destNote); err != nil {
			return fmt.Errorf("CloneHTMLEditors: %v", err)
		}
	}

	return nil
}

func updateExistingNote(trackerDBHandle *sql.DB, noteID string, updatedNote *Note) (*Note, error) {

	if updateErr := common.UpdateTableColumn(trackerDBHandle, noteEntityKind, updatedNote.ParentTableID,
		updatedNote.NoteID, updatedNote.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingNote: error updating existing date editor: %v", updateErr)
	}

	return updatedNote, nil

}
