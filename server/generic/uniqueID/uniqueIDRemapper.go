// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package uniqueID

import "fmt"

type UniqueIDRemapper map[string]string

func (mapper UniqueIDRemapper) AllocNewRemappedID(unmappedID string) (string, error) {
	remappedID, foundExisting := mapper[unmappedID]
	if foundExisting {
		return "", fmt.Errorf("AllocNewRemappedID: remapped ID already exists for ID=%v", unmappedID)
	}
	remappedID = GenerateUniqueID()
	mapper[unmappedID] = remappedID
	return remappedID, nil
}

func (mapper UniqueIDRemapper) AllocNewOrGetExistingRemappedID(unmappedID string) string {
	remappedID, foundExisting := mapper[unmappedID]
	if !foundExisting {
		remappedID = GenerateUniqueID()
		mapper[unmappedID] = remappedID
	}
	return remappedID
}

func (mapper UniqueIDRemapper) GetExistingRemappedID(unmappedID string) (string, error) {
	remappedID, foundExisting := mapper[unmappedID]
	if !foundExisting {
		return "", fmt.Errorf("getExistingRemappedID: unabled to retrieve remapped ID for ID=%v", unmappedID)
	}
	return remappedID, nil
}

func CloneIDList(remapper UniqueIDRemapper, srcIDList []string) []string {

	destIDList := []string{}
	for _, srcID := range srcIDList {
		destID := remapper.AllocNewOrGetExistingRemappedID(srcID)
		destIDList = append(destIDList, destID)
	}
	return destIDList
}
