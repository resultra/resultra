function createInlineFormValidationSettings(specificRules) {
	var settings = {
        highlight: function(element) {
            $(element).closest('.form-group').removeClass('has-success has-feedback').addClass('has-error has-feedback');
			$(element).closest('.form-group').find('i.glyphicon').remove();
			$(element).closest('.form-group').append('<i class="form-control-feedback glyphicon glyphicon-exclamation-sign"></i>');
        },
        unhighlight: function(element) {
            $(element).closest('.form-group').removeClass('has-error  has-feedback').addClass('has-success  has-feedback');
			$(element).closest('.form-group').find('i.glyphicon').remove();
			// Don't show the OK sign when form is being validated in an inline context
        },
        errorElement: 'span',
        errorClass: 'help-block',
        errorPlacement: function(error, element) {
            if(element.parent('.input-group').length) {
                error.insertAfter(element.parent());
            } else {
                error.insertAfter(element);
            }
		},
		// Since there is already a value in place, the validation can be made "eager" by 
		// always triggering the validation when the key goes up. The default behavior is
		// to only trigger the first validation when the input focus is lost (blur), then
		// become more eager when an error is detected.
		onkeyup: function(element) { $(element).valid() }
		
	}
	$.extend(settings,specificRules)
	
	return settings
}

function validateRemoteInlineInput(validator, inputSelector,
	remoteValidationParams, validationSucceedFunc) {
		if(validator.element(inputSelector)) {
		
			var newVal = $(inputSelector).val()
					
			doubleCheckRemoteFormValidation(remoteValidationParams.url,remoteValidationParams.data, 
						function(validationResult) {
			
				if(validationResult == true) {
					validationSucceedFunc(newVal)
				} else {
					console.log("Remote validation failed: " + remoteValidationParams.url)
				}
			})
		
		}
		
	}


function validateRemoteInlineInputControl(validator, $input,
		remoteValidationParams, validationSucceedFunc) {
	if(validator.element($input)) {

		var newVal = $input.val()
			
		doubleCheckRemoteFormValidation(remoteValidationParams.url,remoteValidationParams.data, 
					function(validationResult) {
	
			if(validationResult == true) {
				validationSucceedFunc(newVal)
			} else {
				console.log("Remote validation failed: " + remoteValidationParams.url)
			}
		})

	}
		
}


function initInlineInputControlValidationOnBlur(validator, $input,
				remoteValidationParams, validationSucceedFunc) {
	
	$input.unbind("blur")
	$input.blur(function() {
		validateRemoteInlineInputControl(validator, $input,
			remoteValidationParams, validationSucceedFunc)
	})	
	
}

function initInlineInputValidationOnBlur(validator, inputSelector,
				remoteValidationParams, validationSucceedFunc) {
	
	initInlineInputControlValidationOnBlur(validator, $(inputSelector),
				remoteValidationParams, validationSucceedFunc)	
}
