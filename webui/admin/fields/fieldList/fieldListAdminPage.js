$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#fieldAdminPage'))	
	initAdminPageHeader()
	initAdminSettingsTOC(fieldListAdminContext.databaseID,"settingsTOCFields")
		
	initAdminFieldSettings(fieldListAdminContext.databaseID)
})