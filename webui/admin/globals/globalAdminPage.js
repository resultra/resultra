
$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#globalAdminPage'))	
	initAdminPageHeader(globalAdminPageContext.isSingleUserWorkspace)
	initAdminSettingsTOC(globalAdminPageContext.databaseID,"settingsTOCGlobals", globalAdminPageContext.isSingleUserWorkspace)
			
	initAdminGlobals(globalAdminPageContext.databaseID)
	
})