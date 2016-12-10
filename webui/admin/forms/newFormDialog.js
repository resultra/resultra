function openNewFormDialog(databaseID) {
	
	var $newFormDialogForm = $('#newFormDialogForm')
	
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

	validator.resetForm()
		
	$('#newFormDialog').modal('show')
	
	initButtonClickHandler('#newFormSaveButton',function() {
		console.log("New form save button clicked")
		if($newFormDialogForm.valid()) {	
			console.log("table selection: " + $('#newFormTableSelection').val() )
			
			var newFormParams = { 
				parentDatabaseID: databaseID, 
				name: $('#newFormNameInput').val() }
			jsonAPIRequest("frm/new",newFormParams,function(newFormInfo) {
				console.log("Created new form: " + JSON.stringify(newFormInfo))
				$('#newFormDialog').modal('hide')
			})
			

		}
	})
	
}