function openManageAttachmentsDialog(configParams) {
	var $dialog = $('#manageAttachmentsDialog')
	
	var currAttachments = configParams.attachmentList.slice(0)
	var $attachmentList = $('#manageAttachmentsAttachmentList')
	
	function populateOneAttachmentListItem(attachRef) {
		var $listItem = $('#manageAttachmentsAttachmentListItemTemplate').clone()
		$listItem.attr("id","")
		
		var $thumbnail = $listItem.find(".attachmentThumbnailImage")
		$thumbnail.attr("src",attachRef.url)
		$thumbnail.attr("alt",attachRef.attachmentInfo.origFileName)
		
		var $itemTitle = $listItem.find('.attachmentThumbnailTitle')
		$itemTitle.val(attachRef.attachmentInfo.origFileName)
		
		var $thumbnailContainer = $listItem.find(".attachmentThumbnailContainer")
		
		$attachmentList.append($listItem)
	}
	
	var getRefParams = { attachmentIDs: currAttachments }
	jsonAPIRequest("attachment/getReferences", getRefParams, function(attachRefs) {
		$attachmentList.empty()
		for(var attachIndex=0; attachIndex<attachRefs.length; attachIndex++) {
			var attachRef = attachRefs[attachIndex]
			populateOneAttachmentListItem(attachRef)
		}
	})
	
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
	
	var $addFilesButton = $('#manageAttachmentsAddFilesButton')
	var addAttachmentParams = {
		parentDatabaseID: configParams.parentDatabaseID,
		$addAttachmentInput: $addFilesButton,
		attachDoneCallback: addNewAttachments }
	initAddAttachmentControl(addAttachmentParams)
	
	$dialog.modal("show")
}