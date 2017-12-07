function initWorkspacePermissionSettings(workspaceInfo) {
	
	
	var $usersCanRegisterAccountsCheckbox = $('#usersCanRegisterAccounts')
	initCheckboxControlChangeHandler($usersCanRegisterAccountsCheckbox, 
				workspaceInfo.properties.allowUserRegistration, function (newVal) {
		var props = { allowUserRegistration: newVal }
		jsonAPIRequest("workspace/setAllowUserRegistration",props,function(workspaceInfo) {
		})
	})
	
	
	
}