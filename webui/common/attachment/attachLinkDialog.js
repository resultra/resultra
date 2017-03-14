function openAttachLinkDialog() {
	
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
	initButtonControlClickHandler($saveButton,function() {
		console.log("Save attachment button clicked")
		if($attachLinkForm.valid()) {	
			$dialog.modal("hide")
		}
	})
	
	$dialog.modal("show")
	
//	initAttachmentInfo($dialog,attachRef)
	
}