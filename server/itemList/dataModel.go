package itemList

import (
	"fmt"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/databaseWrapper"
	"resultra/datasheet/server/generic/stringValidation"
	"resultra/datasheet/server/generic/uniqueID"
)

type ItemList struct {
	ListID           string             `json:"listID"`
	ParentDatabaseID string             `json:"parentDatabaseID"`
	FormID           string             `json:"formID"`
	Name             string             `json:"name"`
	Properties       ItemListProperties `json:"properties"`
}

func saveItemList(newList ItemList) error {
	encodedProps, encodeErr := generic.EncodeJSONString(newList.Properties)
	if encodeErr != nil {
		return fmt.Errorf("saveItemList: failure encoding properties: error = %v", encodeErr)
	}

	if _, insertErr := databaseWrapper.DBHandle().Exec(`INSERT INTO item_lists (database_id,list_id,form_id,name,properties) 
					VALUES ($1,$2,$3,$4,$5)`,
		newList.ParentDatabaseID, newList.ListID, newList.FormID, newList.Name, encodedProps); insertErr != nil {
		return fmt.Errorf("saveItemList: Can't create list: error = %v", insertErr)
	}
	return nil

}

type NewItemListParams struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
	FormID           string `json:"formID"`
	Name             string `json:"name"`
}

func newItemList(params NewItemListParams) (*ItemList, error) {

	sanitizedName, sanitizeErr := stringValidation.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	newList := ItemList{ParentDatabaseID: params.ParentDatabaseID,
		FormID:     params.FormID,
		ListID:     uniqueID.GenerateSnowflakeID(),
		Name:       sanitizedName,
		Properties: newDefaultItemListProperties()}

	if err := saveItemList(newList); err != nil {
		return nil, fmt.Errorf("newItemList: error saving list: %v", err)
	}

	return &newList, nil
}

func GetItemList(listID string) (*ItemList, error) {

	listName := ""
	encodedProps := ""
	databaseID := ""
	formID := ""
	getErr := databaseWrapper.DBHandle().QueryRow(`SELECT database_id,form_id,name,properties FROM item_lists
		 WHERE list_id=$1 LIMIT 1`, listID).Scan(&databaseID, &formID, &listName, &encodedProps)
	if getErr != nil {
		return nil, fmt.Errorf("GetItemList: Unabled to get item list: list ID = %v: datastore err=%v",
			listID, getErr)
	}

	listProps := newDefaultItemListProperties()
	if decodeErr := generic.DecodeJSONString(encodedProps, &listProps); decodeErr != nil {
		return nil, fmt.Errorf("GetItemList: can't decode properties: %v", encodedProps)
	}

	retrievedList := ItemList{
		ParentDatabaseID: databaseID,
		FormID:           formID,
		ListID:           listID,
		Name:             listName,
		Properties:       listProps}

	return &retrievedList, nil
}

func GetAllItemLists(parentDatabaseID string) ([]ItemList, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT database_id,list_id,form_id,name,properties FROM item_lists WHERE database_id = $1`,
		parentDatabaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetAllItemLists: Failure querying database: %v", queryErr)
	}

	itemLists := []ItemList{}
	for rows.Next() {
		var currList ItemList
		encodedProps := ""

		if scanErr := rows.Scan(&currList.ParentDatabaseID, &currList.ListID, &currList.FormID, &currList.Name, &encodedProps); scanErr != nil {
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

func CloneItemLists(remappedIDs uniqueID.UniqueIDRemapper, srcParentDatabaseID string) error {

	remappedDatabaseID, err := remappedIDs.GetExistingRemappedID(srcParentDatabaseID)
	if err != nil {
		return fmt.Errorf("CloneTableLists: Error getting remapped database ID: %v", err)
	}

	lists, err := GetAllItemLists(srcParentDatabaseID)
	if err != nil {
		return fmt.Errorf("CloneTableLists: $v", err)
	}

	for _, currList := range lists {

		destList := currList
		destList.ParentDatabaseID = remappedDatabaseID

		destListID, err := remappedIDs.AllocNewRemappedID(currList.ListID)
		if err != nil {
			return fmt.Errorf("CloneTableForms: %v", err)
		}
		destList.ListID = destListID

		destFormID, err := remappedIDs.GetExistingRemappedID(currList.FormID)
		if err != nil {
			return fmt.Errorf("CloneTableForms: %v", err)
		}
		destList.FormID = destFormID

		destProps, err := currList.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneTableLists: %v", err)
		}
		destList.Properties = *destProps

		if err := saveItemList(destList); err != nil {
			return fmt.Errorf("CloneTableLists: %v", err)
		}

	}

	return nil

}

func updateExistingItemList(listID string, updatedItemList *ItemList) (*ItemList, error) {

	encodedProps, encodeErr := generic.EncodeJSONString(updatedItemList.Properties)
	if encodeErr != nil {
		return nil, fmt.Errorf("updateExistingItemList: failure encoding properties: error = %v", encodeErr)
	}

	if _, updateErr := databaseWrapper.DBHandle().Exec(`UPDATE item_lists 
				SET properties=$1, name=$2,form_id=$3
				WHERE list_id=$4`,
		encodedProps, updatedItemList.Name, updatedItemList.FormID, listID); updateErr != nil {
		return nil, fmt.Errorf("updateExistingItemList: Can't update form properties %v: error = %v",
			listID, updateErr)
	}

	return updatedItemList, nil

}

func GetItemListDatabaseID(listID string) (string, error) {

	theList, err := GetItemList(listID)
	if err != nil {
		return "", err
	}
	return theList.ParentDatabaseID, nil
}
