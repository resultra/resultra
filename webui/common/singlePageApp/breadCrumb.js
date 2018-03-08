function initMainWindowBreadcrumbHeader() {
	var $workspaceHomeLink = $("#workspaceHomeBreadcrumbLink")
	$workspaceHomeLink.click(function(e) {
		e.preventDefault()
		$workspaceHomeLink.blur()
		navigateToMainWindowContent("workspaceHome")	
	})	
	
}