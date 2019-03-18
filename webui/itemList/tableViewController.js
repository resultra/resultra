// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.


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
			
			// The table needs to be refreshed/resized after initialization is complete.
			tableViewContext.resizeTable()
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