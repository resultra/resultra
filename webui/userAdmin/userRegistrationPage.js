// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initInviteeRegistrationPage() {
	
	var $registerForm = $("#registerNewUserForm")
	var $userNameInput = $('#registerUserNameInput')
	var $passwordInput = $('#registerPasswordInput')
	var $firstNameInput = $('#registerFirstNameInput')
	var $lastNameInput = $('#registerLastNameInput')
	var $registerControls = $('.registerControls')
	
	var validator = $registerForm.validate({
		rules: {
			registerDialogFirstNameInput: {
				required: true,
				remote: {
					url: '/auth/validateName',
					data: {
						name: function() { return $firstNameInput.val(); }
					}
				} // remote
			},
			registerLastNameInput: {
				required: true,
				remote: {
					url: '/auth/validateName',
					data: {
						name: function() { return $lastNameInput.val(); }
					}
				} // remote
			},
			
			registerUserNameInput: {
				minlength: 6,
				required: true,
				remote: {
					url: '/auth/validateNewUserName',
					data: {
						userName: function() { return $userNameInput.val(); }
					}
				} // remote
			},
			
			registerPasswordInput: {
				minlength: 8,
				required: true,
				remote: {
					url: '/auth/validatePasswordStrength',
					data: {
						password: function() { return $passwordInput.val(); }
					}
				} // remote
			 },
			
			registerPasswordInputRepeat: {
				required: true,
				equalTo: "#registerPasswordInput"
			}
		},
		messages: {
			registerFirstNameInput: {
				required: "First name is required",
				remote:"Name can only include include letters, periods, hyphens and apostrophes."
			},
			registerLastNameInput: {
				required: "Last name is required",
				remote:"Name can only include include letters, periods, hyphens and apostrophes."
			},
			registerUserNameInput: {
				required: "User name is required",
				remote:"This user name is already taken. Please choose another user name."
			},
			registerPasswordInput: {
				required: "Password is required.",
				remote:"Passwords must be at least 8 characters, and include some numbers, upper and lower case lettters and/or symbols."
			},
			registerPasswordInputRepeat: {
				required: "Password is required.",
				equalTo:"Passwords must match."
			}
		}
	})	
	
	
	var $registerButton = $('#userRegisterButton')
		
	var $confirmAlert = $('#resetPasswordConfirm')
	var $errorAlert = $('#registrationErrorAlert')
	
	$confirmAlert.hide()
	$errorAlert.hide()
	$registerControls.show()
		
	initButtonControlClickHandler($registerButton,function() {
		
		if($registerForm.valid()) {
				
			var registerParams = {
				firstName: $firstNameInput.val(),
				lastName: $lastNameInput.val(),
				userName: $userNameInput.val(),
				emailAddr: registrationContext.inviteeEmailAddr,
				password: $passwordInput.val(),
			}
			jsonRequest("/auth/register",registerParams,function(registerResp) {
				if(registerResp.success == true) {
					$registerControls.hide()
					$confirmAlert.show()
				} else {
					$errorAlert.text(registerResp.msg)
					$errorAlert.show()
				}
			})
		}
		
	})
	
}


$(document).ready(function() {	

	initPublicPageHeader()
	
	initInviteeRegistrationPage()
	
}); // document ready