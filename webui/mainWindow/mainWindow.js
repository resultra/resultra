$(document).ready(function() {	
	 
	var mainWinLayout = new MainWindowLayout()
	
	var loadLastViewCallback = null
	
	
	function itemListClicked(listID,$tocItem) {
		function loadView() {
			loadItemListView(mainWinLayout,mainWindowContext.databaseID,listID)
			mainWinLayout.clearSidebarNavigationSelection()
			$tocItem.addClass("active")
		}
		console.log("Main window: item list clicked: " + listID)
		loadView()
		loadLastViewCallback = loadView
	}
	
	function dashboardClicked(dashboardID,$tocItem) {
		console.log("Main window: dashboard navigation clicked: " + dashboardID)
		function loadView() {
			loadDashboardView(mainWinLayout,mainWindowContext.databaseID, dashboardID)	
			mainWinLayout.clearSidebarNavigationSelection()
			$tocItem.addClass("active")
		}
		loadView()
		loadLastViewCallback = loadView
	}
	
	function newItemClicked(linkID,$tocItem) {
		console.log("Main window: new item clicked: " + linkID)
		var newItemParams = {
			pageLayout: mainWinLayout,
			databaseID: mainWindowContext.databaseID,
			formLinkID: linkID,
			loadLastViewCallback: loadLastViewCallback
		}
		
		loadNewItemView(newItemParams)
		mainWinLayout.clearSidebarNavigationSelection()
		$tocItem.addClass("active")
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
	initHelpDropdownMenu()
	initAlertHeader(mainWindowContext.databaseID,seeAllAlertsClicked)
	
	// Listen for events to view a specific record/item in a particular form. This happens in response to
	// clicks to a form button deeper down in the DOM.
	$('#formViewContainer,#tableViewContainer').on(viewFormInViewportEventName,function(e,params) {
		e.stopPropagation()
		console.log("Got event in main window: " + JSON.stringify(params))
		
		params.loadLastViewCallback = loadLastViewCallback
		
		mainWinLayout.clearSidebarNavigationSelection()
		loadExistingItemView(mainWinLayout,mainWindowContext.databaseID,params)
		
	})
	
	
//	hideSiblingsShowOne('#listViewProps')
					
}); // document ready