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
	
	if (btnSizeClass !== undefined && btnSizeClass.length > 0) {
			$button.addClass(btnSizeClass)
	}

}

function setFormButtonColorScheme($container,newScheme) {
	var $button = $container.find('button')
	$button.removeClass("btn-primary btn-default btn-success btn-info btn-warning btn-danger btn-link")
	
	var schemeBtnClassMap = {
		"default": "btn-default",
		"primary":"btn-primary",
		"success":"btn-success",
		"info":"btn-info",
		"warning":"btn-warning",
		"danger":"btn-danger",
		"link":"btn-link"
	}
		
	var schemeClass = schemeBtnClassMap[newScheme]
	
	if (schemeClass !== undefined && schemeClass.length > 0) {
			$button.addClass(schemeClass)
	}
	
}