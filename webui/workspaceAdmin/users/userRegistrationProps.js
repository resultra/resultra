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