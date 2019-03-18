// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.


function initAttachmentThumbnailContainer($parentContainer, attachRef) {
	
	var $thumbnailContainer =  $parentContainer.find(".attachmentThumbnailContainer")
	
	var $thumbnailImage = $thumbnailContainer.find(".attachmentThumbnailImage")
	var $thumbnailIcon = $thumbnailContainer.find(".attachmentThumbnailIcon")
	var $thumbnailText = $thumbnailContainer.find(".attachmentThumbnailText")
	$thumbnailContainer.attr("href",attachRef.url)
	if(attachRef.dataType === "image") {
		$thumbnailImage.attr("src",attachRef.url)
		$thumbnailImage.attr("alt",attachRef.attachmentInfo.origFileName)
		$thumbnailIcon.hide()
		$thumbnailText.hide()
	} else {
		$thumbnailImage.hide()
		$thumbnailText.text(attachRef.extension)
	}
}

function initAttachmentTitleInput($parentContainer,attachRef) {
	
	var $titleInput = $parentContainer.find('.attachmentThumbnailTitle')
	
	$titleInput.val(attachRef.attachmentInfo.title)
	$titleInput.unbind("blur")
	$titleInput.bind("blur",function() {
		var setTitleParams = {
			attachmentID: attachRef.attachmentInfo.attachmentID,
			title: $titleInput.val()
		}
		jsonAPIRequest("attachment/setTitle", setTitleParams, function(updateAttachment) {
		})
	})
	
}

function initAttachmentCaptionInput($parentContainer,attachRef) {
	
	var $captionTextArea = $parentContainer.find(".attachmentCaptionTextArea")
	
	$captionTextArea.val(attachRef.attachmentInfo.caption)
	
	$captionTextArea.unbind("blur")
	$captionTextArea.bind("blur",function() {
		var setCaptionParams = {
			attachmentID: attachRef.attachmentInfo.attachmentID,
			caption: $captionTextArea.val()
		}
		jsonAPIRequest("attachment/setCaption", setCaptionParams, function(updateAttachment) {
		})
	})
	
}

function initAttachmentInfo($parentContainer,attachRef) {
	initAttachmentThumbnailContainer($parentContainer,attachRef)
	initAttachmentTitleInput($parentContainer,attachRef)
	initAttachmentCaptionInput($parentContainer,attachRef)	
}