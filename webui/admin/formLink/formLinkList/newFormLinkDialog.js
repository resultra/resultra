function openNewNewItemPresetDialog(pageContext) {
	
	var $newPresetDialogForm = $('#adminFormLinkDialogForm')	
	var $presetNameInput = $('#newFormLinkNameInput')
	var $newPresetDialog = $('#adminNewFormLinkDialog')
	var formSelectionSelector = '#newFormLinkFormSelection'
	var $formSelection = $(formSelectionSelector)
	
	var $includeInSidebarCheckbox = $('#newFormLinkIncludeInSidebar')
	
	var validator = $newPresetDialogForm.validate({
		rules: {
			newFormLinkNameInput: {
				minlength: 3,
				required: true,
				remote: {
					url: '/api/generic/stringValidation/validateItemLabel',
					data: {
						label: function() { return $presetNameInput.val(); }
					}
				} // remote
			}, // newFormNameInput
			newFormLinkFormSelection: { required:true }
		},
		messages: {
			newFormLinkNameInput: {
				required: "Link name is required"
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
		parentDatabaseID: pageContext.databaseID
	}	
	populateFormSelectionMenu(selectFormParams)
	
		
	$newPresetDialog.modal('show')
	
	initButtonClickHandler('#newFormLinkSaveButton',function() {
		console.log("New form link save button clicked")
		if($newPresetDialogForm.valid()) {	
			
			var newFormLinkParams = { 
				parentDatabaseID: pageContext.databaseID, 
				name: $presetNameInput.val(),
				formID: $formSelection.val(),
				includeInSidebar:$includeInSidebarCheckbox.prop("checked") }
								
			jsonAPIRequest("formLink/new",newFormLinkParams,function(newFormLinkInfo) {
				console.log("Created new form link: " + JSON.stringify(newFormLinkInfo))
				$newPresetDialog.modal('hide')
				
				var editPropsContentURL = '/admin/formLink/' + newFormLinkInfo.linkID
				setSettingsPageContent(editPropsContentURL,function() {
					initNewItemLinkPropsSettingsPageContent(pageContext,newFormLinkInfo)
				})
				
			})
			

		}
	})
	
}