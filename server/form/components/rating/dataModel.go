package rating

import (
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
)

const ratingEntityKind string = "rating"

type RatingProperties struct {
	ComponentLink common.ComponentLink           `json:"componentLink"`
	Geometry      componentLayout.LayoutGeometry `json:"geometry"`
}

type Rating struct {
	ParentFormID string           `json:"parentID"`
	RatingID     string           `json:"ratingID"`
	Properties   RatingProperties `json:"properties"`
}

type NewRatingParams struct {
	ParentFormID  string                         `json:"parentFormID"`
	ComponentLink common.ComponentLink           `json:"componentLink"`
	Geometry      componentLayout.LayoutGeometry `json:"geometry"`
}

func validRatingFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeNumber {
		return true
	} else {
		return false
	}
}

func saveNewRating(params NewRatingParams) (*Rating, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if compLinkErr := common.ValidateComponentLink(params.ComponentLink, validRatingFieldType); compLinkErr != nil {
		return nil, fmt.Errorf("saveNewRating: %v", compLinkErr)
	}

	properties := RatingProperties{
		ComponentLink: params.ComponentLink,
		Geometry:      params.Geometry}

	newRating := Rating{ParentFormID: params.ParentFormID,
		RatingID:   uniqueID.GenerateSnowflakeID(),
		Properties: properties}

	if saveErr := common.SaveNewFormComponent(ratingEntityKind,
		newRating.ParentFormID, newRating.RatingID, newRating.Properties); saveErr != nil {
		return nil, fmt.Errorf("saveNewRating: Unable to save rating with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: New Rating: Created new rating component:  %+v", newRating)

	return &newRating, nil

}

func getRating(parentFormID string, ratingID string) (*Rating, error) {

	ratingProps := RatingProperties{}
	if getErr := common.GetFormComponent(ratingEntityKind, parentFormID, ratingID, &ratingProps); getErr != nil {
		return nil, fmt.Errorf("getRating: Unable to retrieve rating: %v", getErr)
	}

	rating := Rating{
		ParentFormID: parentFormID,
		RatingID:     ratingID,
		Properties:   ratingProps}

	return &rating, nil
}

func GetRatings(parentFormID string) ([]Rating, error) {

	ratings := []Rating{}
	addRating := func(ratingID string, encodedProps string) error {

		var ratingProps RatingProperties
		if decodeErr := generic.DecodeJSONString(encodedProps, &ratingProps); decodeErr != nil {
			return fmt.Errorf("GetRatings: can't decode properties: %v", encodedProps)
		}

		currRating := Rating{
			ParentFormID: parentFormID,
			RatingID:     ratingID,
			Properties:   ratingProps}
		ratings = append(ratings, currRating)

		return nil
	}
	if getErr := common.GetFormComponents(ratingEntityKind, parentFormID, addRating); getErr != nil {
		return nil, fmt.Errorf("GetRatings: Can't get ratings: %v")
	}

	return ratings, nil
}

func updateExistingRating(updatedRating *Rating) (*Rating, error) {

	if updateErr := common.UpdateFormComponent(ratingEntityKind, updatedRating.ParentFormID,
		updatedRating.RatingID, updatedRating.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingRating: failure updating rating: %v", updateErr)
	}
	return updatedRating, nil

}
