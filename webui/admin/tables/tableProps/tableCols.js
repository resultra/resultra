function initTableViewColsProperties(tableRef) {
	
	
	
	loadFieldInfo(tableRef.parentDatabaseID,[fieldTypeAll],function(fieldsByID) {
		
		function populateOneTableColInTableColList(tableCol) {
			var $colListItem = $('#tableColItemTemplate').clone()
			$colListItem.attr("id","")
			var fieldName = fieldsByID[tableCol.properties.fieldID].name
			$colListItem.find('label').text(fieldName)
		
			$('#tableColPropsColList').append($colListItem)
		}
		
		var params = { tableID: tableRef.tableID }	
		jsonAPIRequest("tableView/getColumns",params,function(tableCols) {
			console.log("Loading table column properties: " + JSON.stringify(tableCols))
		
			$.each(tableCols,function(colIndex,tableCol) {
				populateOneTableColInTableColList(tableCol)
			})
		})
		
	})
	
	initButtonClickHandler('#adminNewTableColButton',function() {
		console.log("New table column button clicked")
		openNewTableColDialog(tableRef)
	})
	
}