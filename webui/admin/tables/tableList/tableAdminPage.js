$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#tableAdminPage'))	
	initAdminPageHeader()
	initAdminSettingsTOC(tableAdminContext.databaseID,"settingsTOCTables")
			
	initAdminTableListSettings(tableAdminContext.databaseID)
	
})