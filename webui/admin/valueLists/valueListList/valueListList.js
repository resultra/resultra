function initAdminValueListListSettings(databaseID) {
	
	initButtonClickHandler('#adminNewValueListButton',function() {
		console.log("New value list button clicked")
		openNewValueListDialog(databaseID)
	})
	
	
}