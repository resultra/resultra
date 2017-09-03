
function initLabelValidationProperties(params) {
	
		initCheckboxChangeHandler('#adminLabelComponentValidationRequired', 
					params.valueRequired, function (newVal) {
		
			var validationProps = {
				valueRequired: newVal
			}
			params.setValidation(validationProps)		
		
		})
}
