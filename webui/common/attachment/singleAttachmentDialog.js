function openSingleAttachmentDialog(configParams) {

	var $dialog = $('#singleAttachmentDialog')
	var currAttachmentRef = null
	
	function addNewAttachments(newAttachments) {
		console.log("New attachments added: " + JSON.stringify(newAttachments))
		
		if(newAttachments.length > 0) {
			var firstAttach = newAttachments[0]
			var getRefParams = { attachmentID: firstAttach.attachmentID }
			jsonAPIRequest("attachment/getReference", getRefParams, function(attachRef) {
				currAttachmentRef = attachRef
				var $attachItem = $('#singleAttachmentAttachmentItem')
				initAttachmentInfo($attachItem,attachRef)
			})
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
		attachDoneCallback: addNewAttachments }
	initAddAttachmentControl(addAttachmentParams)
				
	
	$dialog.modal("show")
}