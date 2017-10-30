$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#alertListAdminPage'))	
	
	initUserDropdownMenu()
	
	initAdminSettingsTOC(alertListAdminContext.databaseID,"settingsTOCAlerts")
			
	initAdminAlertSettings(alertListAdminContext.databaseID)
	
})