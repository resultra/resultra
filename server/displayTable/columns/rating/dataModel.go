package rating

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/displayTable/columns/common"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/trackerDatabase"
)

const ratingEntityKind string = "rating"

type Rating struct {
	ParentTableID string           `json:"parentTableID"`
	RatingID      string           `json:"ratingID"`
	ColumnID      string           `json:"columnID"`
	ColType       string           `json:"colType"`
	Properties    RatingProperties `json:"properties"`
}

type NewRatingParams struct {
	ParentTableID string `json:"parentTableID"`
	FieldID       string `json:"fieldID"`
}

func validRatingFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeNumber {
		return true
	} else {
		return false
	}
}

func saveRating(destDBHandle *sql.DB, newRating Rating) error {

	if saveErr := common.SaveNewTableColumn(destDBHandle, ratingEntityKind,
		newRating.ParentTableID, newRating.RatingID, newRating.Properties); saveErr != nil {
		return fmt.Errorf("saveRating: Unable to save rating: error = %v", saveErr)
	}

	return nil
}

func saveNewRating(params NewRatingParams) (*Rating, error) {

	if fieldErr := field.ValidateField(params.FieldID, validRatingFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewRating: %v", fieldErr)
	}

	properties := newDefaultRatingProperties()
	properties.FieldID = params.FieldID

	ratingID := uniqueID.GenerateSnowflakeID()
	newRating := Rating{ParentTableID: params.ParentTableID,
		RatingID:   ratingID,
		ColumnID:   ratingID,
		ColType:    ratingEntityKind,
		Properties: properties}

	if saveErr := saveRating(databaseWrapper.DBHandle(), newRating); saveErr != nil {
		return nil, fmt.Errorf("saveNewRating: Unable to save rating with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: New Rating: Created new rating component:  %+v", newRating)

	return &newRating, nil

}

func getRating(parentTableID string, ratingID string) (*Rating, error) {

	ratingProps := newDefaultRatingProperties()
	if getErr := common.GetTableColumn(ratingEntityKind, parentTableID, ratingID, &ratingProps); getErr != nil {
		return nil, fmt.Errorf("getRating: Unable to retrieve rating: %v", getErr)
	}

	rating := Rating{
		ParentTableID: parentTableID,
		RatingID:      ratingID,
		ColumnID:      ratingID,
		ColType:       ratingEntityKind,
		Properties:    ratingProps}

	return &rating, nil
}

func getRatingsFromSrc(srcDBHandle *sql.DB, parentTableID string) ([]Rating, error) {

	ratings := []Rating{}
	addRating := func(ratingID string, encodedProps string) error {

		ratingProps := newDefaultRatingProperties()
		ratingProps.Tooltips = []string{} // Default to empty set of tooltips
		if decodeErr := generic.DecodeJSONString(encodedProps, &ratingProps); decodeErr != nil {
			return fmt.Errorf("GetRatings: can't decode properties: %v", encodedProps)
		}

		currRating := Rating{
			ParentTableID: parentTableID,
			RatingID:      ratingID,
			ColumnID:      ratingID,
			ColType:       ratingEntityKind,
			Properties:    ratingProps}
		ratings = append(ratings, currRating)

		return nil
	}
	if getErr := common.GetTableColumns(srcDBHandle, ratingEntityKind, parentTableID, addRating); getErr != nil {
		return nil, fmt.Errorf("GetRatings: Can't get ratings: %v")
	}

	return ratings, nil
}

func GetRatings(parentTableID string) ([]Rating, error) {
	return getRatingsFromSrc(databaseWrapper.DBHandle(), parentTableID)
}

func CloneRatings(cloneParams *trackerDatabase.CloneDatabaseParams, parentTableID string) error {

	srcRatings, err := getRatingsFromSrc(cloneParams.SrcDBHandle, parentTableID)
	if err != nil {
		return fmt.Errorf("CloneRatings: %v", err)
	}

	for _, srcRating := range srcRatings {
		remappedRatingID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcRating.RatingID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcRating.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneRatings: %v", err)
		}
		destProperties, err := srcRating.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneRatings: %v", err)
		}
		destRating := Rating{
			ParentTableID: remappedFormID,
			RatingID:      remappedRatingID,
			ColumnID:      remappedRatingID,
			ColType:       ratingEntityKind,
			Properties:    *destProperties}
		if err := saveRating(cloneParams.DestDBHandle, destRating); err != nil {
			return fmt.Errorf("CloneRatings: %v", err)
		}
	}

	return nil
}

func updateExistingRating(updatedRating *Rating) (*Rating, error) {

	if updateErr := common.UpdateTableColumn(ratingEntityKind, updatedRating.ParentTableID,
		updatedRating.RatingID, updatedRating.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingRating: failure updating rating: %v", updateErr)
	}
	return updatedRating, nil

}
