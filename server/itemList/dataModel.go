package itemList

import (
	"fmt"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/databaseWrapper"
	"resultra/datasheet/server/generic/stringValidation"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/table"
)

type ItemList struct {
	ListID        string             `json:"listID"`
	ParentTableID string             `json:"parentTableID"`
	FormID        string             `json:"formID"`
	Name          string             `json:"name"`
	Properties    ItemListProperties `json:"properties"`
}

func saveItemList(newList ItemList) error {
	encodedProps, encodeErr := generic.EncodeJSONString(newList.Properties)
	if encodeErr != nil {
		return fmt.Errorf("saveItemList: failure encoding properties: error = %v", encodeErr)
	}

	if _, insertErr := databaseWrapper.DBHandle().Exec(`INSERT INTO item_lists (table_id,list_id,form_id,name,properties) 
					VALUES ($1,$2,$3,$4,$5)`,
		newList.ParentTableID, newList.ListID, newList.FormID, newList.Name, encodedProps); insertErr != nil {
		return fmt.Errorf("saveItemList: Can't create list: error = %v", insertErr)
	}
	return nil

}

type NewItemListParams struct {
	ParentTableID string `json:"parentTableID"`
	FormID        string `json:"formID"`
	Name          string `json:"name"`
}

func newItemList(params NewItemListParams) (*ItemList, error) {

	sanitizedName, sanitizeErr := stringValidation.SanitizeName(params.Name)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	newList := ItemList{ParentTableID: params.ParentTableID,
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
	tableID := ""
	formID := ""
	getErr := databaseWrapper.DBHandle().QueryRow(`SELECT table_id,form_id,name,properties FROM item_lists
		 WHERE list_id=$1 LIMIT 1`, listID).Scan(&tableID, &formID, &listName, &encodedProps)
	if getErr != nil {
		return nil, fmt.Errorf("GetItemList: Unabled to get item list: list ID = %v: datastore err=%v",
			listID, getErr)
	}

	var listProps ItemListProperties
	if decodeErr := generic.DecodeJSONString(encodedProps, &listProps); decodeErr != nil {
		return nil, fmt.Errorf("GetItemList: can't decode properties: %v", encodedProps)
	}

	retrievedList := ItemList{
		ParentTableID: tableID,
		FormID:        formID,
		ListID:        listID,
		Name:          listName,
		Properties:    listProps}

	return &retrievedList, nil
}

func getAllItemLists(parentTableID string) ([]ItemList, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT table_id,list_id,form_id,name,properties FROM item_lists WHERE table_id = $1`,
		parentTableID)
	if queryErr != nil {
		return nil, fmt.Errorf("getAllItemLists: Failure querying database: %v", queryErr)
	}

	itemLists := []ItemList{}
	for rows.Next() {
		var currList ItemList
		encodedProps := ""

		if scanErr := rows.Scan(&currList.ParentTableID, &currList.ListID, &currList.FormID, &currList.Name, &encodedProps); scanErr != nil {
			return nil, fmt.Errorf("getAllItemLists: Failure querying database: %v", scanErr)
		}

		var listProps ItemListProperties
		if decodeErr := generic.DecodeJSONString(encodedProps, &listProps); decodeErr != nil {
			return nil, fmt.Errorf("getAllForms: can't decode properties: %v", encodedProps)
		}
		currList.Properties = listProps

		itemLists = append(itemLists, currList)
	}

	return itemLists, nil

}

func getDatabaseItemLists(databaseID string) ([]ItemList, error) {

	getTableParams := table.GetTableListParams{DatabaseID: databaseID}
	tables, err := table.GetTableList(getTableParams)
	if err != nil {
		return nil, fmt.Errorf("getDatabaseItemLists: %v", err)
	}

	allItemLists := []ItemList{}

	for _, currTable := range tables {
		tableItemLists, err := getAllItemLists(currTable.TableID)
		if err != nil {
			return nil, fmt.Errorf("getDatabaseItemLists: can't get lists: %v", err)
		}
		allItemLists = append(allItemLists, tableItemLists...)
	}

	return allItemLists, nil

}

func CloneTableItemLists(remappedIDs uniqueID.UniqueIDRemapper, srcParentTableID string) error {

	remappedTableID, err := remappedIDs.GetExistingRemappedID(srcParentTableID)
	if err != nil {
		return fmt.Errorf("CloneTableLists: Error getting remapped table ID: %v", err)
	}

	lists, err := getAllItemLists(srcParentTableID)
	if err != nil {
		return fmt.Errorf("CloneTableLists: Error getting forms for parent table ID = %v: %v",
			srcParentTableID, err)
	}

	for _, currList := range lists {

		destList := currList
		destList.ParentTableID = remappedTableID

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

func CloneItemLists(remappedIDs uniqueID.UniqueIDRemapper, srcDatabaseID string) error {

	getTableParams := table.GetTableListParams{DatabaseID: srcDatabaseID}
	tables, err := table.GetTableList(getTableParams)
	if err != nil {
		return fmt.Errorf("CloneItemLists: %v", err)
	}

	for _, srcTable := range tables {

		if err := CloneTableItemLists(remappedIDs, srcTable.TableID); err != nil {
			return fmt.Errorf("CloneItemLists: %v", err)
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

	databaseID := ""
	getErr := databaseWrapper.DBHandle().QueryRow(
		`SELECT database_id 
			FROM data_tables, item_lists 
			WHERE item_lists.list_id=$1 
				AND item_lists.table_id=data_tables.table_id LIMIT 1`,
		listID).Scan(&databaseID)
	if getErr != nil {
		return "", fmt.Errorf(
			"getItemListDatabaseID: can't get database for list = %v: err=%v",
			listID, getErr)
	}

	return databaseID, nil

}
