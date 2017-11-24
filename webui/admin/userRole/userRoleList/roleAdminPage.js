$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#roleAdminPage'))	
	initAdminPageHeader()
	initAdminSettingsTOC(roleAdminContext.databaseID,"settingsTOCRoles")
			
	initUserRoleSettings(roleAdminContext.databaseID)
	
})