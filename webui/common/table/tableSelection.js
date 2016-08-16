


function populateTableSelectionMenu(menuSelector, databaseID) {
	
	function populateSelectionMenuTableList(tableRefList, menuSelector) {
		$(menuSelector).empty()		
		$(menuSelector).append(defaultSelectOptionPromptHTML("Select a Table"))
		$.each(tableRefList, function(index, tableRef) {
			$(menuSelector).append(selectFieldHTML(tableRef.tableID, tableRef.name))		
		})
	}
	
	var tableListParams =  { databaseID: databaseID }
	jsonAPIRequest("table/getList",tableListParams,function(tableRefs) {
		populateSelectionMenuTableList(tableRefs,menuSelector)
	})
	
}