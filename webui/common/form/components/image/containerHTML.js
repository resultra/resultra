


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
	
	var thumbnailHTML =  '' +
			'<div class="attachGalleryThumbnailContainer"' + 
				'<a class="" href="' + imageURL + '">' + 
					'<img class="img-thumbnail imageContainerImage" src="' + imageURL + '">'+
				'</a>'+
				'<div class="galleryThumbnailHoverButtons">' + 
					'<button class="btn btn-xs btn-default">Edit</button>' + 
				'</div>'+
			'</div>'
	
	var $thumbnail = $(thumbnailHTML)
	
	$thumbnail.hover(
		function() { 
			$(this).find(".galleryThumbnailHoverButtons").show()
		}, 
		function() { 
			$(this).find(".galleryThumbnailHoverButtons").hide()
		}
	)
	
	return $thumbnail
	
}