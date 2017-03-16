package rating

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/form/components/common"
)

type RatingIDInterface interface {
	getRatingID() string
	getParentFormID() string
}

type RatingIDHeader struct {
	ParentFormID string `json:"parentFormID"`
	RatingID     string `json:"ratingID"`
}

func (idHeader RatingIDHeader) getRatingID() string {
	return idHeader.RatingID
}

func (idHeader RatingIDHeader) getParentFormID() string {
	return idHeader.ParentFormID
}

type RatingPropUpdater interface {
	RatingIDInterface
	updateProps(rating *Rating) error
}

func updateRatingProps(propUpdater RatingPropUpdater) (*Rating, error) {

	// Retrieve the bar chart from the data store
	ratingForUpdate, getErr := getRating(propUpdater.getParentFormID(), propUpdater.getRatingID())
	if getErr != nil {
		return nil, fmt.Errorf("updateRatingProps: Unable to get existing rating: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(ratingForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateRatingProps: Unable to update existing rating properties: %v", propUpdateErr)
	}

	updatedRating, updateErr := updateExistingRating(ratingForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateRatingProps: Unable to update existing rating properties: datastore update error =  %v", updateErr)
	}

	return updatedRating, nil
}

type RatingResizeParams struct {
	RatingIDHeader
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (updateParams RatingResizeParams) updateProps(rating *Rating) error {

	if !componentLayout.ValidGeometry(updateParams.Geometry) {
		return fmt.Errorf("set check box dimensions: Invalid geometry: %+v", updateParams.Geometry)
	}

	rating.Properties.Geometry = updateParams.Geometry

	return nil
}

type RatingTooltipParams struct {
	RatingIDHeader
	Tooltips []string `json:"tooltips"`
}

func (updateParams RatingTooltipParams) updateProps(rating *Rating) error {

	rating.Properties.Tooltips = updateParams.Tooltips

	return nil
}

type RatingIconParams struct {
	RatingIDHeader
	Icon string `json:"icon"`
}

func (updateParams RatingIconParams) updateProps(rating *Rating) error {

	// TODO - Validate icon is a valid name

	rating.Properties.Icon = updateParams.Icon

	return nil
}

type RatingLabelFormatParams struct {
	RatingIDHeader
	LabelFormat common.ComponentLabelFormatProperties `json:"labelFormat"`
}

func (updateParams RatingLabelFormatParams) updateProps(rating *Rating) error {

	// TODO - Validate format is well-formed.

	rating.Properties.LabelFormat = updateParams.LabelFormat

	return nil
}
