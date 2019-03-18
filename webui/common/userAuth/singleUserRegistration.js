// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initSingleUserRegistrationPagePage() {
	
	var $registerForm = $("#singleUserRegistrationForm")
	var $userNameInput = $('#singlerUserUserNameInput')
	var $passwordInput = $('#registerPasswordInput')
	var $firstNameInput = $('#singlerUserFirstNameInput')
	var $lastNameInput = $('#singlerUserLastNameInput')
	var $registerControls = $('.registerControls')
	
	var validator = $registerForm.validate({
		rules: {
			singlerUserFirstNameInput: {
				required: true,
				remote: {
					url: '/auth/validateName',
					data: {
						name: function() { return $firstNameInput.val(); }
					}
				} // remote
			},
			singlerUserLastNameInput: {
				required: true,
				remote: {
					url: '/auth/validateName',
					data: {
						name: function() { return $lastNameInput.val(); }
					}
				} // remote
			},
			
			singlerUserUserNameInput: {
				minlength: 6,
				required: true,
				remote: {
					url: '/auth/validateNewUserName',
					data: {
						userName: function() { return $userNameInput.val(); }
					}
				} // remote
			}
						
		},
		messages: {
			singlerUserFirstNameInput: {
				required: "First name is required",
				remote:"Name can only include include letters, periods, hyphens and apostrophes."
			},
			singlerUserLastNameInput: {
				required: "Last name is required",
				remote:"Name can only include include letters, periods, hyphens and apostrophes."
			},
			singlerUserUserNameInput: {
				required: "User name is required",
				remote:"This user name is already taken. Please choose another user name."
			}
		}
	})	
	
	
	var $registerButton = $('#singleUserPersonalizeButton')
		
	var $errorAlert = $('#singleUserRegisterErrorAlert')
	
	$errorAlert.hide()
	$registerControls.show()
		
	initButtonControlClickHandler($registerButton,function() {
		
		if($registerForm.valid()) {
				
			var registerParams = {
				firstName: $firstNameInput.val(),
				lastName: $lastNameInput.val(),
				userName: $userNameInput.val(),
			}
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
		}
		
	})
	
}