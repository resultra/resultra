// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initUserDropdownMenu(isSingleUserWorkspace) {
	
	if (!isSingleUserWorkspace) {
		var getUserInfoParams = {}
		jsonRequest("/auth/getCurrentUserInfo",getUserInfoParams,function(userInfo) {
			var fullName = userInfo.firstName + " " + userInfo.lastName
			$('#userMenuUserName').text(fullName)
		})
	
		$('#userDropdownMenuSignoutMenuItem').click(function(e) {
		    console.log("Sign out button clicked")
		    e.preventDefault();// prevent the default anchor functionality
		
			userAuthSignoutCurrentUser()
		})
	}
	
}