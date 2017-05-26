

function initTextInputValidationProperties(params) {
	
		initCheckboxChangeHandler('#adminTextInputComponentValidationRequired', 
					params.valueRequired, function (newVal) {
		
			var validationProps = {
				valueRequired: newVal
			}
			params.setValidation(validationProps)		
		
		})
}
