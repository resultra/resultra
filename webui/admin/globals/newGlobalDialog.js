function openNewGlobalDialog(databaseID) {
	
	var $newGlobalDialogForm = $('#adminNewGlobalForm')
	var $newGlobalDialog = $('#adminNewGlobalDialog')
	var $nameInput = $('#adminGlobalNewGlobalNameInput')
	var $typeSelection = $("#adminGlobalNewGlobalTypeSelection")
	
	var validator = $newGlobalDialogForm.validate({
		rules: {
			adminGlobalNewGlobalNameInput: {
				minlength: 3,
				required: true,
				remote: {
					url: '/api/global/validateNewName',
					data: {
						databaseID: databaseID,
						globalName: function() { return $nameInput.val(); }
					}
				} // remote
			}, // new
			adminGlobalNewGlobalTypeSelection: { optionSelectionRequired:"type" }
		},
		messages: {
			adminGlobalNewGlobalNameInput: {
				required: "Global name is required"
			}
		}
	})

	$nameInput.val("")
	$typeSelection.val("")
	validator.resetForm()
	
	
	$newGlobalDialog.modal('show')

	initButtonClickHandler('#newGlobalDialogSaveGlobalButton',function() {
		console.log("New global save button clicked")
		if($newGlobalDialogForm.valid()) {				
			var newGlobalParams = { 
				parentDatabaseID:databaseID, 
				name: $nameInput.val(),
				type: $typeSelection.val()}
			jsonAPIRequest("global/new",newGlobalParams,function(newGlobalInfo) {
				console.log("Created new global: " + JSON.stringify(newGlobalInfo))
				addGlobalToAdminList(newGlobalInfo)
				$newGlobalDialog.modal('hide')
			})
		}
	})
	
}