// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function headerFromDashboardHeaderContainer($header) {
	return 	$header.find(".dashboardHeader")
}


function dashboardHeaderContainerHTML(elementID)
{	
	var containerHTML = ''+
		'<div class="layoutContainer dashboardHeaderFormContainer" id="'+elementID+'">' +
			'<span class="h3 dashboardHeader">' +
			'New Header' +
			'</span>' +
		'</div><';
						
	return containerHTML
}

function setHeaderDashboardComponentHeaderSize($container,headerSize) {
	
	var sizeSizeClassMap = {
		"xlarge":"h2",
		"large":"h3",
		"medium":"h4",
		"small":"h5",
		"xsmall":"h6"
	}
	var sizeClass = sizeSizeClassMap[headerSize]
	
	$container.find("span").removeClass("h1 h2 h3 h4 h5 h6")
	$container.find("span").addClass(sizeClass)
}

function setHeaderDashboardComponentUnderlined($container,isUnderlined) {
	
	var $headerSpan = $container.find("span")
	
	$headerSpan.removeClass("underlinedDashboardHeader")
	if(isUnderlined) {
		$headerSpan.addClass("underlinedDashboardHeader")
	}
}

function setHeaderDashboardComponentLabel($container,headerRef) {
	var $label = $container.find("span")
	$label.text(headerRef.properties.title)
	
	setHeaderDashboardComponentHeaderSize($container,headerRef.properties.size)
	setHeaderDashboardComponentUnderlined($container,headerRef.properties.underlined)
}
