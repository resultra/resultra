$(document).ready(function() {	
	 
				
	initUserDropdownMenu()
	initAlertHeader(mainWindowContext.databaseID)
		
	var tocConfig = {
		databaseID: mainWindowContext.databaseID,
		newItemFormButtonFunc: openSubmitFormDialog
	}	
	initDatabaseTOC(tocConfig)
	
	function resizeMainWindow() {
		console.log("Resizing list view")
/*		if (tableViewController !== undefined) {
			tableViewController.refresh()
		} */
	}
	var mainWinLayout = new MainWindowLayout(resizeMainWindow)
	
	
//	hideSiblingsShowOne('#listViewProps')
					
}); // document ready