function openManageAttachmentsDialog(configParams) {
	var $dialog = $('#manageAttachmentsDialog')
	
	var currAttachments = configParams.attachmentList.slice(0)
	var $attachmentList = $('#manageAttachmentsAttachmentList')
	
	function populateOneAttachmentListItem(attachRef) {
		var $listItem = $('#manageAttachmentsAttachmentListItemTemplate').clone()
		$listItem.attr("id","")
		
		initAttachmentInfo($listItem,attachRef)
		
		var $deleteButton = $listItem.find('.deleteAttachmentButton')
		initButtonControlClickHandler($deleteButton,function() {
			console.log("Deleting attachment: " + JSON.stringify(attachRef))
			var attachmentsWithoutDeletedAttachment = []
			for(var currAttachIndex = 0; currAttachIndex < currAttachments.length; currAttachIndex++) {
				var attachID = currAttachments[currAttachIndex]
				if (attachID != attachRef.attachmentInfo.attachmentID) {
					attachmentsWithoutDeletedAttachment.push(attachID)
				}
			}
			currAttachments = attachmentsWithoutDeletedAttachment
			configParams.changeAttachmentsCallback(currAttachments)
			$listItem.remove()
		})
				
		$attachmentList.append($listItem)
	}
	
	function repopulateAttachmentList() {
		var getRefParams = { attachmentIDs: currAttachments }
		jsonAPIRequest("attachment/getReferences", getRefParams, function(attachRefs) {
			$attachmentList.empty()
			for(var attachIndex=0; attachIndex<attachRefs.length; attachIndex++) {
				var attachRef = attachRefs[attachIndex]
				populateOneAttachmentListItem(attachRef)
			}
		})
	}
	
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
		repopulateAttachmentList()
		
	}
	
	var $addFilesButton = $('#manageAttachmentsAddFilesButton')
	var addAttachmentParams = {
		parentDatabaseID: configParams.parentDatabaseID,
		$addAttachmentInput: $addFilesButton,
		attachDoneCallback: addNewAttachments }
	initAddAttachmentControl(addAttachmentParams)
		
	repopulateAttachmentList()
	$dialog.modal("show")
}