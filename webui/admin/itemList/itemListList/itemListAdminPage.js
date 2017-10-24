$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#itemListAdminPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(itemListAdminContext.databaseID)
			
	initAdminListSettings(itemListAdminContext.databaseID)
	
})