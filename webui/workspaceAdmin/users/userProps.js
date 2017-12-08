$(document).ready(function() {
	
	function initAcctActiveProp(userInfo) {
		var $acctActiveCheckbox = $('#userAccountActive')
		initCheckboxControlChangeHandler($acctActiveCheckbox, 
					userInfo.isActive, function (newVal) {
			var params = { 
				userID: userPropsContext.userID,
				isActive: newVal
			 }
			jsonRequest("/auth/setUserActive",params,function(resp) {
			})
		})
		
	}
	
	
	initWorkspaceAdminSettingsPageLayout($('#userPropsPage'))	
	
	initWorkspaceAdminPageHeader()
	
	initWorkspaceAdminSettingsTOC("settingsTOCUsers")
	
	var getUserInfoParams = { userID: userPropsContext.userID }
	jsonRequest("/auth/getAdminUserInfo",getUserInfoParams,function(userInfoResp) {
		initAcctActiveProp(userInfoResp)
	})
	
})