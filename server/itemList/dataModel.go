package itemList

import (
	"database/sql"
	"fmt"
	"net/http"
	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/common/userAuth"
	"resultra/tracker/server/generic"
	"resultra/tracker/server/generic/stringValidation"
	"resultra/tracker/server/generic/uniqueID"
	"resultra/tracker/server/trackerDatabase"
	"resultra/tracker/server/userRole"
)

type ItemList struct {
	ListID           string             `json:"listID"`
	ParentDatabaseID string             `json:"parentDatabaseID"`
	Name             string             `json:"name"`
	Properties       ItemListProperties `json:"properties"`
}

func saveItemList(destDBHandle *sql.DB, newList ItemList) error {
	encodedProps, encodeErr := generic.EncodeJSONString(newList.Properties)
	if encodeErr != nil {
		return fmt.Errorf("saveItemList: failure encoding properties: error = %v", encodeErr)
	}

	if _, insertErr := destDBHandle.Exec(`INSERT INTO item_lists (database_id,list_id,name,properties) 
					VALUES ($1,$2,$3,$4)`,
		newList.ParentDatabaseID, newList.ListID, newList.Name, encodedProps); insertErr != nil {
		return fmt.Errorf("saveItemList: Can't create list: error = %v", insertErr)
	}
	return nil

}

type NewItemListParams struct {
	ParentDatabaseID string                 `json:"parentDatabaseID"`
	DefaultView      ItemListViewProperties `json:"defaultView"`
	Name             string                 `json:"name"`
}

func newItemList(trackerDBHandle *sql.DB, params NewItemListParams) (*ItemList, error) {

	sanitizedName, sanitizeErr := stringValidation.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	props := newDefaultItemListProperties()
	props.DefaultView = params.DefaultView

	newList := ItemList{ParentDatabaseID: params.ParentDatabaseID,
		ListID:     uniqueID.GenerateUniqueID(),
		Name:       sanitizedName,
		Properties: props}

	if err := saveItemList(trackerDBHandle, newList); err != nil {
		return nil, fmt.Errorf("newItemList: error saving list: %v", err)
	}

	return &newList, nil
}

func GetItemList(trackerDBHandle *sql.DB, listID string) (*ItemList, error) {

	listName := ""
	encodedProps := ""
	databaseID := ""
	getErr := trackerDBHandle.QueryRow(`SELECT database_id,name,properties FROM item_lists
		 WHERE list_id=$1 LIMIT 1`, listID).Scan(&databaseID, &listName, &encodedProps)
	if getErr != nil {
		return nil, fmt.Errorf("GetItemList: Unabled to get item list: list ID = %v: datastore err=%v",
			listID, getErr)
	}

	listProps := newDefaultItemListProperties()
	if decodeErr := generic.DecodeJSONString(encodedProps, &listProps); decodeErr != nil {
		return nil, fmt.Errorf("GetItemList: can't decode properties: %v, err = %v", encodedProps, decodeErr)
	}

	retrievedList := ItemList{
		ParentDatabaseID: databaseID,
		ListID:           listID,
		Name:             listName,
		Properties:       listProps}

	return &retrievedList, nil
}

func getAllItemListsFromSrc(srcDBHandle *sql.DB, parentDatabaseID string) ([]ItemList, error) {

	rows, queryErr := srcDBHandle.Query(
		`SELECT database_id,list_id,name,properties FROM item_lists WHERE database_id = $1`,
		parentDatabaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetAllItemLists: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	itemLists := []ItemList{}
	for rows.Next() {
		var currList ItemList
		encodedProps := ""

		if scanErr := rows.Scan(&currList.ParentDatabaseID, &currList.ListID, &currList.Name, &encodedProps); scanErr != nil {
			return nil, fmt.Errorf("GetAllItemLists: Failure querying database: %v", scanErr)
		}

		listProps := newDefaultItemListProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &listProps); decodeErr != nil {
			return nil, fmt.Errorf("GetAllItemLists: can't decode properties: %v", encodedProps)
		}
		currList.Properties = listProps

		itemLists = append(itemLists, currList)
	}

	return itemLists, nil

}

func GetAllItemLists(trackerDBHandle *sql.DB, parentDatabaseID string) ([]ItemList, error) {
	return getAllItemListsFromSrc(trackerDBHandle, parentDatabaseID)
}

