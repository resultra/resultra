function openNewValueListDialog(databaseID) {
	
	var $newValueListForm = $('#adminNewValueListDialogForm')	
	var $valueListNameInput = $('#newValueListNameInput')
	var $valueTypeSelection = $('#newValueListValueTypeSelection')
	var $newValueListDialog = $('#adminValueListDialog')
	
	var $includeInSidebarCheckbox = $('#newFormLinkIncludeInSidebar')
	
	var validator = $newValueListForm.validate({
		rules: {
			newValueListNameInput: {
				minlength: 3,
				required: true,
				remote: {
					url: '/api/generic/stringValidation/validateItemLabel',
					data: {
						label: function() { return $valueListNameInput.val(); }
					}
				} // remote
			}, // newFormNameInput
			newValueListValueTypeSelection: { required:true }
		},
		messages: {
			newValueListNameInput: {
				required: "Value list name is required"
			}
		}
	})	

	resetFormValidationFeedback($newValueListForm)	
	$valueListNameInput.val("")
	$valueTypeSelection.val("")
	validator.resetForm()
		
		
	$newValueListDialog.modal('show')
	
	initButtonClickHandler('#newValueListSaveButton',function() {
		console.log("New value list save button clicked")
		if($newValueListForm.valid()) {	
			
			var newValueListParams = { 
				parentDatabaseID: databaseID, 
				name: $valueListNameInput.val(),
				valueType: $valueTypeSelection.val() }
				
			$newValueListDialog.modal('hide')
			/*
			
								
			jsonAPIRequest("formLink/new",newFormLinkParams,function(newFormLinkInfo) {
				console.log("Created new form link: " + JSON.stringify(newFormLinkInfo))
				$newPresetDialog.modal('hide')
				navigateToURL('/admin/formLink/' + newFormLinkInfo.linkID)
			})
			
			*/
			

		}
	})
	
}