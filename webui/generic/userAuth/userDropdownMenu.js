function initUserDropdownMenu() {
	
	
	$('#userDropdownMenuSignoutMenuItem').click(function(e) {
	    console.log("Sign out button clicked")
	    e.preventDefault();// prevent the default anchor functionality
		
		userAuthSignoutCurrentUser()
	})
}