
function initUserTagValidationProperties(params) {
	
		initCheckboxChangeHandler('#adminUserTagComponentValidationRequired', 
					params.valueRequired, function (newVal) {
		
			var validationProps = {
				valueRequired: newVal
			}
			params.setValidation(validationProps)		
		
		})
}
