$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#itemListAdminPage'))	
	initAdminPageHeader()
	initAdminSettingsTOC(itemListAdminContext.databaseID,"settingsTOCLists")
			
	initAdminListSettings(itemListAdminContext.databaseID)
	
})