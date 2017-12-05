
function openResetPasswordDialog() {
	
	
	var $resetDialog = $('#resetPasswordModalDialog')
	
	var $resetForm = $("#resetPasswordForm")
	var $emailInput = $('#resetPasswordEmailInput')
	
	var $resetAlert = $('#resetPasswordAlert')
	var $resetAlertMsg = $('#resetAlertMsg')
	var $resetInfo = $('#resetPasswordInfo')
	var $resetConfirm = $('#resetPasswordConfirm')
	
	$resetAlert.hide()
	$resetInfo.show()
	$resetConfirm.hide()
	
	$emailInput.val("")
	resetFormValidationFeedback($resetForm)
	
	var validator = $resetForm.validate({
		rules: {
			resetPasswordEmailInput: {
				minlength: 3,
				required: true,
				remote: {
					url: '/auth/validateExistingUserEmail',
					data: {
						emailAddr: function() { return  $emailInput.val(); }
					}
				} // remote
			},

		},
		messages: {
			resetPasswordEmailInput: {
				required: "Email address is required",
				remote:"Please enter the email address used to register with this system."
			},
		}
	})
	
	var $resetButton = $('#resetPasswordButton')
	
	initButtonClickHandler("#resetPasswordButton",function() {
		if($resetForm.valid()) {
				
			var resetParams = {
				emailAddr: $emailInput.val(),
			}
			$resetButton.prop('disabled',true)
			jsonRequest("/auth/sendResetPasswordLink",resetParams,function(resetResp) {
				if(resetResp.success == true) {
					$resetConfirm.show()
					$resetInfo.hide()
					setTimeout(function() {
						$resetDialog.modal('hide')
					}, 2000);
				} else {
					$resetAlertMsg.text(resetResp.msg)
					$resetAlert.show()
				}
			})
		}
		
	})
		
	$resetDialog.modal('show')
}