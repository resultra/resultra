$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#itemListAdminPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(itemListAdminContext.databaseID,"settingsTOCLists")
			
	initAdminListSettings(itemListAdminContext.databaseID)
	
})