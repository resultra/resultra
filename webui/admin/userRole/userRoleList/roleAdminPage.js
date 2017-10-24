$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#roleAdminPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(roleAdminContext.databaseID)
			
	initUserRoleSettings(roleAdminContext.databaseID)
	
})