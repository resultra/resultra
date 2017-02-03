function openNewNewItemPresetDialog(databaseID) {
	
	var $newPresetDialogForm = $('#adminNewNewItemPresetDialogForm')	
	var $presetNameInput = $('#newNewItemPresetNameInput')
	var $newPresetDialog = $('#adminNewNewItemPresetDialog')
	var formSelectionSelector = '#newNewItemPresetFormSelection'
	var $formSelection = $(formSelectionSelector)
	
	var $includeInSidebarCheckbox = $('#newNewItemPresetIncludeInSidebar')
	
	var validator = $newPresetDialogForm.validate({
		rules: {
			newNewItemPresetNameInput: {
				minlength: 3,
				required: true,
				remote: {
					url: '/api/generic/stringValidation/validateItemLabel',
					data: {
						label: function() { return $presetNameInput.val(); }
					}
				} // remote
			}, // newFormNameInput
			newNewItemPresetFormSelection: { required:true }
		},
		messages: {
			newNewItemPresetNameInput: {
				required: "Preset name is required"
			}
		}
	})	

	resetFormValidationFeedback($newPresetDialogForm)	
	$presetNameInput.val("")
	$formSelection.val("")
	validator.resetForm()
	$includeInSidebarCheckbox.prop("checked",false)
	
	var selectFormParams = {
		menuSelector: formSelectionSelector,
		parentDatabaseID: databaseID
	}	
	populateFormSelectionMenu(selectFormParams)
	
		
	$newPresetDialog.modal('show')
	
	initButtonClickHandler('#newNewItemPresetSaveButton',function() {
		console.log("New preset save button clicked")
		if($newPresetDialogForm.valid()) {	
			
			var newPresetParams = { 
				parentDatabaseID: databaseID, 
				name: $presetNameInput.val(),
				formID: $formSelection.val(),
				includeInSidebar:$includeInSidebarCheckbox.prop("checked") }
								
			jsonAPIRequest("formLink/new",newPresetParams,function(newPresetInfo) {
				console.log("Created new item preset: " + JSON.stringify(newPresetInfo))
				$newPresetDialog.modal('hide')
			})
			

		}
	})
	
}