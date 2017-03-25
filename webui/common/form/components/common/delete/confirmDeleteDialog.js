function openFormComponentConfirmDeleteDialog(componentLabel,deleteCallback) {
	
	var $dialog = $('#formComponentConfirmDeleteDialog')
	
	var $whatLabel = $dialog.find(".confirmDeleteWhatLabel")
	$whatLabel.text(componentLabel)
	
	$dialog.modal("show")
	
	var $confirmDeleteButton = $dialog.find(".formComponentConfirmDeleteButton")
	
	initButtonControlClickHandler($confirmDeleteButton,function() {
		deleteCallback()
		$dialog.modal("hide")
	})
	
	
}