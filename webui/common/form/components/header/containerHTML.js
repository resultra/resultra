function headerFromHeaderContainer($header) {
	return 	$header.find(".formHeader")
}


function formHeaderContainerHTML(elementID)
{	
	var containerHTML = ''+
		'<div class="ui-widget-content layoutContainer headerFormContainer" id="'+elementID+'">' +
			'<h3 class="formHeader">' +
			'New Header' +
			'</h3>' +
		'</div><';
						
	return containerHTML
}
