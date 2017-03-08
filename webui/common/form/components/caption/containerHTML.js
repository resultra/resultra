function captionFromCaptionContainer($caption) {
	return 	$caption.find(".formCaption")
}


function formCaptionContainerHTML(elementID)
{	
	var containerHTML = ''+
		'<div class=" layoutContainer captionFormContainer" id="'+elementID+'">' +
			'<div class="well well-sm formCaptionContent">' +
				'<div class="formCaption bg-info"></div>' +
			'</div>' +
		'</div><';
						
	return containerHTML
}

function setFormCaptionColorScheme($captionContainer,colorScheme) {
	
	var captionClassLookup = {
		default:"",
		info:"bg-info",
		primary: "bg-primary",
		success: "bg-success",
		warning: "bg-warning",
		danger:"bg-danger"
	}
	
	var captionClass = captionClassLookup[colorScheme]
	if (captionClass === undefined) {
		captionClass = ""
	}
	
	var $caption = $captionContainer.find(".formCaption")
	$caption.removeClass("bg-info bg-primary bg-success bg-warning bg-danger")
	$caption.addClass(captionClass)
	
}