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
	
	function dashboardClicked(dashboardID) {
		console.log("Main window: dashboard navigation clicked: " + dashboardID)
		loadDashboardView(mainWinLayout,mainWindowContext.databaseID, dashboardID)
	}
	
	function newItemClicked(linkID) {
		console.log("Main window: new item clicked: " + linkID)
		loadNewItemView(mainWinLayout,mainWindowContext.databaseID,linkID)
	}
	
		
	var tocConfig = {
		databaseID: mainWindowContext.databaseID,
		newItemFormButtonFunc: openSubmitFormDialog,
		itemListClickedCallback: itemListClicked,
		dashboardClickedCallback: dashboardClicked,
		newItemLinkClickedCallback: newItemClicked
	}	
	initDatabaseTOC(tocConfig)
	
	
	
//	hideSiblingsShowOne('#listViewProps')
					
}); // document ready