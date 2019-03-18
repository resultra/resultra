// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.


// Post-process basic/raw HTML that is created by the CK editor for read-only display.
// Notably, this includes replacing links with a _blank target, so clicking on the links will
// open a new window or tab.
function formatInlineContentHTMLDisplay(rawHTML) {
	
	var $display = $("<div></div>")
	$display.html(rawHTML)
	
	$display.find('a').attr("target","_blank")
	
	// script tags not allowed
	$display.find('script').remove()
	
	// TODO - Remove any other unapproved content like scripts
	
	var displayHTML = $display.html()
	
	return displayHTML
}

function populateInlineDisplayContainerHTML($container,rawHTML) {
	var displayHTML = formatInlineContentHTMLDisplay(rawHTML)
	$container.html(displayHTML)
}