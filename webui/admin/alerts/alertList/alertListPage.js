$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#alertListAdminPage'))	
	
	initAdminPageHeader()
	
	initAdminSettingsTOC(alertListAdminContext.databaseID,"settingsTOCAlerts")
			
	initAdminAlertSettings(alertListAdminContext.databaseID)
	
	appendPageSpecificBreadcrumbHeader("/admin/alerts/"+alertListAdminContext.databaseID,"Alerts")
	
	
})