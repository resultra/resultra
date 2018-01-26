
$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#globalAdminPage'))	
	initAdminPageHeader()
	initAdminSettingsTOC(globalAdminPageContext.databaseID,"settingsTOCGlobals", globalAdminPageContext.isSingleUserWorkspace)
			
	initAdminGlobals(globalAdminPageContext.databaseID)
	
})