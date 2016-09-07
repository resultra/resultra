function populateGlobalSelectionMenu(menuSelector, databaseID, globalTypes) {
	
	function populateSelectionMenuTableList(globalsInfo, menuSelector) {
		$(menuSelector).empty()		
		$(menuSelector).append(defaultSelectOptionPromptHTML("Select a global value"))
		$.each(globalsInfo, function(index, globalInfo) {
			$(menuSelector).append(selectFieldHTML(globalInfo.globalID, globalInfo.name))		
		})
	}
	
	var includeType = {}
	for (var globalTypeIndex = 0; globalTypeIndex < globalTypes.length; globalTypeIndex++) {
		var type = globalTypes[globalTypeIndex]
		includeType[type] = true
	}
	
	function filterByType(globalInfo) {
		if((includeType['all'] == true) || (includeType[globalInfo.type] == true)) {
			return true
		} else {
			return false
		}
	}
	
	
	
	var listParams =  { parentDatabaseID: databaseID }
	jsonAPIRequest("global/getList",listParams,function(globalsInfo) {
		
		var filteredInfo = globalsInfo.filter(filterByType)
		
		populateSelectionMenuTableList(filteredInfo,menuSelector)
	})
	
}