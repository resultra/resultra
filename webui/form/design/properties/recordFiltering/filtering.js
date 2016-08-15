function initDesignFormFilterProperties(tableID,formInfo) {
	
	function changeDefaultFilterSelection(filterID, isChecked) {
		var selectedFilterIDs = getFilterCheckboxListSelectedFilterIDs('#designFormDefaultFilterList')
		console.log("Default filters: updated selection: " + JSON.stringify(selectedFilterIDs))
		
		var setDefaultFiltersParams = {
			formID: formID,
			defaultFilterIDs: selectedFilterIDs
		}
		
		jsonAPIRequest("frm/setDefaultFilters",setDefaultFiltersParams,function(updatedForm) {
			console.log(" Default filters updated")
		}) // set record's number field value
		
		
	}

	function changeAvailableFilterSelection(filterID, isChecked) {
		
		var selectedFilterIDs = getFilterCheckboxListSelectedFilterIDs('#designFormAvailableFilterList')
		console.log("Available filters: updated selection: " + JSON.stringify(selectedFilterIDs))
		
		var setAvailFiltersParams = {
			formID: formID,
			availableFilterIDs: selectedFilterIDs
		}
		
		jsonAPIRequest("frm/setAvailableFilters",setAvailFiltersParams,function(updatedForm) {
			console.log("Available filters updated")
		}) // set record's number field value
	}
	
	jsonAPIRequest("filter/getList",{parentTableID:tableID},function(filterList) {
						
		initializeFilterCheckboxList('#designFormDefaultFilterList','defaultFilters_',
			filterList,formInfo.properties.defaultFilterIDs,changeDefaultFilterSelection)
		
		initializeFilterCheckboxList('#designFormAvailableFilterList','availFilters_',
			filterList,formInfo.properties.availableFilterIDs,changeAvailableFilterSelection)
			
	})
		

}