func orderListsByManualListOrder(unorderedListInfo []ItemList, manualOrder []string) []ItemList {
	// Map the listID -> ListInfo.
	listInfoByID := map[string]ItemList{}
	for _, currListInfo := range unorderedListInfo {
		listInfoByID[currListInfo.ListID] = currListInfo
	}
	// Iterate throught the manually ordered list of ListIDs, pull items from listInfoByID in
	// the order they are encountered in the ordered list, then re-append the ListInfo's into a
	// new ordered list in the same order they are found.
	orderedListInfo := []ItemList{}
	for _, currListID := range manualOrder {
		listInfo, foundListInfo := listInfoByID[currListID]
		if foundListInfo {
			orderedListInfo = append(orderedListInfo, listInfo)
			delete(listInfoByID, currListID)
		}
	}
	for _, currListInfo := range listInfoByID {
		orderedListInfo = append(orderedListInfo, currListInfo)
	}
	return orderedListInfo

}

func GetAllSortedItemLists(trackerDBHandle *sql.DB, parentDatabaseID string) ([]ItemList, error) {

	unorderedLists, err := GetAllItemLists(trackerDBHandle, parentDatabaseID)
	if err != nil {
		return nil, fmt.Errorf("GetAllSortedItemLists: %v")
	}

	db, getErr := trackerDatabase.GetDatabase(trackerDBHandle, parentDatabaseID)
	if getErr != nil {
		return nil, fmt.Errorf("getDatabaseInfo: Unable to get existing database: %v", getErr)
	}

	orderedLists := orderListsByManualListOrder(unorderedLists, db.Properties.ListOrder)

	return orderedLists, nil
}

func GetAllUserSortedItemLists(req *http.Request, parentDatabaseID string) ([]ItemList, error) {

	currUserID, userErr := userAuth.GetCurrentUserID(req)
	if userErr != nil {
		return nil, fmt.Errorf("getUserDashboards: can't verify user: %v", userErr)
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(req)
	if dbErr != nil {
		return nil, dbErr
	}

	allLists, err := GetAllSortedItemLists(trackerDBHandle, parentDatabaseID)
	if err != nil {
		return nil, fmt.Errorf("GetAllSortedItemLists: can't verify user: %v", err)
	}

	if userRole.CurrUserIsDatabaseAdmin(req, parentDatabaseID) {

		sidebarLists := []ItemList{}
		for _, currList := range allLists {
			if currList.Properties.IncludeInSidebar {
				sidebarLists = append(sidebarLists, currList)
			}
		}
		return sidebarLists, nil
	}

	userListPrivs, userListErr := userRole.GetItemListsWithUserPrivs(trackerDBHandle, parentDatabaseID, currUserID)
	if userListErr != nil {
		return nil, fmt.Errorf("GetAllUserSortedItemLists: %v", userListErr)
	}

	userLists := []ItemList{}
	for _, currList := range allLists {
		_, foundPriv := userListPrivs[currList.ListID]
		if foundPriv && currList.Properties.IncludeInSidebar {
			userLists = append(userLists, currList)
		}
	}

	return userLists, nil

}

func CloneItemLists(cloneParams *trackerDatabase.CloneDatabaseParams) error {

	remappedDatabaseID, err := cloneParams.IDRemapper.GetExistingRemappedID(cloneParams.SourceDatabaseID)
	if err != nil {
		return fmt.Errorf("CloneTableLists: Error getting remapped database ID: %v", err)
	}

	lists, err := getAllItemListsFromSrc(cloneParams.SrcDBHandle, cloneParams.SourceDatabaseID)
	if err != nil {
		return fmt.Errorf("CloneTableLists: $v", err)
	}

	for _, currList := range lists {

		destList := currList
		destList.ParentDatabaseID = remappedDatabaseID

		destListID := cloneParams.IDRemapper.AllocNewOrGetExistingRemappedID(currList.ListID)
		destList.ListID = destListID

		destProps, err := currList.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneTableLists: %v", err)
		}
		destList.Properties = *destProps

		if err := saveItemList(cloneParams.DestDBHandle, destList); err != nil {
			return fmt.Errorf("CloneTableLists: %v", err)
		}

	}

	return nil

}

func updateExistingItemList(trackerDBHandle *sql.DB, listID string, updatedItemList *ItemList) (*ItemList, error) {

	encodedProps, encodeErr := generic.EncodeJSONString(updatedItemList.Properties)
	if encodeErr != nil {
		return nil, fmt.Errorf("updateExistingItemList: failure encoding properties: error = %v", encodeErr)
	}

	if _, updateErr := trackerDBHandle.Exec(`UPDATE item_lists 
				SET properties=$1, name=$2
				WHERE list_id=$3`,
		encodedProps, updatedItemList.Name, listID); updateErr != nil {
		return nil, fmt.Errorf("updateExistingItemList: Can't update form properties %v: error = %v",
			listID, updateErr)
	}

	return updatedItemList, nil

}

func GetItemListDatabaseID(trackerDBHandle *sql.DB, listID string) (string, error) {

	theList, err := GetItemList(trackerDBHandle, listID)
	if err != nil {
		return "", err
	}
	return theList.ParentDatabaseID, nil
}
