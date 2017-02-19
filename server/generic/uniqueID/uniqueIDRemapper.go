package uniqueID

import "fmt"

type UniqueIDRemapper map[string]string

func (mapper UniqueIDRemapper) AllocNewRemappedID(unmappedID string) (string, error) {
	remappedID, foundExisting := mapper[unmappedID]
	if foundExisting {
		return "", fmt.Errorf("AllocNewRemappedID: remapped ID already exists for ID=%v", unmappedID)
	}
	remappedID = GenerateSnowflakeID()
	mapper[unmappedID] = remappedID
	return remappedID, nil
}

func (mapper UniqueIDRemapper) AllocNewOrGetExistingRemappedID(unmappedID string) string {
	remappedID, foundExisting := mapper[unmappedID]
	if !foundExisting {
		remappedID = GenerateSnowflakeID()
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
