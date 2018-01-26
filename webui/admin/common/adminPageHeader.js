function initAdminPageHeader(isSingleUserWorkspace) {
	initUserDropdownMenu(isSingleUserWorkspace)
	initHelpDropdownMenu()
	
}

function appendPageSpecificBreadcrumbHeader(link,label) {
	var $breadcrumbContainer = $('#adminBreadcrumbHeaderContainer')
	
	var $sep = $('<i class="fa fa-chevron-right marginRight5 marginLeft5" aria-hidden="true"></i>')
	
	var $breadCrumb = $('<a class="h4" href="link goes here"><i class="fa" aria-hidden="true">Label goes here</i></a>')
	$breadCrumb.attr("href",link)
	$breadCrumb.find("i").html(label)
	
	$breadcrumbContainer.append($sep)
	$breadcrumbContainer.append($breadCrumb)
	
}