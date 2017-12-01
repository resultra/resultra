function openRegisterUserDialog() {
	
	var $registerForm = $("#registerNewUserForm")
	var $userNameInput = $('#registerDialogUserNameInput')
	var $emailInput = $('#registerDialogEmailInput')
	var $passwordInput = $('#registerDialogPasswordInput')
	var $firstNameInput = $('#registerDialogFirstNameInput')
	var $lastNameInput = $('#registerDialogLastNameInput')
	
	
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
			registerDialogLastNameInput: {
				required: true,
				remote: {
					url: '/auth/validateName',
					data: {
						name: function() { return $lastNameInput.val(); }
					}
				} // remote
			},
			
			registerDialogUserNameInput: {
				minlength: 6,
				required: true,
				remote: {
					url: '/auth/validateNewUserName',
					data: {
						userName: function() { return $userNameInput.val(); }
					}
				} // remote
			},
			registerDialogEmailInput: {
				minlength: 3,
				required: true,
				remote: {
					url: '/auth/validateNewUserEmail',
					data: {
						emailAddr: function() { return  $emailInput.val(); }
					}
				} // remote
			},
			
			registerDialogPasswordInput: {
				minlength: 8,
				required: true,
				remote: {
					url: '/auth/validatePasswordStrength',
					data: {
						password: function() { return $passwordInput.val(); }
					}
				} // remote
			 },
			
			registerDialogRepeatPasswordInput: {
				required: true,
				equalTo: "#registerDialogPasswordInput"
			}
		},
		messages: {
			registerDialogFirstNameInput: {
				required: "First name is required",
				remote:"Name can only include include letters, periods, hyphens and apostrophes."
			},
			registerDialogLastNameInput: {
				required: "Last name is required",
				remote:"Name can only include include letters, periods, hyphens and apostrophes."
			},
			registerDialogUserNameInput: {
				required: "User name is required",
				remote:"This user name is already taken. Please choose another user name."
			},
			registerDialogEmailInput: {
				required: "Email address is required",
				remote:"This email address is already registered."
			},
			registerDialogPasswordInput: {
				required: "Password is required.",
				remote:"Passwords must be at least 8 characters, and include some numbers, upper and lower case lettters and/or symbols."
			},
			registerDialogRepeatPasswordInput: {
				required: "Password is required.",
				equalTo:"Passwords must match."
			}
		}
	})	
	
	$('#registerDialogAlert').hide()
	
	$('#registerDialogRegisterButton').unbind("click")
	$('#registerDialogRegisterButton').click(function(e) {
		
	    console.log("Register button clicked")
		$(this).blur();
	    e.preventDefault();// prevent the default anchor functionality
		
		if($registerForm.valid()) {
				
			var registerParams = {
				firstName: $firstNameInput.val(),
				lastName: $lastNameInput.val(),
				userName: $userNameInput.val(),
				emailAddr: $emailInput.val(),
				password: $passwordInput.val(),
			}
			jsonRequest("/auth/register",registerParams,function(registerResp) {
				if(registerResp.success == true) {
					$('#registerModalDialog').modal('hide')
				
				} else {
					$('#registerDialogAlert').text(registerResp.msg)
					$('#registerDialogAlert').show()
				}
			})
		}
		
	});
	
	$('#registerModalDialog').modal('show')
}