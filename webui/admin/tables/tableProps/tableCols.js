function initTableViewColsProperties(tableRef) {
	
	
	loadFieldInfo(tableRef.parentDatabaseID,[fieldTypeAll],function(fieldsByID) {
		
		function savedUpdatedColumnOrder(tableID) {
			console.log("Saving updated columns: " + tableID)
		}
		
		function populateOneTableColInTableColList(tableCol) {
			var $colListItem = $('#tableColItemTemplate').clone()
			$colListItem.attr("id","")
			
			var fieldName = fieldsByID[tableCol.properties.fieldID].name
			$colListItem.find('label').text(fieldName)
			
			var editColLink = '/admin/tablecol/' + tableCol.columnID
			$colListItem.find('.editTableColButton').attr("href",editColLink)
			
			var $deleteColButton = $colListItem.find('.deleteTableColButton')
			
			initButtonControlClickHandler($deleteColButton,function() {
				openFormComponentConfirmDeleteDialog("column",function() {
					console.log("column deletion confirmed")
					var deleteParams = {
						parentTableID: tableCol.parentTableID,
						columnID: tableCol.columnID
					}
					jsonAPIRequest("tableView/deleteColumn",deleteParams,function(replyStatus) {
						$colListItem.remove()
						console.log("Delete confirmed")
						savedUpdatedColumnOrder(tableCol.parentTableID)
					})
					
				})				
			})
			
		
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