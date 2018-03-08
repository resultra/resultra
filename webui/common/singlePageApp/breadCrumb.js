function initMainWindowBreadcrumbHeader() {
	var $workspaceHomeLink = $("#workspaceHomeBreadcrumbLink")
	$workspaceHomeLink.click(function(e) {
		e.preventDefault()
		$workspaceHomeLink.blur()
		navigateToMainWindowContent("workspaceHome")	
	})	
	
}

function resetWorkspaceBreadcrumbHeader() {
	var $breadcrumbContainer = $('#mainWindowBreadcrumbHeader')
	$breadcrumbContainer.children().not('#workspaceHomeBreadcrumbLink').remove()
}


function appendMainWindowContentSpecificBreadcrumbHeader(label, linkClickedCallback) {
	var $breadcrumbContainer = $('#mainWindowBreadcrumbHeader')
	
	var $sep = $('<i class="fa fa-chevron-right marginRight5 marginLeft5" aria-hidden="true"></i>')
	
	var $breadCrumb = $('<a class="h4" href="link goes here"><i class="fa" aria-hidden="true">Label goes here</i></a>')
	$breadCrumb.find("i").text(label)
	
	$breadCrumb.click(function(e) {
		e.preventDefault()
		$breadCrumb.blur()
		linkClickedCallback()
	})
	
	$breadcrumbContainer.append($sep)
	$breadcrumbContainer.append($breadCrumb)
	
}