package valueList

import (
	"fmt"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/generic/uniqueID"
)

type NewValueListParams struct {
	Name             string `json:"name"`
	ParentDatabaseID string `json:"parentDatabaseID"`
	ValueType        string `json:"valueType"`
}

type ValueList struct {
	ValueListID      string              `json:"valueListID"`
	Name             string              `json:"name"`
	ParentDatabaseID string              `json:"parentDatabaseID"`
	Properties       ValueListProperties `json:"properties"`
}

func saveNewValueList(newValueList ValueList) error {

	encodedProps, encodeErr := generic.EncodeJSONString(newValueList.Properties)
	if encodeErr != nil {
		return fmt.Errorf("saveNewValueList: failure encoding properties: error = %v", encodeErr)
	}

	if _, insertErr := databaseWrapper.DBHandle().Exec(`INSERT INTO value_lists 
				(value_list_id,database_id,name,properties) 
				VALUES ($1,$2,$3,$4)`,
		newValueList.ValueListID,
		newValueList.ParentDatabaseID,
		newValueList.Name,
		encodedProps); insertErr != nil {
		return fmt.Errorf("saveNewValueList: Can't create preset: error = %v", insertErr)
	}
	return nil

}

func newValueList(params NewValueListParams) (*ValueList, error) {

	newProps := newDefaultValueListProperties()

	newValueList := ValueList{
		ValueListID:      uniqueID.GenerateSnowflakeID(),
		Name:             params.Name,
		ParentDatabaseID: params.ParentDatabaseID,
		Properties:       newProps}

	if saveErr := saveNewValueList(newValueList); saveErr != nil {
		return nil, fmt.Errorf("newValueList: %v", saveErr)
	}

	return &newValueList, nil
}

type GetValueListParams struct {
	ValueListID string `json:"valueListID"`
}

func GetValueList(valueListID string) (*ValueList, error) {

	valueList := ValueList{}
	encodedProps := ""
	getErr := databaseWrapper.DBHandle().QueryRow(
		`SELECT value_list_id,name,database_id,properties
			FROM value_lists WHERE
			value_list_id=$1 LIMIT 1`, valueListID).Scan(
		&valueList.ValueListID,
		&valueList.Name,
		&valueList.ParentDatabaseID,
		&encodedProps)
	if getErr != nil {
		return nil, fmt.Errorf("GetValueList: Unabled to get form link: link ID = %v: datastore err=%v",
			valueListID, getErr)
	}

	props := newDefaultValueListProperties()
	if decodeErr := generic.DecodeJSONString(encodedProps, &props); decodeErr != nil {
		return nil, fmt.Errorf("GetValueList: can't decode properties: %v", encodedProps)
	}
	valueList.Properties = props

	return &valueList, nil

}

type GetValueListsParams struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
}

func getAllValueLists(parentDatabaseID string) ([]ValueList, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT value_lists.value_list_id,value_lists.name,
						value_lists.database_id,
						value_lists.properties
				FROM databases,value_lists WHERE 
				databases.database_id=$1 AND value_lists.database_id=databases.database_id`,
		parentDatabaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetAllPresets: Failure querying database: %v", queryErr)
	}

	valueLists := []ValueList{}
	for rows.Next() {
		var currValueList ValueList
		encodedProps := ""

		if scanErr := rows.Scan(&currValueList.ValueListID,
			&currValueList.Name,
			&currValueList.ParentDatabaseID,
			&encodedProps); scanErr != nil {
			return nil, fmt.Errorf("GetAllForms: Failure querying database: %v", scanErr)
		}

		props := newDefaultValueListProperties()
		if decodeErr := generic.DecodeJSONString(encodedProps, &props); decodeErr != nil {
			return nil, fmt.Errorf("getAllValueLists: can't decode properties: %v", encodedProps)
		}
		currValueList.Properties = props

		valueLists = append(valueLists, currValueList)
	}

	return valueLists, nil

}

func updateExistingValueList(updatedValueList *ValueList) (*ValueList, error) {

	encodedProps, encodeErr := generic.EncodeJSONString(updatedValueList.Properties)
	if encodeErr != nil {
		return nil, fmt.Errorf("updateExistingValueList: failure encoding properties: error = %v", encodeErr)
	}

	if _, updateErr := databaseWrapper.DBHandle().Exec(`UPDATE value_lists 
				SET properties=$1,name=$2
				WHERE value_list_id=$3`,
		encodedProps,
		updatedValueList.Name,
		updatedValueList.ValueListID); updateErr != nil {
		return nil, fmt.Errorf("updateExistingValueList: Can't update form value list %v: error = %v",
			updatedValueList.ValueListID, updateErr)
	}

	return updatedValueList, nil

}

func CloneValueLists(remappedIDs uniqueID.UniqueIDRemapper, srcParentDatabaseID string) error {

	valueLists, err := getAllValueLists(srcParentDatabaseID)
	if err != nil {
		return fmt.Errorf("CloneValueLists: Error getting form links for parent database ID = %v: %v",
			srcParentDatabaseID, err)
	}

	for _, currValueList := range valueLists {

		destValueList := currValueList

		destValueListID, err := remappedIDs.AllocNewRemappedID(currValueList.ValueListID)
		if err != nil {
			return fmt.Errorf("CloneValueLists: %v", err)
		}
		destValueList.ValueListID = destValueListID

		destDatabaseID, err := remappedIDs.GetExistingRemappedID(currValueList.ParentDatabaseID)
		if err != nil {
			return fmt.Errorf("CloneValueLists: %v", err)
		}
		destValueList.ParentDatabaseID = destDatabaseID

		destProps, err := currValueList.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("CloneFormLinks: %v", err)
		}
		destValueList.Properties = *destProps

		if err := saveNewValueList(destValueList); err != nil {
			return fmt.Errorf("CloneFormLinks: %v", err)
		}

	}

	return nil

}
