package componentLayout

import "resultra/tracker/server/generic/uniqueID"

type LayoutComponentCol struct {
	ComponentIDs []string `json:"componentIDs"`
}

type LayoutComponentRow struct {
	Columns []LayoutComponentCol `json:"columns"`
}

type ComponentLayout []LayoutComponentRow

func (srcLayout ComponentLayout) Clone(remappedIDs uniqueID.UniqueIDRemapper) ComponentLayout {

	destLayout := ComponentLayout{}
	for _, srcRow := range srcLayout {

		destCols := []LayoutComponentCol{}
		for _, srcCol := range srcRow.Columns {

			remappedComponentIDs := []string{}
			for _, srcComponentID := range srcCol.ComponentIDs {
				remappedComponentID := remappedIDs.AllocNewOrGetExistingRemappedID(srcComponentID)
				remappedComponentIDs = append(remappedComponentIDs, remappedComponentID)
			}
			destCol := LayoutComponentCol{ComponentIDs: remappedComponentIDs}

			destCols = append(destCols, destCol)
		}

		destRow := LayoutComponentRow{Columns: destCols}
		destLayout = append(destLayout, destRow)
	}

	return destLayout
}
