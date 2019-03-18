// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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