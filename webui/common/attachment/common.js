// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function attachmentButtonHTML(className) {
	
	// className is to uniquely identify the button with other HTML elements,
	// such that it can be found with jQuery's find() function.
	
	var buttonHTML = '<button tabindex="-1" class="btn btn-default btn-sm clearButton ' + 
			className + 
			'"><span class="glyphicon glyphicon-paperclip"></span></button>'
	
	return buttonHTML
}

function attachmentLinkButtonHTML(className) {
	
	// className is to uniquely identify the button with other HTML elements,
	// such that it can be found with jQuery's find() function.
	
	var buttonHTML = '<button tabindex="-1" class="btn btn-default btn-sm clearButton ' + 
			className + 
			'"><span class="glyphicon glyphicon-link"></span></button>'
	
	return buttonHTML
}

function attachmentCaptionHTML(attachRef) {
	return '<small class="attachCaptionText">' 
		+  escapeHTML(attachRef.attachmentInfo.caption) + '</small>';
}

function attachmentTitleAndCaptionHTML(attachRef) {
	var label = '<label>' + escapeHTML(attachRef.attachmentInfo.title) + "</label>"

	return label + '<small class="attachCaptionText">' 
		+  escapeHTML(attachRef.attachmentInfo.caption) + '</small>';
	
}
