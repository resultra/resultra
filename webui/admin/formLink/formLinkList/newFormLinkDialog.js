// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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
	$includeInSidebarCheckbox.prop("checked",true)
	
	var selectFormParams = {
		menuSelector: formSelectionSelector,
		parentDatabaseID: pageContext.databaseID
	}	
	populateFormSelectionMenu(selectFormParams)
	
		
	$newPresetDialog.modal('show')
	
	var linkCreated = false
	
	initButtonClickHandler('#newFormLinkSaveButton',function() {
		console.log("New form link save button clicked")
		if($newPresetDialogForm.valid()) {	
			
			var newFormLinkParams = { 
				parentDatabaseID: pageContext.databaseID, 
				name: $presetNameInput.val(),
				formID: $formSelection.val(),
				includeInSidebar:$includeInSidebarCheckbox.prop("checked") }
						
				// Only create a single link when the new link button is pressed. Otherwise,
				// there's a possibility of creating more than one link with the same
				// properties from the same dialog.		
				if (linkCreated === false) {
					linkCreated = true
					jsonAPIRequest("formLink/new",newFormLinkParams,function(newFormLinkInfo) {
						console.log("Created new form link: " + JSON.stringify(newFormLinkInfo))
						$newPresetDialog.modal('hide')
				
						var editPropsContentURL = '/admin/formLink/' + newFormLinkInfo.linkID
						setSettingsPageContent(editPropsContentURL,function() {
							initNewItemLinkPropsSettingsPageContent(pageContext,newFormLinkInfo)
						})
				
					})
				}					
								
			

		}
	})
	
}