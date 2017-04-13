
function initRecordFilterViewPanel(filterPaneParams) {
		
	initDefaultFilterRules(filterPaneParams) 	
	
	var resetFiltersButtonSelector = createPrefixedSelector(filterPaneParams.elemPrefix,'FilterRecordsResetFiltersButton')
	initButtonClickHandler(resetFiltersButtonSelector,function () {
		updateDefaultFilterRules(filterPaneParams, function () {
			filterPaneParams.updateFilterRules(filterPaneParams.defaultFilterRules)
			filterPaneParams.refilterWithCurrentFilterRules()		
		})		
	})
	
}