



function initFilterPropertyPanel(panelParams) {
	
	var defaultFilterListSelector = createPrefixedSelector(panelParams.elemPrefix,'DefaultFilterList')			
	var availListSelector = createPrefixedSelector(panelParams.elemPrefix,'AvailableFilterList')	
	
	function changeOneAvailableFilterSelection(filterID, isChecked) {
	
		var selectedFilterIDs = getFilterCheckboxListSelectedFilterIDs(availListSelector)
		console.log("Available filters: updated selection: " + JSON.stringify(selectedFilterIDs))
		
		panelParams.setAvailableFilterFunc(selectedFilterIDs)
	
	}
	
	function changeOneDefaultFilterSelection(filterID, isChecked) {
	
		var selectedFilterIDs = getFilterCheckboxListSelectedFilterIDs(defaultFilterListSelector)
		console.log("Available filters: updated selection: " + JSON.stringify(selectedFilterIDs))
		
		panelParams.setDefaultFilterFunc(selectedFilterIDs)
	
	}
	
	jsonAPIRequest("filter/getList",{parentTableID:panelParams.tableID},function(filterList) {
							
		initializeFilterCheckboxList(defaultFilterListSelector,'defaultFilters_',
			filterList,panelParams.defaultFilterIDs,changeOneDefaultFilterSelection)
	
		initializeFilterCheckboxList(availListSelector,'availFilters_',
			filterList,panelParams.availableFilterIDs,changeOneAvailableFilterSelection)
		
	})
	
	
	var manageFiltersSelector = createPrefixedSelector(panelParams.elemPrefix,'FilterRecordsManageFiltersButton')	
	initButtonClickHandler(manageFiltersSelector,function() {
		openRecordFilterManageFiltersDialog(panelParams.tableID)	
	}) 
	
}
