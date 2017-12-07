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
