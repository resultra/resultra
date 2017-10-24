
$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#globalAdminPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(globalAdminPageContext.databaseID)
			
	initAdminGlobals(globalAdminPageContext.databaseID)
	
})