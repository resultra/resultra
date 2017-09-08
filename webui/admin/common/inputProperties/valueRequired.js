
function initValueRequiredValidationProperties(params) {
	
		initCheckboxChangeHandler('#' + params.elemPrefix+'ComponentValidationRequired', 
					params.initialValidationProps.valueRequired, function (newVal) {
		
			var validationProps = {
				valueRequired: newVal
			}
			params.setValidation(validationProps)		
		
		})
}
