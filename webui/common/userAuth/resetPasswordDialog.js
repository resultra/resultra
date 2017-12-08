
function openResetPasswordDialog() {
	
	
	var $resetDialog = $('#resetPasswordModalDialog')
	
	var $resetForm = $("#resetPasswordForm")
	var $emailInput = $('#resetPasswordEmailInput')
	
	var $resetAlert = $('#resetPasswordAlert')
	var $resetAlertMsg = $('#resetAlertMsg')
	var $resetInfo = $('#resetPasswordInfo')
	var $resetConfirm = $('#resetPasswordConfirm')
	var $resetButton = $('#resetPasswordButton')
	
	$resetAlert.hide()
	$resetInfo.show()
	$resetConfirm.hide()
	$resetButton.prop('disabled',false)
	
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
					}, 3000);
				} else {
					$resetAlertMsg.text(resetResp.msg)
					$resetAlert.show()
				}
			})
		}
		
	})
		
	$resetDialog.modal('show')
}