$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#alertListAdminPage'))	
	
	initAdminPageHeader()
	
	initAdminSettingsTOC(alertListAdminContext.databaseID,"settingsTOCAlerts",alertListAdminContext.isSingleUserWorkspace)
			
	initAdminAlertSettings(alertListAdminContext.databaseID)
	
	appendPageSpecificBreadcrumbHeader("/admin/alerts/"+alertListAdminContext.databaseID,"Alerts")
	
	
})