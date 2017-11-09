package rating

import (
	"database/sql"
	"fmt"
	"resultra/datasheet/server/form/components/common"
)

type RatingIDInterface interface {
	getRatingID() string
	getParentTableID() string
}

type RatingIDHeader struct {
	ParentTableID string `json:"parentTableID"`
	RatingID      string `json:"ratingID"`
}

func (idHeader RatingIDHeader) getRatingID() string {
	return idHeader.RatingID
}

func (idHeader RatingIDHeader) getParentTableID() string {
	return idHeader.ParentTableID
}

type RatingPropUpdater interface {
	RatingIDInterface
	updateProps(rating *Rating) error
}

func updateRatingProps(trackerDBHandle *sql.DB, propUpdater RatingPropUpdater) (*Rating, error) {

	// Retrieve the bar chart from the data store
	ratingForUpdate, getErr := getRating(trackerDBHandle, propUpdater.getParentTableID(), propUpdater.getRatingID())
	if getErr != nil {
		return nil, fmt.Errorf("updateRatingProps: Unable to get existing rating: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(ratingForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateRatingProps: Unable to update existing rating properties: %v", propUpdateErr)
	}

	updatedRating, updateErr := updateExistingRating(trackerDBHandle, ratingForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateRatingProps: Unable to update existing rating properties: datastore update error =  %v", updateErr)
	}

	return updatedRating, nil
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

type RatingPermissionParams struct {
	RatingIDHeader
	Permissions common.ComponentValuePermissionsProperties `json:"permissions"`
}

func (updateParams RatingPermissionParams) updateProps(rating *Rating) error {

	// TODO - Validate conditions

	rating.Properties.Permissions = updateParams.Permissions

	return nil
}

type RatingValidationParams struct {
	RatingIDHeader
	Validation RatingValidationProperties `json:"validation"`
}

func (updateParams RatingValidationParams) updateProps(rating *Rating) error {

	// TODO - Validate conditions

	rating.Properties.Validation = updateParams.Validation

	return nil
}

type RatingClearValueSupportedParams struct {
	RatingIDHeader
	ClearValueSupported bool `json:"clearValueSupported"`
}

func (updateParams RatingClearValueSupportedParams) updateProps(rating *Rating) error {

	rating.Properties.ClearValueSupported = updateParams.ClearValueSupported

	return nil
}

type HelpPopupMsgParams struct {
	RatingIDHeader
	PopupMsg string `json:"popupMsg"`
}

func (updateParams HelpPopupMsgParams) updateProps(rating *Rating) error {

	rating.Properties.HelpPopupMsg = updateParams.PopupMsg

	return nil
}

type RangeParams struct {
	RatingIDHeader
	MinVal int `json:"minVal"`
	MaxVal int `json:"maxVal"`
}

func (updateParams RangeParams) updateProps(rating *Rating) error {

	if updateParams.MinVal >= updateParams.MaxVal {
		return fmt.Errorf("Error setting rating range: min = %v, max = %v",
			updateParams.MinVal, updateParams.MaxVal)
	}

	rating.Properties.MinVal = updateParams.MinVal
	rating.Properties.MaxVal = updateParams.MaxVal

	return nil
}
