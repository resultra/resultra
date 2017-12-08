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
	
	
	function initResetPasswordButton(userInfo) {
		var $resetButton = $('#resetUserPasswordButton')
		initButtonControlClickHandler($resetButton,function() {
			
			var resetParams = { userID: userInfo.userID }
			$resetButton.prop('disabled',true)
			jsonRequest("/auth/sendResetPasswordLinkByUserID",resetParams,function(resetResp) {
				$resetButton.prop('disabled',false)	
				
				var $resetSuccessAlert = $('#resetUserPasswordSuccess')	
				var $resetErrorAlert = $('#resetUserPasswordError')
				
				if(resetResp.success == true) {
					$resetSuccessAlert.show()
					setTimeout(function() {
						$resetSuccessAlert.hide()
					},3000)
				} else {
					$resetErrorAlert.text(resetResp.msg)
					$resetErrorAlert.show()
				}
				
			})
			
		})
	}
	
	
	
	initWorkspaceAdminSettingsPageLayout($('#userPropsPage'))	
	
	initWorkspaceAdminPageHeader()
	
	initWorkspaceAdminSettingsTOC("settingsTOCUsers")
	
	var getUserInfoParams = { userID: userPropsContext.userID }
	jsonRequest("/auth/getAdminUserInfo",getUserInfoParams,function(userInfoResp) {
		initAcctActiveProp(userInfoResp)
		initResetPasswordButton(userInfoResp)
	})
	
})