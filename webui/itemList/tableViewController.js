

function ItemListTableViewController($parentContainer,databaseID,resortCallback) {
	
	var currRecordData
	
	var tableViewContext
	
	function updateDataTableData(sortRules) {
		if (tableViewContext !== undefined && currRecordData !== undefined) {
			tableViewContext.updateData(currRecordData,sortRules)
		}
	}
	
	this.setTable = function(tableID,sortRules) {
		console.log("ItemListTableViewController: setting table: " + tableID)
		
		function initTableDoneCallback(tableContext) {
			tableViewContext = tableContext
			updateDataTableData(sortRules)
		}
		
		var tableViewParams = {
			$tableContainer: $parentContainer,
			databaseID: databaseID,
			tableID: tableID,
			initDoneCallback: initTableDoneCallback,
			resortCallback:resortCallback
		}
		initItemListTableView(tableViewParams)
	}
	
	this.setRecordData = function(recordData,sortRules) {
		currRecordData = recordData
		updateDataTableData(sortRules)
	}
	
	this.refresh = function() {
		if (tableViewContext !== undefined) {
			tableViewContext.resizeTable()
		}
		
	}
	
}