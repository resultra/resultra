package alert

import (
	"database/sql"
	"fmt"
	"regexp"
	"resultra/tracker/server/field"
	"resultra/tracker/server/trackerDatabase"
)

type identReplacementMap map[string]string

var msgTemplateFieldRefPattern = regexp.MustCompile(`\[[^\]]+\]`)

func replaceTemplateMsgIdents(templateMsg string, replMap identReplacementMap) (string, error) {

	mapIdentVal := func(s []byte) []byte {
		keyVal := string(s[1 : len(s)-1])
		val, foundVal := replMap[keyVal]
		if foundVal {
			return []byte(`[` + val + `]`)
		} else {
			return s
		}
	}

	mappedTemplateMsg := msgTemplateFieldRefPattern.ReplaceAllFunc([]byte(templateMsg), mapIdentVal)

	return string(mappedTemplateMsg), nil

}

func replaceFieldRefWithFieldID(trackerDBHandle *sql.DB, templateMsg string, databaseID string) (string, error) {

	fieldRefIndex, indexErr := field.GetFieldRefIDIndex(trackerDBHandle,
		field.GetFieldListParams{ParentDatabaseID: databaseID})
	if indexErr != nil {
		return "", fmt.Errorf("replaceFieldRefWithFieldID: %v", indexErr)
	}

	fieldRefFieldIDMap := identReplacementMap{}
	for fieldRefName, currField := range fieldRefIndex.FieldsByRefName {
		fieldRefFieldIDMap[fieldRefName] = currField.FieldID
	}

	return replaceTemplateMsgIdents(templateMsg, fieldRefFieldIDMap)

}

func replaceFieldIDWithFieldRef(trackerDBHandle *sql.DB, templateMsg string, databaseID string) (string, error) {

	fieldRefIndex, indexErr := field.GetFieldRefIDIndex(trackerDBHandle,
		field.GetFieldListParams{ParentDatabaseID: databaseID})
	if indexErr != nil {
		return "", fmt.Errorf("replaceFieldRefWithFieldID: %v", indexErr)
	}

	fieldIDFieldRefMap := identReplacementMap{}
	for fieldRefName, currField := range fieldRefIndex.FieldsByRefName {
		fieldIDFieldRefMap[currField.FieldID] = fieldRefName
	}

	return replaceTemplateMsgIdents(templateMsg, fieldIDFieldRefMap)

}

func cloneAlertCaptionMsg(cloneParams *trackerDatabase.CloneDatabaseParams, captionMsg string) string {

	mapFieldIDVal := func(s []byte) []byte {
		srcFieldID := string(s[1 : len(s)-1])
		remappedFieldID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(srcFieldID)
		return []byte(`[` + remappedFieldID + `]`)
	}

	remappedCaptionMsg := msgTemplateFieldRefPattern.ReplaceAllFunc([]byte(captionMsg), mapFieldIDVal)
	return string(remappedCaptionMsg)

}
