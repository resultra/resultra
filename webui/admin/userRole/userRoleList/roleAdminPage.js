$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#roleAdminPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(roleAdminContext.databaseID,"settingsTOCRoles")
			
	initUserRoleSettings(roleAdminContext.databaseID)
	
})