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
	
	initButtonClickHandler('#newAlertSaveButton',function() {
		console.log("New alert save button clicked")
		if($newAlertDialogForm.valid()) {	
			var newAlertParams = { 
				parentDatabaseID: databaseID, 
				name: $alertNameInput.val(),
				formID: $formSelection.val()}
			jsonAPIRequest("alert/new",newAlertParams,function(newAlertInfo) {
				console.log("Created new alert: " + JSON.stringify(newAlertInfo))
				$newAlertDialog.modal('hide')
				// TODO - Include database ID in link
				navigateToURL('/admin/alert/' + newAlertInfo.alertID)
			})
			

		}
	})
	
}