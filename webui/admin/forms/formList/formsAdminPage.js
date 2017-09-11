$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#formAdminPage'))	
	initUserDropdownMenu()
	initAlertHeader(formsAdminContext.databaseID)
	initAdminSettingsTOC(formsAdminContext.databaseID)
			
	initAdminFormSettings(formsAdminContext.databaseID)
	
})