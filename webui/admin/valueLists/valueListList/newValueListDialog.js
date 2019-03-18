// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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
	
	var listCreated = false
	
	initButtonClickHandler('#newValueListSaveButton',function() {
		console.log("New value list save button clicked")
		if($newValueListForm.valid()) {	
			
			var newValueListParams = { 
				parentDatabaseID: databaseID, 
				name: $valueListNameInput.val(),
				valueType: $valueTypeSelection.val() }
						
			// Only support the creation of a single value list from the dialog.
			if (listCreated === false) {
				listCreated = true
				jsonAPIRequest("valueList/new",newValueListParams,function(newValueListInfo) {
					console.log("Created new value list: " + JSON.stringify(newValueListInfo))
					$newValueListDialog.modal('hide')
			
					var editPropsContentURL = '/admin/valueList/' + newValueListInfo.valueListID
					setSettingsPageContent(editPropsContentURL,function() {
						initValueListSettingsPageContent(newValueListInfo)
					})
				})
				
			}
			

		}
	})
	
}