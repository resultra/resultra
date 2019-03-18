// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package progress

import (
	"database/sql"
	"fmt"
	"resultra/tracker/server/form/components/common"
	"resultra/tracker/server/generic/numberFormat"
)

type ProgressIDInterface interface {
	getProgressID() string
	getParentTableID() string
}

type ProgressIDHeader struct {
	ParentTableID string `json:"parentTableID"`
	ProgressID    string `json:"progressID"`
}

func (idHeader ProgressIDHeader) getProgressID() string {
	return idHeader.ProgressID
}

func (idHeader ProgressIDHeader) getParentTableID() string {
	return idHeader.ParentTableID
}

type ProgressPropUpdater interface {
	ProgressIDInterface
	updateProps(progress *Progress) error
}

func updateProgressProps(trackerDBHandle *sql.DB, propUpdater ProgressPropUpdater) (*Progress, error) {

	// Retrieve the bar chart from the data store
	progressForUpdate, getErr := getProgress(trackerDBHandle, propUpdater.getParentTableID(), propUpdater.getProgressID())
	if getErr != nil {
		return nil, fmt.Errorf("updateProgressProps: Unable to get existing progress indicator: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(progressForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateProgressProps: Unable to update existing progress indicator properties: %v", propUpdateErr)
	}

	progress, updateErr := updateExistingProgress(trackerDBHandle, progressForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateProgressProps: Unable to update existing progress indicator properties: datastore update error =  %v", updateErr)
	}

	return progress, nil
}

type SetRangeParams struct {
	ProgressIDHeader
	MinVal float64 `json:"minVal"`
	MaxVal float64 `json:"maxVal"`
}

func (updateParams SetRangeParams) updateProps(progress *Progress) error {

	if updateParams.MaxVal <= updateParams.MinVal {
		return fmt.Errorf("invalid progress indicator range: %v %v", updateParams.MinVal, updateParams.MaxVal)
	}

	progress.Properties.MinVal = updateParams.MinVal
	progress.Properties.MaxVal = updateParams.MaxVal

	return nil
}

type SetThresholdsParams struct {
	ProgressIDHeader
	ThresholdVals []ThresholdValues `json:"thresholdVals"`
}

func (updateParams SetThresholdsParams) updateProps(progress *Progress) error {

	progress.Properties.ThresholdVals = updateParams.ThresholdVals

	return nil
}

type ProgressLabelFormatParams struct {
	ProgressIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams ProgressLabelFormatParams) updateProps(progress *Progress) error {

	// TODO - Validate format is well-formed.

	progress.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}

type ProgressValueFormatParams struct {
	ProgressIDHeader
	ValueFormat numberFormat.NumberFormatProperties `json:"valueFormat"`
}

func (updateParams ProgressValueFormatParams) updateProps(progress *Progress) error {

	progress.Properties.ValueFormat = updateParams.ValueFormat

	return nil
}

type HelpPopupMsgParams struct {
	ProgressIDHeader
	PopupMsg string `json:"popupMsg"`
}

func (updateParams HelpPopupMsgParams) updateProps(progress *Progress) error {

	progress.Properties.HelpPopupMsg = updateParams.PopupMsg

	return nil
}
