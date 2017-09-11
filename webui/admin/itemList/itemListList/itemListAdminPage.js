$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#itemListAdminPage'))	
	initUserDropdownMenu()
	initAlertHeader(itemListAdminContext.databaseID)
	initAdminSettingsTOC(itemListAdminContext.databaseID)
			
	initAdminListSettings(itemListAdminContext.databaseID)
	
})