$(document).ready(function() {
	
	function initNewUserButton() {
		var $newUserButton = $('#inviteNewUsersButton')
		initButtonControlClickHandler($newUserButton,function() {
			console.log("invite users button clicked")
			openUserInviteDialog()
		})
	}
	
	initWorkspaceAdminSettingsPageLayout($('#userMgmtAdminPage'))	
	
	initWorkspaceAdminPageHeader()
	
	initWorkspaceAdminSettingsTOC("settingsTOCUsers")
			
	initUserRegistrationProps()
	
	initNewUserButton()
})