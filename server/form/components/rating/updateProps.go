package rating

import (
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
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
