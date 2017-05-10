$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#formAdminPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(formsAdminContext.databaseID)
			
	initAdminFormSettings(formsAdminContext.databaseID)
	
})