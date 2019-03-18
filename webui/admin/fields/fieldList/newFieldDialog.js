// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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