function openSingleAttachmentDialog(configParams) {

	var $dialog = $('#singleAttachmentDialog')
	var currAttachmentRef = null
	
	function populateAttachmentInfo(attachmentID) {
		var getRefParams = { attachmentID: attachmentID }
		jsonAPIRequest("attachment/getReference", getRefParams, function(attachRef) {
			currAttachmentRef = attachRef
			var $attachItem = $('#singleAttachmentAttachmentItem')
			initAttachmentInfo($attachItem,attachRef)
			
			// Only show the dialog once the attachment information is ready for display.
			// This also handles the scenario where no initial attachment is given and 
			// the user is prompted for an initial upload.
			$dialog.modal("show")
		})		
	}
	
	function addNewAttachments(newAttachments) {
		console.log("New attachments added: " + JSON.stringify(newAttachments))
		
		if(newAttachments.length > 0) {
			var firstAttach = newAttachments[0]
			populateAttachmentInfo(firstAttach.attachmentID)
		}	
	}
	
	var $doneButton = $("#singleAttachmentDoneButton")
	initButtonControlClickHandler($doneButton,function() {
		
		if (currAttachmentRef != null) {
			configParams.setAttachmentCallback(currAttachmentRef.attachmentInfo.attachmentID)	
		}
		$dialog.modal("hide")
	})
	
	var $replaceButton = $('#singleAttachmentReplaceFileButton')
	var addAttachmentParams = {
		parentDatabaseID: configParams.parentDatabaseID,
		$addAttachmentInput: $replaceButton,
		attachDoneCallback: addNewAttachments,
		 acceptedFileTypes: configParams.acceptedFileTypes }
	initAddAttachmentControl(addAttachmentParams)
		
	if(configParams.attachmentID !== null) {
		populateAttachmentInfo(configParams.attachmentID)		
	} else {
		// If no initial attachment is given, immediately trigger a file upload.
		$replaceButton.trigger("click")
	}

}

function initAddAttachmentThenOpenInfoDialogButton(configParams) {
	
	function addNewAttachments(newAttachments) {
		console.log("New attachments added: " + JSON.stringify(newAttachments))
		
		if(newAttachments.length > 0) {
			var firstAttach = newAttachments[0]			
			var attachInfoParams = {
				parentDatabaseID: configParams.parentDatabaseID,
				setAttachmentCallback: configParams.setAttachmentCallback,
				attachmentID: firstAttach.attachmentID
			}
			openSingleAttachmentDialog(attachInfoParams)
		}	
	}
	
	
	var addAttachmentParams = {
		parentDatabaseID: configParams.parentDatabaseID,
		$addAttachmentInput: configParams.$addAttachmentInput,
		attachDoneCallback: addNewAttachments,
		acceptedFileTypes: configParams.acceptedFileTypes }
	initAddAttachmentControl(addAttachmentParams)
	
}