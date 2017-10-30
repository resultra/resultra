$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#fieldAdminPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(fieldListAdminContext.databaseID,"settingsTOCFields")
		
	initAdminFieldSettings(fieldListAdminContext.databaseID)
})