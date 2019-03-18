// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

function userAuthSignoutCurrentUser() {
	
	var logoutParams = {}
	jsonRequest("/auth/signout",logoutParams,function(logoutResp) {
		// Navigate to the main home page. Since the user is now
		// signed out, the sign-in form will be presented.
		navigateToURL("/")
	})
}