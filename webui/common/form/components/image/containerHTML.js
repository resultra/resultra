


function imageContainerLabelImageComponentContainer($image) {
	return 	$image.find(".imageContainerLabel")
}

function imageUploadInputFromImageComponentContainer($image) {
	return 	$image.find(".imageComponentUploadInput")
}

function imageInnerContainerFromImageComponentContainer($image) {
	return $image.find(".imageInnerContainer")
}

function manageAttachmentsButtonFromImageComponentContainer($image) {
	return $image.find(".imageComponentManageAttachmentsButtton")
}

function imageContainerHTML(elementID)
{		
	// Adding title=" " to the file input field prevents the jQuery File Upload
	// plugin from displaying it's own messages.
	
	var containerHTML = ''+
	'<div class="ui-widget-content layoutContainer imageContainer draggable resizable">' +
		'<div class="imageContainerHeader">' +
			'<label class="imageContainerLabel">Image Label</label>' +
			attachmentButtonHTML("imageComponentManageAttachmentsButtton") + 
		'</div>' +
		'<div class="imageInnerContainer text-center"">'+
		'</div>'+
	'</div>';
	
		
	return containerHTML
}

function imageLinkIDFromContainerElemID(imageElemID) {
	return 	imageElemID + '_imageLink'
}

function imageGalleryThumbnailContainer(imageURL) {
	var thumbnailHTML =  '<a href="' + imageURL + '">' + 
		'<img class="img-thumbnail imageContainerImage" src="' + imageURL + '">'+
	'</a>'
	
	return $(thumbnailHTML)
	
}