$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#alertListAdminPage'))	
	
	initUserDropdownMenu()
	initAlertHeader(alertListAdminContext.databaseID)
	
	initAdminSettingsTOC(alertListAdminContext.databaseID)
			
	initAdminAlertSettings(alertListAdminContext.databaseID)
	
})