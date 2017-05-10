$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#fieldAdminPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(fieldListAdminContext.databaseID)
		
	initAdminFieldSettings(fieldListAdminContext.databaseID)
})