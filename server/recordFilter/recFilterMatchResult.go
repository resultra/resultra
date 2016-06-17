package recordFilter

type RecFilterMatchResults map[string]bool

func (matchResults RecFilterMatchResults) MatchAllFilterIDs(filterIDs []string) bool {
	for _, currFilterID := range filterIDs {
		matchResult, resultFound := matchResults[currFilterID]
		if (resultFound == false) || (matchResult == false) {
			return false
		}
	}
	return true
}
