


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

function addLinkButtonFromAttachmentComponentContainer($image) {
	return $image.find(".attachmentComponentAddLinkButton")
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
			attachmentLinkButtonHTML("attachmentComponentAddLinkButton") +
		'</div>' +
		'<div class="imageInnerContainer lightGreyBorder text-center"">'+
		'</div>'+
	'</div>';
	
		
	return containerHTML
}

function initAttachmentFormComponentViewModeGeometry($container,attachRef) {
	// In view mode, the height will be flexible, up the maximum set in the form designer.
	// This ensures there isn't any "dead space" when there aren't enough attachments to
	// fill up the attachment area below the header.
	setElemFixedWidthFlexibleHeight($container,attachRef.properties.geometry.sizeWidth)
	
	var $header = $container.find(".imageContainerHeader")
	
	// Set the maximum height of the attachment area to be the remainder after the header
	// is accounted for.
	var headerBottom = $header.position().top + $header.outerHeight(true);
	var attachMaxHeightPx = (attachRef.properties.geometry.sizeHeight - (headerBottom+5)) + "px"
	
	var $innerAttachmentContainer = imageInnerContainerFromImageComponentContainer($container)
	
	$innerAttachmentContainer.css('max-height',attachMaxHeightPx)
	
}

function attachmentGalleryThumbnailContainer(attachRef,deleteAttachmentCallback) {
	
	var attachURL = attachRef.url
	
	var thumbnailHTML =  '' +
			'<div class="attachGalleryThumbnailContainer thumbnail">' + 
				'<a class="attachLink"></a>' + 
				'<div class="galleryThumbnailHoverButtons">' + 
					'<button class="btn btn-default btn-sm clearButton attachmentInfoButton marginRight10"><span class="glyphicon glyphicon-pencil"></span></button>' +
					'<button type="button" class="close deleteAttachButton" aria-label="Close"><span aria-hidden="true">&times;</span></button>'+
				'</div>'+
			'</div>'
	
	var $thumbnail = $(thumbnailHTML)
	
	// Initialize the link, depending upon if the data type is image or regular file.
	var $attachLink = $thumbnail.find(".attachLink")
	$attachLink.data("attachRef",attachRef)
	if(attachRef.dataType === "image") {
		$attachLink.attr("href",attachRef.url)
		$attachLink.append('<img class="imageContainerImage" src="' + attachURL + '">')
		$attachLink.addClass("mfp-image")
	} else if (attachRef.dataType === "link") {
		var $linkThumbnail = $('<div><i class="smallAttachmentThumbnailIcon glyphicon glyphicon-link"></i>' +
				'<span class="smallAttachmentThumbnailText"></span></div>')
		var $thumbnailText = $linkThumbnail.find(".smallAttachmentThumbnailText")
		$thumbnailText.text("link")
		$attachLink.append($linkThumbnail)
		$attachLink.addClass("mfp-inline")
	} else {
		var $fileThumbnail = $('<div><i class="smallAttachmentThumbnailIcon glyphicon glyphicon-file"></i>' +
				'<span class="smallAttachmentThumbnailText"></span></div>')
		var $thumbnailText = $fileThumbnail.find(".smallAttachmentThumbnailText")
		$thumbnailText.text(attachRef.extension)
		$attachLink.append($fileThumbnail)
		$attachLink.addClass("mfp-inline")
	}	
	
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