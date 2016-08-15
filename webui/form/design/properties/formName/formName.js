


function initFormPropertiesFormName(formInfo) {
	
	$('#formPropsFormNameInput').val(formInfo.name)
	
	var $formNameForm = $('#formNamePropertyPanelForm')
	
	var validator = $formNameForm.validate({
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
		// Since there is already a value in place, the validation can be made "eager" by 
		// always triggering the validation when the key goes up. The default behavior is
		// to only trigger the first validation when the input focus is lost (blur), then
		// become more eager when an error is detected.
		onkeyup: function(element) { $(element).valid() }
	})
	
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