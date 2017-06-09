

function ItemListTableViewController($parentContainer,databaseID,resortCallback) {
	
	var dataTable
	var currRecordData
	var resizeTableFunc
	
	function updateDataTableData() {
		if (dataTable !== undefined && currRecordData !== undefined) {
			dataTable.clear()
			dataTable.rows.add(currRecordData)
			dataTable.draw()
		}
	}
	
	this.setTable = function(tableID) {
		console.log("ItemListTableViewController: setting table: " + tableID)
		
		function initTableDoneCallback(tableViewDataTable,resizeFunc) {
			dataTable = tableViewDataTable
			resizeTableFunc = resizeFunc
			updateDataTableData()
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
	
	this.setRecordData = function(recordData) {
		currRecordData = recordData
		updateDataTableData()
	}
	
	this.refresh = function() {
		if (resizeTableFunc !== undefined) {
			resizeTableFunc()
		}
		
	}
	
}