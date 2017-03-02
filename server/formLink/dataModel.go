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
	LinkID            string             `json:"linkID"`
	Name              string             `json:"name"`
	FormID            string             `json:"formID"`
	IncludeInSidebar  bool               `json:"includeInSidebar"`
	SharedLinkEnabled bool               `json:"sharedLinkEnabled"`
	SharedLinkID      string             `json:"sharedLinkID"`
	Properties        FormLinkProperties `json:"properties"`
}

const FormLinkDisabledSharedLink string = ""

func saveNewFormLink(newLink FormLink) error {

	encodedProps, encodeErr := generic.EncodeJSONString(newLink.Properties)
	if encodeErr != nil {
		return fmt.Errorf("saveNewFormLink: failure encoding properties: error = %v", encodeErr)
	}

	if _, insertErr := databaseWrapper.DBHandle().Exec(`INSERT INTO form_links 
				(link_id,form_id,name,include_in_sidebar,shared_link_enabled,shared_link_id,properties) 
				VALUES ($1,$2,$3,$4,$5,$6,$7)`,
		newLink.LinkID,
		newLink.FormID,
		newLink.Name,
		newLink.IncludeInSidebar,
		newLink.SharedLinkEnabled,
		newLink.SharedLinkID,
		encodedProps); insertErr != nil {
		return fmt.Errorf("savePreset: Can't create preset: error = %v", insertErr)
	}
	return nil

}

func newFormLink(params NewFormLinkParams) (*FormLink, error) {

	newProps := newDefaultNewItemProperties()

	newLink := FormLink{
		LinkID:            uniqueID.GenerateSnowflakeID(),
		Name:              params.Name,
		FormID:            params.FormID,
		IncludeInSidebar:  params.IncludeInSidebar,
		SharedLinkEnabled: false,
		SharedLinkID:      FormLinkDisabledSharedLink,
		Properties:        newProps}

	if saveErr := saveNewFormLink(newLink); saveErr != nil {
		return nil, fmt.Errorf("newFormLink: %v", saveErr)
	}

	return &newLink, nil
}

type GetFormLinkParams struct {
	FormLinkID string `json:"formLinkID"`
}

func GetFormLink(linkID string) (*FormLink, error) {

	formLink := FormLink{}
	encodedProps := ""
	getErr := databaseWrapper.DBHandle().QueryRow(
		`SELECT link_id,name,form_id,include_in_sidebar,shared_link_enabled,shared_link_id,properties
			FROM form_links WHERE
			link_id=$1 LIMIT 1`, linkID).Scan(&formLink.LinkID,
		&formLink.Name,
		&formLink.FormID,
		&formLink.IncludeInSidebar,
		&formLink.SharedLinkEnabled,
		&formLink.SharedLinkID,
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

func GetFormLinkFromSharedLinkID(sharedLinkID string) (*FormLink, error) {

	formLink := FormLink{}
	encodedProps := ""
	getErr := databaseWrapper.DBHandle().QueryRow(
		`SELECT link_id,name,form_id,include_in_sidebar,shared_link_enabled,shared_link_id,properties
			FROM form_links WHERE
			shared_link_id=$1 LIMIT 1`, sharedLinkID).Scan(&formLink.LinkID,
		&formLink.Name,
		&formLink.FormID,
		&formLink.IncludeInSidebar,
		&formLink.SharedLinkEnabled,
		&formLink.SharedLinkID,
		&encodedProps)
	if getErr != nil {
		return nil, fmt.Errorf("GetFormLinkFromSharedLinkID: Unabled to get form link: shared link ID = %v: datastore err=%v",
			sharedLinkID, getErr)
	}

	props := newDefaultNewItemProperties()
	if decodeErr := generic.DecodeJSONString(encodedProps, &props); decodeErr != nil {
		return nil, fmt.Errorf("GetForm: can't decode properties: %v", encodedProps)
	}
	formLink.Properties = props

	return &formLink, nil

}

type GetFormLinkListParams struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
}

func getAllFormLinks(parentDatabaseID string) ([]FormLink, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT form_links.link_id,form_links.name,form_links.form_id,
						form_links.include_in_sidebar,
						form_links.shared_link_enabled,
						form_links.shared_link_id,
						form_links.properties
				FROM forms,form_links WHERE 
				forms.database_id=$1 AND form_links.form_id=forms.form_id`,
		parentDatabaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetAllPresets: Failure querying database: %v", queryErr)
	}

	links := []FormLink{}
	for rows.Next() {
		var currLink FormLink
		encodedProps := ""

		if scanErr := rows.Scan(&currLink.LinkID,
			&currLink.Name,
			&currLink.FormID,
			&currLink.IncludeInSidebar,
			&currLink.SharedLinkEnabled,
			&currLink.SharedLinkID,
			&encodedProps); scanErr != nil {
			return nil, fmt.Errorf("GetAllForms: Failure querying database: %v", scanErr)
		}

		props := newDefaultNewItemProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &props); decodeErr != nil {
			return nil, fmt.Errorf("GetAllPresets: can't decode properties: %v", encodedProps)
		}
		currLink.Properties = props

		links = append(links, currLink)
	}

	return links, nil

}

func updateExistingFormLink(updatedFormLink *FormLink) (*FormLink, error) {

	encodedProps, encodeErr := generic.EncodeJSONString(updatedFormLink.Properties)
	if encodeErr != nil {
		return nil, fmt.Errorf("updateExistingFormLink: failure encoding properties: error = %v", encodeErr)
	}

	if _, updateErr := databaseWrapper.DBHandle().Exec(`UPDATE form_links 
				SET properties=$1,name=$2,include_in_sidebar=$3,shared_link_enabled=$4,shared_link_id=$5
				WHERE link_id=$6`,
		encodedProps,
		updatedFormLink.Name,
		updatedFormLink.IncludeInSidebar,
		updatedFormLink.SharedLinkEnabled,
		updatedFormLink.SharedLinkID,
		updatedFormLink.LinkID); updateErr != nil {
		return nil, fmt.Errorf("updateExistingFormLink: Can't update form link properties %v: error = %v",
			updatedFormLink.LinkID, updateErr)
	}

	return updatedFormLink, nil

}

func CloneFormLinks(remappedIDs uniqueID.UniqueIDRemapper, srcParentDatabaseID string) error {

	links, err := getAllFormLinks(srcParentDatabaseID)
	if err != nil {
		return fmt.Errorf("CloneFormLinks: Error getting form links for parent database ID = %v: %v",
			srcParentDatabaseID, err)
	}

	for _, currLink := range links {

		destLink := currLink

		destLinkID, err := remappedIDs.AllocNewRemappedID(currLink.LinkID)
		if err != nil {
			return fmt.Errorf("CloneTableForms: %v", err)
		}
		destLink.LinkID = destLinkID

		destFormID, err := remappedIDs.GetExistingRemappedID(currLink.FormID)
		if err != nil {
			return fmt.Errorf("CloneTableForms: %v", err)
		}
		destLink.FormID = destFormID

		// If there is a shared link, it must be replaced with a new ID around which to uniquely
		// link to the form. In other words, each clone must have unique links to its forms.
		if len(destLink.SharedLinkID) > 0 {
			destLink.SharedLinkID = uniqueID.GenerateSnowflakeID()
		}

		destProps, err := currLink.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneFormLinks: %v", err)
		}
		destLink.Properties = *destProps

		if err := saveNewFormLink(destLink); err != nil {
			return fmt.Errorf("CloneFormLinks: %v", err)
		}

	}

	return nil

}
