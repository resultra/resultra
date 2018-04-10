function initHomePageSignedInPageContent(pageContext) {	
	
	initTrackerList(pageContext)
		
}

function navigateToWorkspaceHomePageContent() {
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

