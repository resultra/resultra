// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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
