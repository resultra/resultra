
function userAuthSignoutCurrentUser() {
	
	var logoutParams = {}
	jsonRequest("/auth/signout",logoutParams,function(logoutResp) {
		// Reload the current page, while forcing the page to refresh
		// from the server, rather than the cache. This will cause
		// the server to authenticate the current user while serving
		// up the current page.
		location.reload(true)
		
	})
}