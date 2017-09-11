
$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#globalAdminPage'))	
	initUserDropdownMenu()
	initAlertHeader(globalAdminPageContext.databaseID)
	initAdminSettingsTOC(globalAdminPageContext.databaseID)
			
	initAdminGlobals(globalAdminPageContext.databaseID)
	
})