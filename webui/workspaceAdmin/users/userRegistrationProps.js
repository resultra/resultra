function initUserRegistrationProps() {
	var infoParams = {}
	jsonRequest("/auth/getAllUsersInfo",infoParams,function(usersInfo) {
		
		var $usersTableBody = $('#registeredUserListTableBody')
		
		for(var userIndex = 0; userIndex < usersInfo.length; userIndex++) {
			var currUserInfo = usersInfo[userIndex]
		
			var $userInfo = $('#registerUserListItemTemplate').clone()
			$userInfo.removeAttr('id')
			$userInfo.find('.firstName').text(currUserInfo.firstName)
			$userInfo.find('.lastName').text(currUserInfo.lastName)
			$userInfo.find('.emailAddress').text(currUserInfo.emailAddress)
			$userInfo.find('.userName').text(currUserInfo.userName)
			
			$usersTableBody.append($userInfo)
			
		}
		
	})
	
}