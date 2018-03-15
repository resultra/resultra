

$(document).ready(function() {	
	 
	initMainWindowLayout()
	
	function loadWorkspaceHomePageContent() {
		theMainWindowLayout.disableRHSSidebar()
		theMainWindowLayout.disableLHSSidebar()
		clearMainWindowHeaderButtonsContent()
		resetWorkspaceBreadcrumbHeader()
		
		setMainWindowContent('/homePage',function() {
			initHomePageSignedInPageContent(mainWindowContext)
		})
		
		setMainWindowOffPageContent('/homePage/offPageContent',function() {
		})
	}
	
	function loadSettingsPageContent() {
		theMainWindowLayout.disableRHSSidebar()		
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
	
	initMainWindowBreadcrumbHeader()
	initUserDropdownMenu(mainWindowContext.isSingleUserWorkspace)
	initHelpDropdownMenu()
						
}); // document ready