// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initWorkspaceNameProperty(workspaceName) {
		
	var $nameForm = $('#workspaceNamePropertyForm')
	var $nameInput = $('#workspacePropsNameInput')
	
	$nameInput.val(workspaceName)
		
	var remoteValidationParams = {
		url: '/api/generic/stringValidation/validateItemLabel',
		data: {
			label: function() { return $nameInput.val() }
		}	
	}
	
	var validationSettings = createInlineFormValidationSettings({
		rules: {
			workspacePropsNameInput: {
				minlength: 3,
				required: true,
				remote: remoteValidationParams
			} // newRoleNameInput
		}
	})	
	
	var validator = $nameForm.validate(validationSettings)
	
	initInlineInputValidationOnBlur(validator,'#workspacePropsNameInput',
		remoteValidationParams, function(validatedName) {		
			var setNameParams = {
				newName:validatedName
			}
			jsonAPIRequest("workspace/setName",setNameParams,function(dbInfo) {
				console.log("Done changing workspace name: " + validatedName)
			})
	})	

	validator.resetForm()
	
}
