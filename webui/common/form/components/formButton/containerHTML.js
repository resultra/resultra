function buttonFromFormButtonContainer($formButton) {
	return 	$formButton.find(".formButton")
}


function formButtonContainerHTML(elementID)
{	
	var containerHTML = ''+
		'<div class="ui-widget-content layoutContainer buttonFormContainer">' +
			'<button type="button" class="btn btn-primary formButton">' + 
			'Open Form' +
			'</button>' +
		'</div><';
						
	return containerHTML
}
