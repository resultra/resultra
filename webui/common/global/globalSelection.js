function populateGlobalSelectionMenu(menuSelector, databaseID) {
	
	function populateSelectionMenuTableList(globalsInfo, menuSelector) {
		$(menuSelector).empty()		
		$(menuSelector).append(defaultSelectOptionPromptHTML("Select a global value"))
		$.each(globalsInfo, function(index, globalInfo) {
			$(menuSelector).append(selectFieldHTML(globalInfo.globalID, globalInfo.name))		
		})
	}
	
	var listParams =  { parentDatabaseID: databaseID }
	jsonAPIRequest("global/getList",listParams,function(globalsInfo) {
		populateSelectionMenuTableList(globalsInfo,menuSelector)
	})
	
}