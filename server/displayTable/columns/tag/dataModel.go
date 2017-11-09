package tag

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/datasheet/server/displayTable/columns/common"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/trackerDatabase"
)

const tagEntityKind string = "tags"

type Tag struct {
	ParentTableID string        `json:"parentTableID"`
	TagID         string        `json:"tagID"`
	ColumnID      string        `json:"columnID"`
	ColType       string        `json:"colType"`
	Properties    TagProperties `json:"properties"`
}

type NewTagParams struct {
	ParentTableID string `json:"parentTableID"`
	FieldID       string `json:"fieldID"`
}

func validTagFieldType(fieldType string) bool {
	if fieldType == field.FieldTypeLabel {
		return true
	} else {
		return false
	}
}

func saveTag(destDBHandle *sql.DB, newTag Tag) error {
	if saveErr := common.SaveNewTableColumn(destDBHandle, tagEntityKind,
		newTag.ParentTableID, newTag.TagID, newTag.Properties); saveErr != nil {
		return fmt.Errorf("saveNewTag: Unable to save tag: error = %v", saveErr)
	}
	return nil
}

func saveNewTag(trackerDBHandle *sql.DB, params NewTagParams) (*Tag, error) {

	if fieldErr := field.ValidateField(trackerDBHandle, params.FieldID, validTagFieldType); fieldErr != nil {
		return nil, fmt.Errorf("saveNewTag: %v", fieldErr)
	}

	properties := newDefaultTagProperties()
	properties.FieldID = params.FieldID

	tagID := uniqueID.GenerateSnowflakeID()
	newTag := Tag{ParentTableID: params.ParentTableID,
		TagID:      tagID,
		ColumnID:   tagID,
		ColType:    tagEntityKind,
		Properties: properties}

	if saveErr := saveTag(trackerDBHandle, newTag); saveErr != nil {
		return nil, fmt.Errorf("saveNewTag: Unable to save tag with params=%+v: error = %v", params, saveErr)
	}

	log.Printf("INFO: API: New Tag: Created new tag component:  %+v", newTag)

	return &newTag, nil

}

func getTag(trackerDBHandle *sql.DB, parentTableID string, tagID string) (*Tag, error) {

	tagProps := newDefaultTagProperties()
	if getErr := common.GetTableColumn(trackerDBHandle, tagEntityKind, parentTableID,
		tagID, &tagProps); getErr != nil {
		return nil, fmt.Errorf("getTag: Unable to retrieve tag: %v", getErr)
	}

	tag := Tag{
		ParentTableID: parentTableID,
		TagID:         tagID,
		ColumnID:      tagID,
		ColType:       tagEntityKind,
		Properties:    tagProps}

	return &tag, nil
}

func getTagsFromSrc(srcDBHandle *sql.DB, parentTableID string) ([]Tag, error) {

	tags := []Tag{}
	addTag := func(tagID string, encodedProps string) error {

		tagProps := newDefaultTagProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &tagProps); decodeErr != nil {
			return fmt.Errorf("GetTags: can't decode properties: %v", encodedProps)
		}

		currTag := Tag{
			ParentTableID: parentTableID,
			TagID:         tagID,
			ColumnID:      tagID,
			ColType:       tagEntityKind,
			Properties:    tagProps}
		tags = append(tags, currTag)

		return nil
	}
	if getErr := common.GetTableColumns(srcDBHandle, tagEntityKind, parentTableID, addTag); getErr != nil {
		return nil, fmt.Errorf("GetTags: Can't get tags: %v")
	}

	return tags, nil
}

func GetTags(trackerDBHandle *sql.DB, parentTableID string) ([]Tag, error) {
	return getTagsFromSrc(trackerDBHandle, parentTableID)
}

func CloneTags(cloneParams *trackerDatabase.CloneDatabaseParams, parentTableID string) error {

	srcTags, err := getTagsFromSrc(cloneParams.SrcDBHandle, parentTableID)
	if err != nil {
		return fmt.Errorf("CloneTags: %v", err)
	}

	for _, srcTag := range srcTags {
		remappedTagID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcTag.TagID)
		remappedFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcTag.ParentTableID)
		if err != nil {
			return fmt.Errorf("CloneTags: %v", err)
		}
		destProperties, err := srcTag.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneTags: %v", err)
		}
		destTag := Tag{
			ParentTableID: remappedFormID,
			TagID:         remappedTagID,
			ColumnID:      remappedTagID,
			ColType:       tagEntityKind,
			Properties:    *destProperties}
		if err := saveTag(cloneParams.DestDBHandle, destTag); err != nil {
			return fmt.Errorf("CloneTags: %v", err)
		}
	}

	return nil
}

func updateExistingTag(trackerDBHandle *sql.DB, updatedTag *Tag) (*Tag, error) {

	if updateErr := common.UpdateTableColumn(trackerDBHandle, tagEntityKind, updatedTag.ParentTableID,
		updatedTag.TagID, updatedTag.Properties); updateErr != nil {
		return nil, fmt.Errorf("updateExistingTag: failure updating tag: %v", updateErr)
	}
	return updatedTag, nil

}
