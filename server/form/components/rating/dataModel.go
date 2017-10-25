package rating

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/form/components/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/trackerDatabase"
)

const ratingEntityKind string = "rating"

type Rating struct {
	ParentFormID string           `json:"parentFormID"`
	RatingID     string           `json:"ratingID"`
	Properties   RatingProperties `json:"properties"`
}

type NewRatingParams struct {
	ParentFormID string                         `json:"parentFormID"`
	FieldID      string                         `json:"fieldID"`
	Geometry     componentLayout.LayoutGeometry `json:"geometry"`
}

func validRatingFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeNumber {
		return true
	} else {
		return false
	}
}

func saveRating(destDBHandle *sql.DB, newRating Rating) error {

	if saveErr := common.SaveNewFormComponent(destDBHandle, ratingEntityKind,
		newRating.ParentFormID, newRating.RatingID, newRating.Properties); saveErr != nil {
		return fmt.Errorf("saveRating: Unable to save rating: error = %v", saveErr)
	}

	return nil
}

func saveNewRating(params NewRatingParams) (*Rating, error) {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("Invalid layout container parameters: %+v", params)
	}

	if fieldErr := field.ValidateField(params.FieldID, validRatingFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewRating: %v", fieldErr)
	}

	properties := newDefaultRatingProperties()
	properties.FieldID = params.FieldID
	properties.Geometry = params.Geometry

	newRating := Rating{ParentFormID: params.ParentFormID,
		RatingID:   uniqueID.GenerateSnowflakeID(),
		Properties: properties}

	if saveErr := saveRating(databaseWrapper.DBHandle(), newRating); saveErr != nil {
		return nil, fmt.Errorf("saveNewRating: Unable to save rating with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: New Rating: Created new rating component:  %+v", newRating)

	return &newRating, nil

}

func getRating(parentFormID string, ratingID string) (*Rating, error) {

	ratingProps := newDefaultRatingProperties()
	if getErr := common.GetFormComponent(ratingEntityKind, parentFormID, ratingID, &ratingProps); getErr != nil {
		return nil, fmt.Errorf("getRating: Unable to retrieve rating: %v", getErr)
	}

	rating := Rating{
		ParentFormID: parentFormID,
		RatingID:     ratingID,
		Properties:   ratingProps}

	return &rating, nil
}

func getRatingsFromSrc(srcDBHandle *sql.DB, parentFormID string) ([]Rating, error) {

	ratings := []Rating{}
	addRating := func(ratingID string, encodedProps string) error {

		ratingProps := newDefaultRatingProperties()
		ratingProps.Tooltips = []string{} // Default to empty set of tooltips
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
	if getErr := common.GetFormComponents(srcDBHandle, ratingEntityKind, parentFormID, addRating); getErr != nil {
		return nil, fmt.Errorf("GetRatings: Can't get ratings: %v")
	}

	return ratings, nil
}

func GetRatings(parentFormID string) ([]Rating, error) {
	return getRatingsFromSrc(databaseWrapper.DBHandle(), parentFormID)
}

func CloneRatings(cloneParams *trackerDatabase.CloneDatabaseParams, parentFormID string) error {

	srcRatings, err := getRatingsFromSrc(cloneParams.SrcDBHandle, parentFormID)
	if err != nil {
		return fmt.Errorf("CloneRatings: %v", err)
	}

	for _, srcRating := range srcRatings {
		remappedRatingID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcRating.RatingID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcRating.ParentFormID)
		if err != nil {
			return fmt.Errorf("CloneRatings: %v", err)
		}
		destProperties, err := srcRating.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneRatings: %v", err)
		}
		destRating := Rating{
			ParentFormID: remappedFormID,
			RatingID:     remappedRatingID,
			Properties:   *destProperties}
		if err := saveRating(cloneParams.DestDBHandle, destRating); err != nil {
			return fmt.Errorf("CloneRatings: %v", err)
		}
	}

	return nil
}

func updateExistingRating(updatedRating *Rating) (*Rating, error) {

	if updateErr := common.UpdateFormComponent(ratingEntityKind, updatedRating.ParentFormID,
		updatedRating.RatingID, updatedRating.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingRating: failure updating rating: %v", updateErr)
	}
	return updatedRating, nil

}
