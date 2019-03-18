// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function openAttachLinkDialog(params) {
	
	var $dialog = $('#attachLinkDialog')
	var $form = $('#attachLinkForm')
	var $titleInput = $dialog.find(".attachmentThumbnailTitle")
	var $urlInput = $dialog.find(".attachmentURLLink")
	var $captionInput = $dialog.find(".attachmentCaptionTextArea")
	
	$titleInput.val("")
	$urlInput.val("")
	$captionInput.val("")
	resetFormValidationFeedback($form)
	
	var remoteValidationParams = {
		url: '/api/generic/stringValidation/validateOptionalItemLabel',
		data: {
			label: function() { return $titleInput.val(); }
		}
	}
	
	var $attachLinkForm = $('#attachLinkForm')
	var validator = $attachLinkForm.validate({
		rules: {
			attachmentURLLink: {
				required: true,
				url:true
			},
			attachmentThumbnailTitle: {
				remote: remoteValidationParams
			}
		},
		messages: {
			attachmentURLLink: {
				required: "A valid URL is required"
			},
			attachmentThumbnailTitle: {
				required: "Invalid title. A title must consist of letters, numbers and spaces, but no special characters."
			}
		}
	})
	
	var $saveButton = $('#attachLinkSaveButton')
	initButtonControlClickHandler($saveButton,function() {
		console.log("Save attachment button clicked")
		if($attachLinkForm.valid()) {
			
			var saveURLParams = {
				parentDatabaseID: params.parentDatabaseID,
				url: $urlInput.val(),
				title: $titleInput.val(),
				caption: $captionInput.val()
			}
			jsonAPIRequest("attachment/saveURL", saveURLParams, function(attachInfo) {
				console.log("URL saved: " + JSON.stringify(attachInfo))
				params.addLinkCallback(attachInfo.attachmentID)
			})
				
			$dialog.modal("hide")
		}
	})
	
	$dialog.modal("show")
	
//	initAttachmentInfo($dialog,attachRef)
	
}