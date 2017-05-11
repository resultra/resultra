$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#formLinkAdminPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(formLinkAdminContext.databaseID)
			
	initAdminFormLinkSettings(formLinkAdminContext.databaseID)
	
})