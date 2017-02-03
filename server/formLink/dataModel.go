package formLink

import (
	"fmt"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/databaseWrapper"
	//	"resultra/datasheet/server/generic/stringValidation"
	"resultra/datasheet/server/generic/uniqueID"
)

type NewFormLinkParams struct {
	Name             string `json:"name"`
	FormID           string `json:"formID"`
	IncludeInSidebar bool   `json:"includeInSidebar"`
}

type FormLink struct {
	LinkID           string             `json:"linkID"`
	Name             string             `json:"name"`
	FormID           string             `json:"formID"`
	IncludeInSidebar bool               `json:"includeInSidebar"`
	Properties       FormLinkProperties `json:"properties"`
}

func saveNewFormLink(newPreset FormLink) error {

	encodedProps, encodeErr := generic.EncodeJSONString(newPreset.Properties)
	if encodeErr != nil {
		return fmt.Errorf("savePreset: failure encoding properties: error = %v", encodeErr)
	}

	if _, insertErr := databaseWrapper.DBHandle().Exec(`INSERT INTO form_links 
				(link_id,form_id,name,include_in_sidebar,properties) VALUES ($1,$2,$3,$4,$5)`,
		newPreset.LinkID,
		newPreset.FormID,
		newPreset.Name,
		newPreset.IncludeInSidebar,
		encodedProps); insertErr != nil {
		return fmt.Errorf("savePreset: Can't create preset: error = %v", insertErr)
	}
	return nil

}

func newFormLink(params NewFormLinkParams) (*FormLink, error) {

	newProps := newDefaultNewItemProperties()

	newPreset := FormLink{
		LinkID:           uniqueID.GenerateSnowflakeID(),
		Name:             params.Name,
		FormID:           params.FormID,
		IncludeInSidebar: params.IncludeInSidebar,
		Properties:       newProps}

	if saveErr := saveNewFormLink(newPreset); saveErr != nil {
		return nil, fmt.Errorf("newFormLink: %v", saveErr)
	}

	return &newPreset, nil
}

type GetFormLinkListParams struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
}

func GetFormLink(linkID string) (*FormLink, error) {

	formLink := FormLink{}
	encodedProps := ""
	getErr := databaseWrapper.DBHandle().QueryRow(`SELECT link_id,name,form_id,include_in_sidebar,properties
			FROM form_links WHERE
			link_id=$1 LIMIT 1`, linkID).Scan(&formLink.LinkID,
		&formLink.Name,
		&formLink.FormID,
		&formLink.IncludeInSidebar,
		&encodedProps)
	if getErr != nil {
		return nil, fmt.Errorf("GetFormLink: Unabled to get form link: link ID = %v: datastore err=%v",
			linkID, getErr)
	}

	props := newDefaultNewItemProperties()
	if decodeErr := generic.DecodeJSONString(encodedProps, &props); decodeErr != nil {
		return nil, fmt.Errorf("GetForm: can't decode properties: %v", encodedProps)
	}
	formLink.Properties = props

	return &formLink, nil

}

func getAllFormLinks(parentDatabaseID string) ([]FormLink, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT form_links.link_id,form_links.name,form_links.form_id,
						form_links.include_in_sidebar,form_links.properties
				FROM forms,form_links WHERE 
				forms.database_id=$1 AND form_links.form_id=forms.form_id`,
		parentDatabaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetAllPresets: Failure querying database: %v", queryErr)
	}

	presets := []FormLink{}
	for rows.Next() {
		var currPreset FormLink
		encodedProps := ""

		if scanErr := rows.Scan(&currPreset.LinkID,
			&currPreset.Name,
			&currPreset.FormID,
			&currPreset.IncludeInSidebar, &encodedProps); scanErr != nil {
			return nil, fmt.Errorf("GetAllForms: Failure querying database: %v", scanErr)
		}

		props := newDefaultNewItemProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &props); decodeErr != nil {
			return nil, fmt.Errorf("GetAllPresets: can't decode properties: %v", encodedProps)
		}
		currPreset.Properties = props

		presets = append(presets, currPreset)
	}

	return presets, nil

}

func CloneFormLinks(remappedIDs uniqueID.UniqueIDRemapper, srcParentDatabaseID string) error {

	presets, err := getAllFormLinks(srcParentDatabaseID)
	if err != nil {
		return fmt.Errorf("CloneFormLinks: Error getting presets for parent database ID = %v: %v",
			srcParentDatabaseID, err)
	}

	for _, currPreset := range presets {

		destPreset := currPreset

		destLinkID, err := remappedIDs.AllocNewRemappedID(currPreset.LinkID)
		if err != nil {
			return fmt.Errorf("CloneTableForms: %v", err)
		}
		destPreset.LinkID = destLinkID

		destFormID, err := remappedIDs.GetExistingRemappedID(currPreset.FormID)
		if err != nil {
			return fmt.Errorf("CloneTableForms: %v", err)
		}
		destPreset.FormID = destFormID

		destProps, err := currPreset.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneFormLinks: %v", err)
		}
		destPreset.Properties = *destProps

		if err := saveNewFormLink(destPreset); err != nil {
			return fmt.Errorf("CloneFormLinks: %v", err)
		}

	}

	return nil

}

/*

func updateExistingForm(formID string, updatedForm *Form) (*Form, error) {

	encodedProps, encodeErr := generic.EncodeJSONString(updatedForm.Properties)
	if encodeErr != nil {
		return nil, fmt.Errorf("updateExistingForm: failure encoding properties: error = %v", encodeErr)
	}

	if _, updateErr := databaseWrapper.DBHandle().Exec(`UPDATE forms
				SET properties=$1, name=$2
				WHERE form_id=$3`,
		encodedProps, updatedForm.Name, formID); updateErr != nil {
		return nil, fmt.Errorf("updateExistingForm: Can't update form properties %v: error = %v",
			formID, updateErr)
	}

	return updatedForm, nil

}

*/
