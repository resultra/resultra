

function ItemListTableViewController($parentContainer,databaseID) {
	
	this.setTable = function(tableID) {
		console.log("ItemListTableViewController: setting table: " + tableID)
		initItemListTableView($parentContainer,databaseID,tableID)
	}
	
}