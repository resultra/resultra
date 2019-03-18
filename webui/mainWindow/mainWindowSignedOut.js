// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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
