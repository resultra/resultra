// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.


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