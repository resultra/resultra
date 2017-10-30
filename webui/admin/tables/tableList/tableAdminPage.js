$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#tableAdminPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(tableAdminContext.databaseID,"settingsTOCTables")
			
	initAdminTableListSettings(tableAdminContext.databaseID)
	
})