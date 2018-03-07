$(document).ready(function() {	

	if(homePageContext.isSingleUserWorkspace) {
		initSingleUserRegistrationPagePage()	
	} else {
		initUserSigninComponents()

		$('#homePageRegisterButton').click(function(e) {
		    console.log("Register button clicked")
			openRegisterUserDialog()
			$(this).blur();
		    e.preventDefault();// prevent the default anchor functionality
		});
	
		initButtonClickHandler('#homePageForgotPasswordButton',function() {
			console.log("Reset password clicked")
			openResetPasswordDialog()
		})
		
	}

	initPublicPageHeader()

		
}); // document ready
