function buttonFromFormButtonContainer($formButton) {
	return 	$formButton.find(".formButton")
}


function formButtonContainerHTML(elementID)
{	
	var containerHTML = ''+
		'<div class=" layoutContainer buttonFormContainer">' +
			'<button type="button" class="btn btn-primary formButton">' + 
			'Open Form' +
			'</button>' +
		'</div><';
						
	return containerHTML
}


function setFormButtonSize($container,newSize) {
	
	var $button = $container.find('button')
	$button.removeClass("btn-lg btn-sm btn-xs")
	
	var sizeBtnClassMap = {
		"large": "btn-lg",
		"medium":"",
		"small":"btn-sm",
		"xsmall":"btn-xs"
	}
	
	var btnSizeClass = sizeBtnClassMap[newSize]
	
	if (btnSizeClass !== null && btnSizeClass.length > 0) {
			$button.addClass(btnSizeClass)
	}

}