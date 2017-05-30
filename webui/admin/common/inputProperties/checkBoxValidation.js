


function initCheckBoxValidationProps(params) {
	
	initCheckboxChangeHandler('#adminCheckboxComponentValidationRequired', 
				params.initialValidationConfig.valueRequired, function (newVal) {
			
		var validationProps = {
			valueRequired: newVal
		}
		params.setValidation(validationProps)
			
	})
	
}