package progress

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
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

func updateProgressProps(propUpdater ProgressPropUpdater) (*Progress, error) {

	// Retrieve the bar chart from the data store
	progressForUpdate, getErr := getProgress(propUpdater.getParentFormID(), propUpdater.getProgressID())
	if getErr != nil {
		return nil, fmt.Errorf("updateProgressProps: Unable to get existing progress indicator: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(progressForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateProgressProps: Unable to update existing progress indicator properties: %v", propUpdateErr)
	}

	progress, updateErr := updateExistingProgress(progressForUpdate)
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
