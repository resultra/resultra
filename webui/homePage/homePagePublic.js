$(document).ready(function() {	

	initUserSigninComponents()

	$('#homePageRegisterButton').click(function(e) {
	    console.log("Register button clicked")
		openRegisterUserDialog()
		$(this).blur();
	    e.preventDefault();// prevent the default anchor functionality
	});

	
}); // document ready
