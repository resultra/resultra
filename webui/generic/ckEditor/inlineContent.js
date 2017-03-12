

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