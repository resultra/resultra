// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function openNewAlertDialog(databaseID) {
	
	var $newAlertDialogForm = $('#newAlertDialogForm')
	var $alertNameInput = $('#newAlertNameInput')
	
	var validator = $newAlertDialogForm.validate({
		rules: {
			newAlertNameInput: {
				minlength: 3,
				required: true,
				remote: {
					url: '/api/alert/validateNewAlertName',
					data: {
						databaseID: databaseID,
						alertName: function() { return $alertNameInput.val(); }
					}
				} // remote
			}, // newAlertNameInput
			newAlertFormSelection: {
				required:true,
			}
		},
		messages: {
			newAlertNameInput: {
				required: "Alert name is required"
			},
			newAlertFormSelection: {
				required: "Form selection is required"
			}
		}
	})
	
	var selectFormParams = {
		menuSelector: "#newAlertFormSelection",
		parentDatabaseID: databaseID,
		initialFormID: null
	}	
	populateFormSelectionMenu(selectFormParams)
	var $formSelection = $("#newAlertFormSelection")
		
	resetFormValidationFeedback($newAlertDialogForm)
	$alertNameInput.val("")
	validator.resetForm()
	
	var $newAlertDialog = $('#newAlertDialog')
		
	$newAlertDialog.modal('show')

	var alertCreated = false
	initButtonClickHandler('#newAlertSaveButton',function() {
		console.log("New alert save button clicked")
		if($newAlertDialogForm.valid()) {	
			
			if (alertCreated === false) {
				
				// Only support the creation of a single alert from the dialog box.
				// This prevents the creation of multiple alerts with the same properties,
				// in the event the user "double tapped" the OK button.
				alertCreated = true
				
				var newAlertParams = { 
					parentDatabaseID: databaseID, 
					name: $alertNameInput.val(),
					formID: $formSelection.val()}
				jsonAPIRequest("alert/new",newAlertParams,function(newAlertInfo) {
					console.log("Created new alert: " + JSON.stringify(newAlertInfo))
					$newAlertDialog.modal('hide')
				
					var editPropsContentURL = '/admin/alert/' + newAlertInfo.alertID
					setSettingsPageContent(editPropsContentURL,function() {
						initAlertSettingsAdminPageContent(databaseID,newAlertInfo)
					})
				
				})
				
				
			}
			
		}
	})
	
}