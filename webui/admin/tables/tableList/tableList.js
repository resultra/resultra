

function initAdminTableListSettings(databaseID) {
	
	initButtonClickHandler('#adminNewTableButton',function() {
		console.log("New table button clicked")
		openNewTableDialog(databaseID)
	})
	
	
}