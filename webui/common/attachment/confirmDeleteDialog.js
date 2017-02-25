function openAttachmentConfirmDeleteDialog(deleteAttachmentCallback) {
	
	var $dialog = $('#attachmentConfirmDeleteDialog')
	
	$dialog.modal("show")
	
	var $confirmDeleteButton = $dialog.find(".attachmentConfirmDeleteButton")
	
	initButtonControlClickHandler($confirmDeleteButton,function() {
		deleteAttachmentCallback()
		$dialog.modal("hide")
	})
	
	
}