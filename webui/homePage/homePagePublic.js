$(document).ready(function() {	

	$('#signInButton').click(function(e) {
	    console.log("Sign in button clicked")
		openSigninDialog()
		$(this).blur();
	    e.preventDefault();// prevent the default anchor functionality
	});

	$('#homePageRegisterButton').click(function(e) {
	    console.log("Register button clicked")
		openRegisterUserDialog()
		$(this).blur();
	    e.preventDefault();// prevent the default anchor functionality
	});

	
}); // document ready
