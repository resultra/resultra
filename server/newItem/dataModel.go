package newItem

import (
	"fmt"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/databaseWrapper"
	//	"resultra/datasheet/server/generic/stringValidation"
	"resultra/datasheet/server/generic/uniqueID"
)

type NewNewItemPresetParams struct {
	Name             string `json:"name"`
	FormID           string `json:"formID"`
	IncludeInSidebar bool   `json:"includeInSidebar"`
}

type NewItemPreset struct {
	PresetID         string                  `json:"presetID"`
	Name             string                  `json:"name"`
	FormID           string                  `json:"formID"`
	IncludeInSidebar bool                    `json:"includeInSidebar"`
	Properties       NewItemPresetProperties `json:"properties"`
}

func saveNewPreset(newPreset NewItemPreset) error {

	encodedProps, encodeErr := generic.EncodeJSONString(newPreset.Properties)
	if encodeErr != nil {
		return fmt.Errorf("savePreset: failure encoding properties: error = %v", encodeErr)
	}

	if _, insertErr := databaseWrapper.DBHandle().Exec(`INSERT INTO new_item_presets 
				(preset_id,form_id,name,include_in_sidebar,properties) VALUES ($1,$2,$3,$4,$5)`,
		newPreset.PresetID,
		newPreset.FormID,
		newPreset.Name,
		newPreset.IncludeInSidebar,
		encodedProps); insertErr != nil {
		return fmt.Errorf("savePreset: Can't create preset: error = %v", insertErr)
	}
	return nil

}

func newNewItemPreset(params NewNewItemPresetParams) (*NewItemPreset, error) {

	newProps := newDefaultNewItemProperties()

	newPreset := NewItemPreset{
		PresetID:         uniqueID.GenerateSnowflakeID(),
		Name:             params.Name,
		FormID:           params.FormID,
		IncludeInSidebar: params.IncludeInSidebar,
		Properties:       newProps}

	if saveErr := saveNewPreset(newPreset); saveErr != nil {
		return nil, fmt.Errorf("newNewItemPreset: %v", saveErr)
	}

	return &newPreset, nil
}

type GetPresetListParams struct {
	ParentDatabaseID string `json:"parentDatabaseID"`
}

func getAllPresets(parentDatabaseID string) ([]NewItemPreset, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT new_item_presets.preset_id,new_item_presets.name,new_item_presets.form_id,
						new_item_presets.include_in_sidebar,new_item_presets.properties
				FROM forms,new_item_presets WHERE 
				forms.database_id=$1 AND new_item_presets.form_id=forms.form_id`,
		parentDatabaseID)
	if queryErr != nil {
		return nil, fmt.Errorf("GetAllPresets: Failure querying database: %v", queryErr)
	}

	presets := []NewItemPreset{}
	for rows.Next() {
		var currPreset NewItemPreset
		encodedProps := ""

		if scanErr := rows.Scan(&currPreset.PresetID,
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

func ClonePresets(remappedIDs uniqueID.UniqueIDRemapper, srcParentDatabaseID string) error {

	presets, err := getAllPresets(srcParentDatabaseID)
	if err != nil {
		return fmt.Errorf("ClonePresets: Error getting presets for parent database ID = %v: %v",
			srcParentDatabaseID, err)
	}

	for _, currPreset := range presets {

		destPreset := currPreset

		destPresetID, err := remappedIDs.AllocNewRemappedID(currPreset.PresetID)
		if err != nil {
			return fmt.Errorf("CloneTableForms: %v", err)
		}
		destPreset.PresetID = destPresetID

		destFormID, err := remappedIDs.GetExistingRemappedID(currPreset.FormID)
		if err != nil {
			return fmt.Errorf("CloneTableForms: %v", err)
		}
		destPreset.FormID = destFormID

		destProps, err := currPreset.Properties.Clone(remappedIDs)
		if err != nil {
			return fmt.Errorf("ClonePresets: %v", err)
		}
		destPreset.Properties = *destProps

		if err := saveNewPreset(destPreset); err != nil {
			return fmt.Errorf("ClonePresets: %v", err)
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
