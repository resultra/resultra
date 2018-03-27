function openNewFieldDialog(databaseID) {
	
	var $newFieldDialogForm = $('#newFieldDialogForm')
	var newFieldPanel = new NewFieldPanel(databaseID,$newFieldDialogForm)
	var $newFieldDialog = $('#newFieldDialog')	
	
	var fieldCreated = false
	initButtonClickHandler('#newFieldSaveButton',function() {
		console.log("New field save button clicked")
		if(newFieldPanel.validateNewFieldParams()) {
			
			// Prevent the creation of multiple fields at once.
			if (fieldCreated === false) {
				fieldCreated = true
				newFieldPanel.createNewField(function(newField) {
					if (newField !== null) {
						$newFieldDialog.modal('hide')
					
						var editPropsContentURL = '/admin/field/'+newField.fieldID
						setSettingsPageContent(editPropsContentURL,function() {
							initFieldPropsSettingsPageContent(newField.fieldID)
						})			
					}
				})			
				
			}
			

		}
	})
	
	$newFieldDialog.modal('show')
	
}