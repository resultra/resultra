function captionFromCaptionContainer($caption) {
	return 	$caption.find(".formCaption")
}


function formCaptionContainerHTML(elementID)
{	
	var containerHTML = ''+
		'<div class="ui-widget-content layoutContainer captionFormContainer" id="'+elementID+'">' +
			'<div class="well well-sm formCaptionContent">' +
				'<div class="formCaption"></div>' +
			'</div>' +
		'</div><';
						
	return containerHTML
}