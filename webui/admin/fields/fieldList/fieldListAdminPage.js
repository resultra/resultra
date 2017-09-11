$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#fieldAdminPage'))	
	initUserDropdownMenu()
	initAlertHeader(fieldListAdminContext.databaseID)
	initAdminSettingsTOC(fieldListAdminContext.databaseID)
		
	initAdminFieldSettings(fieldListAdminContext.databaseID)
})