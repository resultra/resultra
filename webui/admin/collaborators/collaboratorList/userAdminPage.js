
$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#userAdminPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(userAdminPageContext.databaseID)
	
	initAlertHeader(userAdminPageContext.databaseID)

	initUserListSettings(userAdminPageContext.databaseID)

				
})