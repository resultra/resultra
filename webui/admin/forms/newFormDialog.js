function openNewFormDialog(databaseID) {
	
	var $newFormDialogForm = $('#newFormDialogForm')
	var $formNameInput = $('#newFormNameInput')
	
	var validator = $newFormDialogForm.validate({
		rules: {
			newFormNameInput: {
				minlength: 3,
				required: true,
				remote: {
					url: '/api/frm/validateNewFormName',
					data: {
						databaseID: databaseID,
						formName: function() { return $('#newFormNameInput').val(); }
					}
				} // remote
			}, // newFormNameInput
		},
		messages: {
			newFormNameInput: {
				required: "Form name is required"
			}
		}
	})

	resetFormValidationFeedback($newFormDialogForm)
	$formNameInput.val("")
	validator.resetForm()
	
	var $newFormDialog = $('#newFormDialog')
		
	$newFormDialog.modal('show')
	
	initButtonClickHandler('#newFormSaveButton',function() {
		console.log("New form save button clicked")
		if($newFormDialogForm.valid()) {	
			console.log("table selection: " + $('#newFormTableSelection').val() )
			
			var newFormParams = { 
				parentDatabaseID: databaseID, 
				name: $formNameInput.val() }
			jsonAPIRequest("frm/new",newFormParams,function(newFormInfo) {
				console.log("Created new form: " + JSON.stringify(newFormInfo))
				$newFormDialog.modal('hide')
				// TODO - Include database ID in link
				var formDesignUrl = '/admin/frm/' + newFormInfo.formID
				window.location.href = formDesignUrl
			})
			

		}
	})
	
}