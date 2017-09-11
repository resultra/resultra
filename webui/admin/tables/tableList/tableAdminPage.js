$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#tableAdminPage'))	
	initUserDropdownMenu()
	initAlertHeader(tableAdminContext.databaseID)
	initAdminSettingsTOC(tableAdminContext.databaseID)
			
	initAdminTableListSettings(tableAdminContext.databaseID)
	
})