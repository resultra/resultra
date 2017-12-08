package formLink

import (
	"database/sql"
	"fmt"
	"net/http"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/common/userAuth"
	"resultra/datasheet/server/trackerDatabase"
	"resultra/datasheet/server/userRole"
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

func saveNewFormLink(destDBHandle *sql.DB, newLink FormLink) error {

	encodedProps, encodeErr := generic.EncodeJSONString(newLink.Properties)
	if encodeErr != nil {
		return fmt.Errorf("saveNewFormLink: failure encoding properties: error = %v", encodeErr)
	}

	if _, insertErr := destDBHandle.Exec(`INSERT INTO form_links 
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

func newFormLink(trackerDBHandle *sql.DB, params NewFormLinkParams) (*FormLink, error) {

	newProps := newDefaultNewItemProperties()

	newLink := FormLink{
		LinkID:            uniqueID.GenerateSnowflakeID(),
		Name:              params.Name,
		FormID:            params.FormID,
		IncludeInSidebar:  params.IncludeInSidebar,
		SharedLinkEnabled: false,
		SharedLinkID:      FormLinkDisabledSharedLink,
		Properties:        newProps}

	if saveErr := saveNewFormLink(trackerDBHandle, newLink); saveErr != nil {
		return nil, fmt.Errorf("newFormLink: %v", saveErr)
	}

	return &newLink, nil
}

type GetFormLinkParams struct {
	FormLinkID string `json:"formLinkID"`
}

func GetFormLink(trackerDBHandle *sql.DB, linkID string) (*FormLink, error) {

	formLink := FormLink{}
	encodedProps := ""
	getErr := trackerDBHandle.QueryRow(
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

func GetFormLinkFromSharedLinkID(trackerDBHandle *sql.DB, sharedLinkID string) (*FormLink, error) {

	formLink := FormLink{}
	encodedProps := ""
	getErr := trackerDBHandle.QueryRow(
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

func getAllFormLinksFromSrc(srcDBHandle *sql.DB, parentDatabaseID string) ([]FormLink, error) {

	rows, queryErr := srcDBHandle.Query(
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
	defer rows.Close()

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

func getAllFormLinks(trackerDBHandle *sql.DB, parentDatabaseID string) ([]FormLink, error) {
	return getAllFormLinksFromSrc(trackerDBHandle, parentDatabaseID)
}

func getUserSortedFormLinks(trackerDBHandle *sql.DB, req *http.Request, databaseID string) ([]FormLink, error) {

	allLinks, err := getAllFormLinks(trackerDBHandle, databaseID)
	if err != nil {
		return nil, fmt.Errorf("getUserSortedFormLinks: %v", err)
	}

	if userRole.CurrUserIsDatabaseAdmin(req, databaseID) {
		return allLinks, nil
	}

	currUserID, userErr := userAuth.GetCurrentUserID(req)
	if userErr != nil {
		return nil, fmt.Errorf("getUserSortedFormLinks: %v", userErr)
	}

	visibleLinks, privsErr := userRole.GetNewItemLinksWithUserPrivs(trackerDBHandle, databaseID, currUserID)
	if privsErr != nil {
		return nil, fmt.Errorf("getUserSortedFormLinks: %v", privsErr)
	}

	userLinks := []FormLink{}
	for _, currLink := range allLinks {
		_, foundVisiblePriv := visibleLinks[currLink.LinkID]
		if foundVisiblePriv {
			userLinks = append(userLinks, currLink)
		}
	}

	return userLinks, nil

}

func sortFormLinksByManualOrder(unorderedLinks []FormLink, manualOrder []string) []FormLink {
	// Map the listID -> ListInfo.
	infoByID := map[string]FormLink{}
	for _, currInfo := range unorderedLinks {
		infoByID[currInfo.LinkID] = currInfo
	}
	// Iterate throught the manually ordered list of ListIDs, pull items from listInfoByID in
	// the order they are encountered in the ordered list, then re-append the ListInfo's into a
	// new ordered list in the same order they are found.
	orderedInfo := []FormLink{}
	for _, currID := range manualOrder {
		linkInfo, foundInfo := infoByID[currID]
		if foundInfo {
			orderedInfo = append(orderedInfo, linkInfo)
			delete(infoByID, currID)
		}
	}
	for _, currInfo := range infoByID {
		orderedInfo = append(orderedInfo, currInfo)
	}
	return orderedInfo

}

func getAllSortedFormLinks(trackerDBHandle *sql.DB, parentDatabaseID string) ([]FormLink, error) {

	unsortedLinks, err := getAllFormLinks(trackerDBHandle, parentDatabaseID)
	if err != nil {
		return nil, fmt.Errorf("GetAllSortedFormLinks: %v", err)
	}

	db, getErr := trackerDatabase.GetDatabase(trackerDBHandle, parentDatabaseID)
	if getErr != nil {
		return nil, fmt.Errorf("getDatabaseInfo: Unable to get existing database: %v", getErr)
	}

	sortedLinks := sortFormLinksByManualOrder(unsortedLinks, db.Properties.FormLinkOrder)

	return sortedLinks, nil

}

func updateExistingFormLink(trackerDBHandle *sql.DB, updatedFormLink *FormLink) (*FormLink, error) {

	encodedProps, encodeErr := generic.EncodeJSONString(updatedFormLink.Properties)
	if encodeErr != nil {
		return nil, fmt.Errorf("updateExistingFormLink: failure encoding properties: error = %v", encodeErr)
	}

	if _, updateErr := trackerDBHandle.Exec(`UPDATE form_links 
				SET properties=$1,name=$2,include_in_sidebar=$3,shared_link_enabled=$4,shared_link_id=$5,form_id=$6
				WHERE link_id=$7`,
		encodedProps,
		updatedFormLink.Name,
		updatedFormLink.IncludeInSidebar,
		updatedFormLink.SharedLinkEnabled,
		updatedFormLink.SharedLinkID,
		updatedFormLink.FormID,
		updatedFormLink.LinkID); updateErr != nil {
		return nil, fmt.Errorf("updateExistingFormLink: Can't update form link properties %v: error = %v",
			updatedFormLink.LinkID, updateErr)
	}

	return updatedFormLink, nil

}

func CloneFormLinks(cloneParams *trackerDatabase.CloneDatabaseParams) error {

	links, err := getAllFormLinksFromSrc(cloneParams.SrcDBHandle, cloneParams.SourceDatabaseID)
	if err != nil {
		return fmt.Errorf("CloneFormLinks: Error getting form links for parent database ID = %v: %v",
			cloneParams.SourceDatabaseID, err)
	}

	for _, currLink := range links {

		destLink := currLink

		destLinkID, err := cloneParams.IDRemapper.AllocNewRemappedID(currLink.LinkID)
		if err != nil {
			return fmt.Errorf("CloneTableForms: %v", err)
		}
		destLink.LinkID = destLinkID

		destFormID, err := cloneParams.IDRemapper.GetExistingRemappedID(currLink.FormID)
		if err != nil {
			return fmt.Errorf("CloneTableForms: %v", err)
		}
		destLink.FormID = destFormID

		// If there is a shared link, it must be replaced with a new ID around which to uniquely
		// link to the form. In other words, each clone must have unique links to its forms.
		if len(destLink.SharedLinkID) > 0 {
			destLink.SharedLinkID = uniqueID.GenerateSnowflakeID()
		}

		destProps, err := currLink.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneFormLinks: %v", err)
		}
		destLink.Properties = *destProps

		if err := saveNewFormLink(cloneParams.DestDBHandle, destLink); err != nil {
			return fmt.Errorf("CloneFormLinks: %v", err)
		}

	}

	return nil

}
