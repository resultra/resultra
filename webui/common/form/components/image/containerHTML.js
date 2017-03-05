


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
	'<div class="layoutContainer imageContainer">' +
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

function attachmentGalleryThumbnailContainer(attachRef,deleteAttachmentCallback) {
	
	var attachURL = attachRef.url
	
	var thumbnailHTML =  '' +
			'<div class="attachGalleryThumbnailContainer"' + 
				'<a class="" href="' + attachURL + '">' + 
					'<img class="img-thumbnail imageContainerImage" src="' + attachURL + '">'+
				'</a>'+
				'<div class="galleryThumbnailHoverButtons">' + 
					'<button class="btn btn-default btn-sm clearButton attachmentInfoButton marginRight10"><span class="glyphicon glyphicon-pencil"></span></button>' +
					'<button type="button" class="close deleteAttachButton" aria-label="Close"><span aria-hidden="true">&times;</span></button>'+
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
	
	function deleteAttachmentInThumbnail() {
		$thumbnail.remove()
		deleteAttachmentCallback()
	}
	
	var $deleteButton = $thumbnail.find(".deleteAttachButton")
	initButtonControlClickHandler($deleteButton,function() {
		openAttachmentConfirmDeleteDialog(deleteAttachmentInThumbnail)
	})
	
	var $infoButton = $thumbnail.find(".attachmentInfoButton")
	initButtonControlClickHandler($infoButton,function() {
		openAttachmentInfoDialog(attachRef)
	})
	
	// Prevent click-through from the buttons onto the thumbnail itself. This prevent the attachment from 
	// being displayed when a button is pressed in the button area.
	var $buttons = $thumbnail.find(".galleryThumbnailHoverButtons")
	$buttons.click(function (event){
		event.stopPropagation();
   	 	//   ... your code here
		return false;
	});
	
	
	return $thumbnail
	
}