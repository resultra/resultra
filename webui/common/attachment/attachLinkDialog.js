function openAttachLinkDialog(params) {
	
	var $dialog = $('#attachLinkDialog')
	
	var $attachLinkForm = $('#attachLinkForm')
	var validator = $attachLinkForm.validate({
		rules: {
			attachmentURLLink: {
				required: true,
				url:true
			},
		},
		messages: {
			attachmentURLLink: {
				required: "A valid URL is required"
			}
		}
	})
	
	var $saveButton = $('#attachLinkSaveButton')
	
	var $urlInput = $dialog.find(".attachmentURLLink")
	
	initButtonControlClickHandler($saveButton,function() {
		console.log("Save attachment button clicked")
		if($attachLinkForm.valid()) {
			
			var saveURLParams = {
				parentDatabaseID: params.parentDatabaseID,
				url: $urlInput.val()
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