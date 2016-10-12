package recordSortDataModel

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
