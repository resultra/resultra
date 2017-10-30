
$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#userAdminPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(userAdminPageContext.databaseID,"settingsTOCUsers")
	

	initUserListSettings(userAdminPageContext.databaseID)

				
})