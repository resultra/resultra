$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#alertListAdminPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(alertListAdminContext.databaseID)
			
	initAdminAlertSettings(alertListAdminContext.databaseID)
	
})