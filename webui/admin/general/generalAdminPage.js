
$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#generalAdminPage'))	
	initUserDropdownMenu()
	initAlertHeader(generalAdminPageContext.databaseID)
	initAdminSettingsTOC(generalAdminPageContext.databaseID)
			
	initAdminGeneralProperties(generalAdminPageContext.databaseID)
	
})