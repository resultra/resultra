function openConfirmDeleteDialog(componentLabel,deleteCallback) {
	
	var $dialog = $('#confirmDeleteDialog')
	
	var $whatLabel = $dialog.find(".confirmDeleteWhatLabel")
	$whatLabel.text(componentLabel)
	
	$dialog.modal("show")
	
	var $confirmDeleteButton = $dialog.find(".confirmDeleteButton")
	
	initButtonControlClickHandler($confirmDeleteButton,function() {
		deleteCallback()
		$dialog.modal("hide")
	})
	
	
}