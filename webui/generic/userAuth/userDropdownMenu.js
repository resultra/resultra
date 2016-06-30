function initUserDropdownMenu() {
	
	var getUserInfoParams = {}
	jsonRequest("/auth/getUserInfo",getUserInfoParams,function(userInfo) {
		$('#userMenuUserName').text(userInfo.userName)
	})
	
	$('#userDropdownMenuSignoutMenuItem').click(function(e) {
	    console.log("Sign out button clicked")
	    e.preventDefault();// prevent the default anchor functionality
		
		userAuthSignoutCurrentUser()
	})
}