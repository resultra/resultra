function headerIDFromContainerElemID(headerElemID) {
	return 	headerElemID + '_header'
}


function formHeaderContainerHTML(elementID)
{	
	var headerID = headerIDFromContainerElemID(elementID)
	
	var containerHTML = ''+
		'<div class="ui-widget-content layoutContainer headerFormContainer draggable resizable" id="'+elementID+'">' +
			'<h3 class="formHeader" id="' + headerID + '">' +
			'New Header (placeholder)' +
			'</h3>' +
		'</div><';
						
	return containerHTML
}
