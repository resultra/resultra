function initAdminGlobals(databaseID) {
	initButtonClickHandler('#adminGlobalsNewGlobalButton',function() {
		console.log("New Global button clicked")
		openNewGlobalDialog(databaseID)
	})
	
}