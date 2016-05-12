

function imageIDFromContainerElemID(imageElemID) {
	return 	imageElemID + '_image'
}


function imageContainerHTML(elementID)
{	
	var imageID = imageIDFromContainerElemID(elementID)
	
	var containerHTML = ''+
	'<div class="ui-widget-content layoutContainer imageContainer  draggable resizable" id="'+elementID+'">' +
		'<label>Image Label</label>' +
		'<div id="' + imageID + '" class="imageInnerContainer">'+
		' drop zone inside div' +
		' drop zone inside div' +
		' drop zone inside div' +
		' drop zone inside div' +
		' drop zone inside div' +
		' drop zone inside div' +
		' drop zone inside div' +
		'</div>'+
	'</div>';
	
		
	return containerHTML
}
