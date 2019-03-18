// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function headerFromHeaderContainer($header) {
	return 	$header.find(".formHeader")
}


function formHeaderContainerHTML(elementID)
{	
	var containerHTML = ''+
		'<div class="layoutContainer headerFormContainer" id="'+elementID+'">' +
			'<span class="h4 formHeader">' +
			'New Header' +
			'</span>' +
		'</div><';
						
	return containerHTML
}

function setHeaderFormComponentHeaderSize($container,headerSize) {
	
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

function setHeaderFormComponentUnderlined($container,isUnderlined) {
	
	var $header = $container.find("span")
	
	$header.removeClass("underlinedFormHeader")
	if(isUnderlined) {
		$header.addClass("underlinedFormHeader")
	}
}