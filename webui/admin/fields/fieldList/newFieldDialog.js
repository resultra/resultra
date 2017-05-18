function openNewFieldDialog(databaseID) {
	
	var $newFieldDialogForm = $('#newFieldDialogForm')
	var newFieldPanel = new NewFieldPanel(databaseID,$newFieldDialogForm)
	var $newFieldDialog = $('#newFieldDialog')	
	
	initButtonClickHandler('#newFieldSaveButton',function() {
		console.log("New field save button clicked")
		if(newFieldPanel.validateNewFieldParams()) {
			newFieldPanel.createNewField(function(newField) {
				if (newField !== null) {
					$newFieldDialog.modal('hide')
					navigateToURL('/admin/field/'+newField.fieldID)				
				}
			})			

		}
	})
	
	$newFieldDialog.modal('show')
	
}