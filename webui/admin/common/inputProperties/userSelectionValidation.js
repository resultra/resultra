

function initUserSelectionValidationProperties(params) {
	
		initCheckboxChangeHandler('#adminUserSelectionComponentValidationRequired', 
					params.valueRequired, function (newVal) {
		
			var validationProps = {
				valueRequired: newVal
			}
			params.setValidation(validationProps)		
		
		})
}
