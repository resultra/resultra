$(document).ready(function() {	
	 
				
	
	function resizeMainWindow() {
		console.log("Resizing list view")
/*		if (tableViewController !== undefined) {
			tableViewController.refresh()
		} */
	}
	var mainWinLayout = new MainWindowLayout(resizeMainWindow)
	
	var loadLastViewCallback = null
	
	
	function itemListClicked(listID) {
		function loadView() {
			loadItemListView(mainWinLayout,mainWindowContext.databaseID,listID)
		}
		console.log("Main window: item list clicked: " + listID)
		loadView()
		loadLastViewCallback = loadView
	}
	
	function dashboardClicked(dashboardID) {
		console.log("Main window: dashboard navigation clicked: " + dashboardID)
		function loadView() {
			loadDashboardView(mainWinLayout,mainWindowContext.databaseID, dashboardID)	
		}
		loadView()
		loadLastViewCallback = loadView
	}
	
	function newItemClicked(linkID) {
		console.log("Main window: new item clicked: " + linkID)
		loadNewItemView(mainWinLayout,mainWindowContext.databaseID,linkID)
	}
	
	function seeAllAlertsClicked() {
		function loadView() {
			mainWinLayout.clearSidebarNavigationSelection()
			initAlertNotificationList(mainWinLayout,mainWindowContext.databaseID)			
		}
		loadView()
		loadLastViewCallback = loadView
	}
	
		
	var tocConfig = {
		databaseID: mainWindowContext.databaseID,
		newItemFormButtonFunc: openSubmitFormDialog,
		itemListClickedCallback: itemListClicked,
		dashboardClickedCallback: dashboardClicked,
		newItemLinkClickedCallback: newItemClicked
	}	
	initDatabaseTOC(tocConfig)
	
	initUserDropdownMenu()
	initAlertHeader(mainWindowContext.databaseID,seeAllAlertsClicked)
	
	// Listen for events to view a specific record/item in a particular form. This happens in response to
	// clicks to a form button deeper down in the DOM.
	$('#formViewContainer').on(viewFormInViewportEventName,function(e,params) {
		e.stopPropagation()
		console.log("Got event in main window: " + JSON.stringify(params))
		
		params.loadLastViewCallback = loadLastViewCallback
		
		mainWinLayout.clearSidebarNavigationSelection()
		loadExistingItemView(mainWinLayout,mainWindowContext.databaseID,params)
		
	})
	
	
//	hideSiblingsShowOne('#listViewProps')
					
}); // document ready