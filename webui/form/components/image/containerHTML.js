

function imageIDFromContainerElemID(imageElemID) {
	return 	imageElemID + '_image'
}

function imageUploadInputIDFromContainerElemID(imageElemID) {
	return 	imageElemID + '_uploadFileInput'
}

function fileNameLabelFromContainerElemID(imageElemID) {
	return 	imageElemID + '_fileNameLabel'
}

function imageContainerLabelIDFromContainerElemID(imageElemID) {
	return 	imageElemID + '_containerLabel'
}

function imageLinkIDFromContainerElemID(imageElemID) {
	return 	imageElemID + '_imageLink'
}



function imageContainerHTML(elementID)
{	
	var imageID = imageIDFromContainerElemID(elementID)
	var uploadInputID = imageUploadInputIDFromContainerElemID(elementID)
	var fileNameLabelID = fileNameLabelFromContainerElemID(elementID)
	var containerLabelID = imageContainerLabelIDFromContainerElemID(elementID)
	
	// Adding title=" " to the file input field prevents the jQuery File Upload
	// plugin from displaying it's own messages.
	
	var containerHTML = ''+
	'<div class="ui-widget-content layoutContainer imageContainer  draggable resizable" id="'+elementID+'">' +
		'<label id="' + containerLabelID + '">Image Label</label>' +
		'<div>' +
			'<input id="'+ uploadInputID + '" type="file" title=" " single>'+
			'<label id="' + fileNameLabelID + '"></label>' +
		'</div>' +
		'<div id="' + imageID + '" class="imageInnerContainer">'+
		'</div>'+
	'</div>';
	
		
	return containerHTML
}


function imageLinkHTML(elementID, imageURL) {
	var linkID = imageLinkIDFromContainerElemID(elementID)
	return '<a href="' + imageURL + '" id="' + linkID + '">' + 
		'<img class="imageComponentImage" src="' + imageURL + '">'+
	'</a>'
}