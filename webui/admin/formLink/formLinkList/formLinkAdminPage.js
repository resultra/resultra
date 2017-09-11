$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#formLinkAdminPage'))	
	initUserDropdownMenu()
	initAlertHeader(formLinkAdminContext.databaseID)
	initAdminSettingsTOC(formLinkAdminContext.databaseID)
			
	initAdminFormLinkSettings(formLinkAdminContext.databaseID)
	
})