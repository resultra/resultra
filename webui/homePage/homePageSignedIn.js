$(document).ready(function() {	
	$('#signOutButton').click(function(e) {
	    console.log("Sign out button clicked")
		$(this).blur();
	    e.preventDefault();// prevent the default anchor functionality
		userAuthSignoutCurrentUser()
	});
	
}); // document ready
