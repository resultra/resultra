

$(document).ready(function() {	
	 
	initMainWindowLayout()
	
	function loadWorkspaceHomePageContent() {
		theMainWindowLayout.hideRHSSidebar()
		theMainWindowLayout.hideLHSSidebar()
		
		setMainWindowContent('/homePage',function() {
			initHomePageSignedInPageContent(mainWindowContext)
		})
	}
	
	function loadHomePageSignedOut() {
		initHomePagePublicPageContent(mainWindowContext)
	}
	
	registerMainWindowContentLoader("workspaceHome",loadWorkspaceHomePageContent)
	registerMainWindowContentLoader("",loadWorkspaceHomePageContent)
	
	const linkID = getMainWindowLinkIDAnchorName()
	
	loadWorkspaceHomePageContent()
	
	
	
/*	var loadLastViewCallback = null */
	
	
	/*
	
	*/
	
	initUserDropdownMenu(mainWindowContext.isSingleUserWorkspace)
	initHelpDropdownMenu()
//	initAlertHeader(mainWindowContext.databaseID,seeAllAlertsClicked)
	
	// Listen for events to view a specific record/item in a particular form. This happens in response to
	// clicks to a form button deeper down in the DOM.
	
/*
	$('#formViewContainer,#tableViewContainer').on(viewFormInViewportEventName,function(e,params) {
		e.stopPropagation()
		console.log("Got event in main window: " + JSON.stringify(params))
		
		params.loadLastViewCallback = loadLastViewCallback
		
		mainWinLayout.clearSidebarNavigationSelection()
		loadExistingItemView(mainWinLayout,mainWindowContext.databaseID,params)
		
	})
*/	
	
//	hideSiblingsShowOne('#listViewProps')
					
}); // document ready