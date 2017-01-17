function buttonIDFromContainerElemID(buttonElemID) {
	return 	buttonElemID + '_button'
}


function formButtonContainerHTML(elementID)
{	
	var buttonID = buttonIDFromContainerElemID(elementID)
	
	var containerHTML = ''+
		'<div class="ui-widget-content layoutContainer buttonFormContainer" id="'+elementID+'">' +
			'<button type="button" class="btn btn-primary formButton" + id="' + buttonID + '">' + 
			'Open Form' +
			'</button>' +
		'</div><';
						
	return containerHTML
}
