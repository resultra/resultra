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
						
			jsonAPIRequest("valueList/new",newValueListParams,function(newValueListInfo) {
				console.log("Created new value list: " + JSON.stringify(newValueListInfo))
				$newValueListDialog.modal('hide')
				
				var editPropsContentURL = '/admin/valueList/' + newValueListInfo.valueListID
				setSettingsPageContent(editPropsContentURL,function() {
					initValueListSettingsPageContent(newValueListInfo)
				})
			})
			

		}
	})
	
}