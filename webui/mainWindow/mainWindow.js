$(document).ready(function() {	
	 
				
	initUserDropdownMenu()
	initAlertHeader(mainWindowContext.databaseID)
	
	function resizeMainWindow() {
		console.log("Resizing list view")
/*		if (tableViewController !== undefined) {
			tableViewController.refresh()
		} */
	}
	var mainWinLayout = new MainWindowLayout(resizeMainWindow)
	
	
	function itemListClicked(listID) {
		console.log("Main window: item list clicked: " + listID)
		loadItemListView(mainWinLayout,mainWindowContext.databaseID,listID)
	}
	
		
	var tocConfig = {
		databaseID: mainWindowContext.databaseID,
		newItemFormButtonFunc: openSubmitFormDialog,
		itemListClickedCallback: itemListClicked
	}	
	initDatabaseTOC(tocConfig)
	
	
	
//	hideSiblingsShowOne('#listViewProps')
					
}); // document ready