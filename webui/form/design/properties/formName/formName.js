


function initFormPropertiesFormName(formInfo) {
	
	$('#formPropsFormNameInput').val(formInfo.name)
	
	var $formNameForm = $('#formNamePropertyPanelForm')
	
	var validationSettings = createInlineFormValidationSettings({
		rules: {
			formPropsFormNameInput: {
				minlength: 3,
				required: true,
				remote: {
					url: '/api/frm/validateFormName',
					data: {
						formID: function() { return formInfo.formID },
						formName: function() { return $('#formPropsFormNameInput').val(); }
					}
				} // remote
			} // newRoleNameInput
		},
	})	
	var validator = $formNameForm.validate(validationSettings)
	
	$('#formPropsFormNameInput').unbind("blur")
	$('#formPropsFormNameInput').blur(function() {
		if(validator.element('#formPropsFormNameInput')) {
			
			var newFormName = $('#formPropsFormNameInput').val()
			
			console.log("Starting form name change (pending remote validation): " + newFormName)
			
			var validationParams = { formID: formInfo.formID, formName: newFormName }
			doubleCheckRemoteFormValidation('/api/frm/validateFormName',validationParams, function(validationResult) {
				
				if(validationResult == true) {
					console.log("Changing form name: " + newFormName)
					jsonAPIRequest("frm/setName",{formID:formID,newFormName:newFormName},function(formInfo) {
						console.log("Done changing form name: " + newFormName)
					})
				} else {
					console.log("Remote validation failed, aborting form name change: " + newFormName)
				}
			})
			
		}
	})	
	

	validator.resetForm()
	
}