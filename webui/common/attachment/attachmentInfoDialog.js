function openAttachmentInfoDialog(attachRef) {
	
	var $dialog = $('#attachmentInfoDialog')
	
	$dialog.modal("show")
	
	initAttachmentInfo($dialog,attachRef)
	
}