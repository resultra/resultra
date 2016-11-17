
function initRecordFilterViewPanel(filterPaneParams) {
		
	initDefaultFilterRules(filterPaneParams) 
		
	var refilterButtonSelector = createPrefixedSelector(filterPaneParams.elemPrefix,'FilterRecordsPanelRefilterButton')
	initButtonClickHandler(refilterButtonSelector,function (){
		filterPaneParams.refilterWithCurrentFilterRules()		
	})
	
	
	var resetFiltersButtonSelector = createPrefixedSelector(filterPaneParams.elemPrefix,'FilterRecordsResetFiltersButton')
	initButtonClickHandler(resetFiltersButtonSelector,function () {
		updateDefaultFilterRules(filterPaneParams, function () {
			filterPaneParams.updateFilterRules(filterPaneParams.defaultFilterRules)
		})		
	})
	
}