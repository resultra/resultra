function openNewFormDialog(pageContext) {
	
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
						databaseID: pageContext.databaseID,
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
	
	var newFormCreated = false
	initButtonClickHandler('#newFormSaveButton',function() {
		console.log("New form save button clicked")
		if($newFormDialogForm.valid()) {	
			console.log("table selection: " + $('#newFormTableSelection').val() )
			
			if(newFormCreated === false) {
				newFormCreated = true
				
				var newFormParams = { 
					parentDatabaseID: pageContext.databaseID, 
					name: $formNameInput.val() }
				jsonAPIRequest("frm/new",newFormParams,function(newFormInfo) {
					console.log("Created new form: " + JSON.stringify(newFormInfo))
					$newFormDialog.modal('hide')
					navigateToFormDesignerPageContent(pageContext,newFormInfo)
				
				})
				
			}
			

		}
	})
	
}