function attachmentButtonHTML(className) {
	
	// className is to uniquely identify the button with other HTML elements,
	// such that it can be found with jQuery's find() function.
	
	var buttonHTML = '<button class="btn btn-default btn-sm clearButton ' + 
			className + 
			'"><span class="glyphicon glyphicon-paperclip"></span></button>'
	
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
