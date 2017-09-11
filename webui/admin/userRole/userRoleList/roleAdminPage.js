$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#roleAdminPage'))	
	initUserDropdownMenu()
	initAlertHeader(roleAdminContext.databaseID)
	initAdminSettingsTOC(roleAdminContext.databaseID)
			
	initUserRoleSettings(roleAdminContext.databaseID)
	
})