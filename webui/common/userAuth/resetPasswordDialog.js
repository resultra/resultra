// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

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