function initTableViewColsProperties(tableRef) {
	
	var $columnList = $('#tableColPropsColList')
	
	function saveUpdatedColumnOrder() {
		console.log("Saving updated columns: " + tableRef.tableID)
		
		var columns = []
		
		$columnList.find('.list-group-item').each(function() {
			var columnID = $(this).attr('data-column-id')
			columns.push(columnID)
		})
		
		console.log("savedUpdatedColumnOrder: " + JSON.stringify(columns))
		
		var setColsParams = {
			tableID: tableRef.tableID,
			orderedColumns: columns
		}
		jsonAPIRequest("tableView/setOrderedCols",setColsParams,function(replyStatus) {
			
		})
	}
	
    $columnList.sortable({
		placeholder: "ui-state-highlight",
		cursor:"move",
		update: function( event, ui ) {
			saveUpdatedColumnOrder(tableRef.tableID)
		}
    });
	
	
	loadFieldInfo(tableRef.parentDatabaseID,[fieldTypeAll],function(fieldsByID) {
		
		function populateOneTableColInTableColList(tableCol) {
			var $colListItem = $('#tableColItemTemplate').clone()
			$colListItem.attr("id","")
			
			if (tableCol.colType !== 'button') {
				var fieldName = fieldsByID[tableCol.properties.fieldID].name
				$colListItem.find('label').text(fieldName)
			} else {
				$colListItem.find('label').text("Button: open form")
			}
			
			var editColLink = '/admin/tablecol/' + tableCol.columnID
			$colListItem.find('.editTableColButton').attr("href",editColLink)
			
			$colListItem.attr('data-column-id',tableCol.columnID)
			
			var $deleteColButton = $colListItem.find('.deleteTableColButton')
			
			initButtonControlClickHandler($deleteColButton,function() {
				openConfirmDeleteDialog("column",function() {
					console.log("column deletion confirmed")
					var deleteParams = {
						parentTableID: tableCol.parentTableID,
						columnID: tableCol.columnID
					}
					jsonAPIRequest("tableView/deleteColumn",deleteParams,function(replyStatus) {
						$colListItem.remove()
						console.log("Delete confirmed")
						saveUpdatedColumnOrder()
					})
					
				})				
			})
			
		
			$columnList.append($colListItem)
		}
		
		var params = { tableID: tableRef.tableID }	
		jsonAPIRequest("tableView/getTableDisplayInfo",params,function(tableInfo) {
			console.log("Loading table column properties: " + JSON.stringify(tableInfo))
		
			$columnList.empty()
			$.each(tableInfo.cols,function(colIndex,tableCol) {
				populateOneTableColInTableColList(tableCol)
			})
		
		})
		
	})
	
	initButtonClickHandler('#adminNewTableColButton',function() {
		console.log("New table column button clicked")
		openNewTableColDialog(tableRef)
	})
	
	initButtonClickHandler('#adminNewFormButtonColButton',function() {
		console.log("New table column button clicked")
		openNewButtonTableColDialog(tableRef)
	})
	
	
}