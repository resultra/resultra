
$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#userAdminPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(userAdminPageContext.databaseID)
	

	initUserListSettings(userAdminPageContext.databaseID)

				
})