

function ItemListTableViewController($parentContainer,databaseID) {
	
	var dataTable
	var currRecordData
	
	function updateDataTableData() {
		if (dataTable !== undefined && currRecordData !== undefined) {
			dataTable.clear()
			dataTable.rows.add(currRecordData)
			dataTable.draw()
		}
	}
	
	this.setTable = function(tableID) {
		console.log("ItemListTableViewController: setting table: " + tableID)
		initItemListTableView($parentContainer,databaseID,tableID, function (tableViewDataTable) {
			dataTable = tableViewDataTable
			updateDataTableData()
		})
	}
	
	this.setRecordData = function(recordData) {
		currRecordData = recordData
		updateDataTableData()
	}
	
}