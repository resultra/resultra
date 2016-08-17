


function initFormPropertiesFormName(formInfo) {
	
	$('#formPropsFormNameInput').val(formInfo.name)
	
	var $formNameForm = $('#formNamePropertyPanelForm')
	
	var remoteValidationParams = {
		url: '/api/frm/validateFormName',
		data: {
			formID: function() { return formInfo.formID },
			formName: function() { return $('#formPropsFormNameInput').val(); }
		}
	}
	
	var validationSettings = createInlineFormValidationSettings({
		rules: {
			formPropsFormNameInput: {
				minlength: 3,
				required: true,
				remote:  remoteValidationParams
			}
		},
	})	
	var validator = $formNameForm.validate(validationSettings)
	
	
	initInlineInputValidationOnBlur(validator,'#formPropsFormNameInput',
			remoteValidationParams, function(validatedName) {
				jsonAPIRequest("frm/setName",{formID:formInfo.formID,newFormName:validatedName},function(formInfo) {
					console.log("Done changing form name: " + validatedName)
				})		
	})
		

	validator.resetForm()
	
}