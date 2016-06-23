$(document).ready(function() {	

	$('#signInButton').click(function(e) {
	    console.log("Sign in button clicked")
		openSigninDialog()
		$(this).blur();
	    e.preventDefault();// prevent the default anchor functionality
	});
	
}); // document ready
