package progress

import (
	"database/sql"
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic/numberFormat"
)

type ProgressIDInterface interface {
	getProgressID() string
	getParentFormID() string
}

type ProgressIDHeader struct {
	ParentFormID string `json:"parentFormID"`
	ProgressID   string `json:"progressID"`
}

func (idHeader ProgressIDHeader) getProgressID() string {
	return idHeader.ProgressID
}

func (idHeader ProgressIDHeader) getParentFormID() string {
	return idHeader.ParentFormID
}

type ProgressPropUpdater interface {
	ProgressIDInterface
	updateProps(progress *Progress) error
}

func updateProgressProps(trackerDBHandle *sql.DB, propUpdater ProgressPropUpdater) (*Progress, error) {

	// Retrieve the bar chart from the data store
	progressForUpdate, getErr := getProgress(trackerDBHandle, propUpdater.getParentFormID(), propUpdater.getProgressID())
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

type ProgressResizeParams struct {
	ProgressIDHeader
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (updateParams ProgressResizeParams) updateProps(progress *Progress) error {

	if !componentLayout.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set progress indicator dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	progress.Properties.Geometry = updateParams.Geometry

	return nil
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

type ProgressVisibilityParams struct {
	ProgressIDHeader
	common.ComponentVisibilityProperties
}

func (updateParams ProgressVisibilityParams) updateProps(progress *Progress) error {

	// TODO - Validate conditions

	progress.Properties.VisibilityConditions = updateParams.VisibilityConditions

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
