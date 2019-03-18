// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initUserRegistrationProps() {
	var infoParams = {}
	
	jsonRequest("/auth/getAllUsersInfo",infoParams,function(usersInfo) {
		
		var $usersTableBody = $('#registeredUserListTableBody')
		
		function addUserInfoToTable(userInfo) {
		
			var $userInfo = $('#registerUserListItemTemplate').clone()
			$userInfo.removeAttr('id')
			
			$userInfo.find('.firstName').text(userInfo.firstName)
			$userInfo.find('.lastName').text(userInfo.lastName)
			$userInfo.find('.emailAddress').text(userInfo.emailAddress)
			$userInfo.find('.userName').text(userInfo.userName)
			
			var editUserLink = "/workspace-admin/user/"+userInfo.userID
			$userInfo.find('.editPropsButton').attr("href",editUserLink)
						
			$usersTableBody.append($userInfo)
			
		}
		
		for(var userIndex = 0; userIndex < usersInfo.length; userIndex++) {
			var currUserInfo = usersInfo[userIndex]
			addUserInfoToTable(currUserInfo)
		}
		
	})
	
}