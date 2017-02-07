


function imageContainerLabelImageComponentContainer($image) {
	return 	$image.find(".imageContainerLabel")
}

function imageUploadInputFromImageComponentContainer($image) {
	return 	$image.find(".imageComponentUploadInput")
}

function fileNameLabelFromImageComponentContainer($image) {
	return  $image.find(".imageComponentFileNameLabel")
}

function imageInnerContainerFromImageComponentContainer($image) {
	return $image.find(".imageInnerContainer")
}


function imageContainerHTML(elementID)
{		
	// Adding title=" " to the file input field prevents the jQuery File Upload
	// plugin from displaying it's own messages.
	
	var containerHTML = ''+
	'<div class="ui-widget-content layoutContainer imageContainer draggable resizable">' +
		'<div class="imageContainerHeader">' +
			'<label class="imageContainerLabel">Image Label</label>' +
			'<input class="imageComponentUploadInput" type="file" title=" " single>'+
			'<label class="imageComponentFileNameLabel"></label>' +
		'</div>' +
		'<div class="imageInnerContainer text-center"">'+
		'</div>'+
	'</div>';
	
		
	return containerHTML
}

function imageLinkIDFromContainerElemID(imageElemID) {
	return 	imageElemID + '_imageLink'
}

function imageLinkHTML(elementID, imageURL) {
	var linkID = imageLinkIDFromContainerElemID(elementID)
	return '<a href="' + imageURL + '" id="' + linkID + '">' + 
		'<img class="img-thumbnail imageContainerImage" src="' + imageURL + '">'+
	'</a>'
}