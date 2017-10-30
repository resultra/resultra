
$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#globalAdminPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(globalAdminPageContext.databaseID,"settingsTOCGlobals")
			
	initAdminGlobals(globalAdminPageContext.databaseID)
	
})