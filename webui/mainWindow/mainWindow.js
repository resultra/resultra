

$(document).ready(function() {	
	 
	initMainWindowLayout()
	
	
	function loadSettingsPageContent() {
		theMainWindowLayout.disableRHSSidebar()		
		setMainWindowContent('/homePage',function() {
			initHomePageSignedInPageContent(mainWindowContext)
		})
	}
	
	function loadHomePageSignedOut() {
		initHomePagePublicPageContent(mainWindowContext)
	}
	
	registerMainWindowContentLoader("workspaceHome",navigateToWorkspaceHomePageContent)
	registerMainWindowContentLoader("",navigateToWorkspaceHomePageContent)
	registerMainWindowContentLoader("workspaceTemplates",navigateToTemplatesPage)
	
	
	
	const linkID = getMainWindowLinkIDAnchorName()
	
	navigateToWorkspaceHomePageContent()
	
	initMainWindowBreadcrumbHeader()
	initUserDropdownMenu(mainWindowContext.isSingleUserWorkspace)
	initHelpDropdownMenu()
						
}); // document ready