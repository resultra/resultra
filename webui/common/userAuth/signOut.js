
function userAuthSignoutCurrentUser() {
	
	var logoutParams = {}
	jsonRequest("/auth/signout",logoutParams,function(logoutResp) {
		// Navigate to the main home page. Since the user is now
		// signed out, the sign-in form will be presented.
		navigateToURL("/")
	})
}