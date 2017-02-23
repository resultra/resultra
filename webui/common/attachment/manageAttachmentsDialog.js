function openManageAttachmentsDialog(configParams) {
	var $dialog = $('#manageAttachmentsDialog')
	
	var currAttachments = configParams.attachmentList.slice(0)
	
	var $addFilesButton = $('#manageAttachmentsAddFilesButton')
	
	function addNewAttachments(newAttachments) {
		console.log("New attachments added: " + JSON.stringify(newAttachments))
		
		// update currAttachments (list of attachment IDs) to include the 
		// attachments which were was just added.
		var attachmentList = currAttachments.slice(0)
		for(var attachIndex = 0; attachIndex < newAttachments.length; attachIndex++) {
			var newAttachment = newAttachments[attachIndex]
			attachmentList.push(newAttachment.attachmentID)
		}
		currAttachments = attachmentList
		
		configParams.changeAttachmentsCallback(currAttachments)
		
	}
	
	var addAttachmentParams = {
		parentDatabaseID: configParams.parentDatabaseID,
		$addAttachmentInput: $addFilesButton,
		attachDoneCallback: addNewAttachments }
	initAddAttachmentControl(addAttachmentParams)
	
	$dialog.modal("show")
}