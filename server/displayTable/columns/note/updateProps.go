// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package note

import (
	"database/sql"
	"fmt"
	"github.com/resultra/resultra/server/form/components/common"
)

type NoteIDInterface interface {
	getNoteID() string
	getParentTableID() string
}

type NoteIDHeader struct {
	NoteID        string `json:"noteID"`
	ParentTableID string `json:"parentTableID"`
}

func (idHeader NoteIDHeader) getNoteID() string {
	return idHeader.NoteID
}

func (idHeader NoteIDHeader) getParentTableID() string {
	return idHeader.ParentTableID
}

type NotePropUpdater interface {
	NoteIDInterface
	updateProps(note *Note) error
}

func updateNoteProps(trackerDBHandle *sql.DB, propUpdater NotePropUpdater) (*Note, error) {

	// Retrieve the bar chart from the data store
	noteForUpdate, getErr := getNote(trackerDBHandle, propUpdater.getParentTableID(), propUpdater.getNoteID())
	if getErr != nil {
		return nil, fmt.Errorf("updateNoteProps: Unable to get existing html editor: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(noteForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateNoteProps: Unable to update existing html editor properties: %v", propUpdateErr)
	}

	note, updateErr := updateExistingNote(trackerDBHandle, propUpdater.getNoteID(), noteForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateNoteProps: Unable to update existing html editor properties: datastore update error =  %v", updateErr)
	}

	return note, nil
}

type EditorLabelFormatParams struct {
	NoteIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams EditorLabelFormatParams) updateProps(note *Note) error {

	// TODO - Validate format is well-formed.

	note.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}

type EditorPermissionParams struct {
	NoteIDHeader
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
}

func (updateParams EditorPermissionParams) updateProps(editor *Note) error {

	editor.Properties.Permissions = updateParams.Permissions

	return nil
}

type EditorValidationParams struct {
	NoteIDHeader
	Validation ValidationProperties `json:"validation"`
}

func (updateParams EditorValidationParams) updateProps(editor *Note) error {

	editor.Properties.Validation = updateParams.Validation

	return nil
}

type HelpPopupMsgParams struct {
	NoteIDHeader
	PopupMsg string `json:"popupMsg"`
}

func (updateParams HelpPopupMsgParams) updateProps(editor *Note) error {

	editor.Properties.HelpPopupMsg = updateParams.PopupMsg

	return nil
}
