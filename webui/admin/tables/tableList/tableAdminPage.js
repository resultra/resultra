$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#tableAdminPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(tableAdminContext.databaseID)
			
	initAdminTableListSettings(tableAdminContext.databaseID)
	
})