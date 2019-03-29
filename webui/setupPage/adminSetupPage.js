// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

function initAdminRegistrationPage() {
	
	var $registerForm = $("#adminRegistrationForm")
	var $userNameInput = $('#adminUserNameInput')
	var $passwordInput = $('#registerPasswordInput')
	var $firstNameInput = $('#adminFirstNameInput')
	var $lastNameInput = $('#adminLastNameInput')
	var $registerControls = $('.registerControls')
	var $passwordInput = $('#adminPasswordInput')

	
	var validator = $registerForm.validate({
		rules: {
			adminFirstNameInput: {
				required: true,
				remote: {
					url: '/auth/validateName',
					data: {
						name: function() { return $firstNameInput.val(); }
					}
				} // remote
			},
			adminLastNameInput: {
				required: true,
				remote: {
					url: '/auth/validateName',
					data: {
						name: function() { return $lastNameInput.val(); }
					}
				} // remote
			},
			
			adminUserNameInput: {
				minlength: 6,
				required: true,
				remote: {
					url: '/auth/validateNewUserName',
					data: {
						userName: function() { return $userNameInput.val(); }
					}
				} // remote
			},
			adminPasswordInput: {
				minlength: 8,
				required: true,
				remote: {
					url: '/auth/validatePasswordStrength',
					data: {
						password: function() { return $passwordInput.val(); }
					}
				} // remote
			 },
			adminPasswordInputRepeat: {
				required: true,
				equalTo: "#adminPasswordInput"
			}
						
		},
		messages: {
			adminFirstNameInput: {
				required: "First name is required",
				remote:"Name can only include include letters, periods, hyphens and apostrophes."
			},
			adminLastNameInput: {
				required: "Last name is required",
				remote:"Name can only include include letters, periods, hyphens and apostrophes."
			},
			adminUserNameInput: {
				required: "User name is required",
				remote:"This user name is already taken. Please choose another user name."
			},
			adminPasswordInput: {
				required: "Password is required.",
				remote:"Passwords must be at least 8 characters, and include some numbers, upper and lower case lettters and/or symbols."
			},
			adminPasswordInputRepeat: {
				required: "Password is required.",
				equalTo:"Passwords must match."
			}
		}
	})	
	
	
	var $registerButton = $('#adminPersonalizeButton')
		
	var $errorAlert = $('#adminRegisterErrorAlert')
	
	$errorAlert.hide()
	$registerControls.show()
		
	initButtonControlClickHandler($registerButton,function() {
		
		if($registerForm.valid()) {
				
			var registerParams = {
				firstName: $firstNameInput.val(),
				lastName: $lastNameInput.val(),
				userName: $userNameInput.val(),
			}
			/* TBD
			jsonRequest("/auth/registerSingleUser",registerParams,function(registerResp) {
				if(registerResp.success == true) {
					$registerControls.hide()
					// Navigate to the main page. Now that the registration is complete, the 
					// list of empty trackers will be shown.
					navigateToURL("/")
				} else {
					$errorAlert.text(registerResp.msg)
					$errorAlert.show()
				}
			})
			*/
		}
		
	})
	
}


$(document).ready(function() {	

	initPublicPageHeader()
	
	initAdminRegistrationPage()
	
}); // document ready