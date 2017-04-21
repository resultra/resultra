function validateFormValues($parentFormLayout, validationCompleteCallback) {
	
	var validationFuncs = []
	
	$parentFormLayout.find(".layoutContainer").each(function() {
		var viewFormConfig = $(this).data("viewFormConfig")
		if(viewFormConfig.hasOwnProperty('validateValue')) {
			validationFuncs.push(viewFormConfig.validateValue)
		}
	})
	
	if (validationFuncs.length > 0) {
		
		// The validations of individual components may use remote AJAX calls to validate
		// their inputs. validationsRemaining is used as a counter for remaining validations
		// to complete. When this counter hits zero, the validation is complete.
		var validationsRemaining = validationFuncs.length
		var validationSucceeded = true
		
		
		for(var currValIndex = 0; currValIndex < validationFuncs.length; currValIndex++) {
			var currValidator = validationFuncs[currValIndex]
			currValidator(function(validationResult) {
				if (validationResult === false) {
					// The validation of any single component makes the entire
					// validation fail
					validationSucceeded = false 
				}
				validationsRemaining--
				if(validationsRemaining <= 0) {
					validationCompleteCallback(validationSucceeded)
				}
			})
		}
	} else {
		validationCompleteCallback(true)
	}
}