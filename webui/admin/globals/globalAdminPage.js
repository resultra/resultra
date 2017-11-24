
$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#globalAdminPage'))	
	initAdminPageHeader()
	initAdminSettingsTOC(globalAdminPageContext.databaseID,"settingsTOCGlobals")
			
	initAdminGlobals(globalAdminPageContext.databaseID)
	
})