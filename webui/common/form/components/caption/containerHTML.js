function captionFromCaptionContainer($caption) {
	return 	$caption.find(".formCaption")
}


function formCaptionContainerHTML(elementID)
{	
	var containerHTML = ''+
		'<div class="ui-widget-content layoutContainer captionFormContainer" id="'+elementID+'">' +
			'<h3 class="formCaption">' +
			'New Caption' +
			'</h3>' +
		'</div><';
						
	return containerHTML
}
