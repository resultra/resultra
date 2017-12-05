
function initPasswordResetForm() {
	var $resetAlert = $('#resetPasswordAlert')
	
	$resetAlert.hide()
	
	var $resetForm = $("#passwordResetForm")
	var $passwordInput = $('#resetPasswordInput')
	
	var validator = $resetForm.validate({
		rules: {	
			resetPasswordInput: {
				minlength: 8,
				required: true,
				remote: {
					url: '/auth/validatePasswordStrength',
					data: {
						password: function() { return $passwordInput.val(); }
					}
				} // remote
			 },
			
			resetPasswordInputRepeat: {
				required: true,
				equalTo: "#resetPasswordInput"
			}
		},
		messages: {
			resetPasswordInput: {
				required: "Password is required.",
				remote:"Passwords must be at least 8 characters, and include some numbers, upper and lower case lettters and/or symbols."
			},
			resetPasswordInputRepeat: {
				required: "Password is required.",
				equalTo:"Passwords must match."
			}
		}
	})	
	
	
	initButtonClickHandler('#passwordResetButton',function() {
		if($resetForm.valid()) {
				
			var resetParams = {
				resetID: resetContext.resetID,
				password: $passwordInput.val(),
			}
			jsonRequest("/auth/resetPassword",resetParams,function(resetResp) {
				if(resetResp.success == true) {
					
				} else {
					$resetAlert.text(registerResp.msg)
					$resetAlert.show()
				}
			})
		}
		
	})

	
}

$(document).ready(function() {	
	
	initPasswordResetForm()

	initPublicPageHeader()
	
}); // document ready