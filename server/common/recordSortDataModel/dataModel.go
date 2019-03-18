package recordSortDataModel

import (
	"fmt"
	"resultra/tracker/server/generic/uniqueID"
)

const SortDirectionAsc string = "asc"
const SortDirectionDesc string = "desc"

type RecordSortRule struct {
	SortFieldID   string `json:"fieldID"`
	SortDirection string `json:"direction"`
}

func ValidSortDirection(sortDir string) bool {
	if (sortDir == SortDirectionAsc) || (sortDir == SortDirectionDesc) {
		return true
	} else {
		return false
	}
}

func (srcRule RecordSortRule) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*RecordSortRule, error) {

	remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcRule.SortFieldID)
	if err != nil {
		return nil, fmt.Errorf("RecordSortRule.Clone: %v", err)
	}

	destRule := srcRule
	destRule.SortFieldID = remappedFieldID

	return &destRule, nil
}

func CloneSortRules(remappedIDs uniqueID.UniqueIDRemapper, srcRules []RecordSortRule) ([]RecordSortRule, error) {

	destRules := []RecordSortRule{}

	for _, srcRule := range srcRules {
		destRule, err := srcRule.Clone(remappedIDs)
		if err != nil {
			return nil, fmt.Errorf("CloneSortRules: %v", err)
		}
		destRules = append(destRules, *destRule)
	}

	return destRules, nil
}
