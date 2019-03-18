// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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