function initUserDropdownMenu() {
	
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