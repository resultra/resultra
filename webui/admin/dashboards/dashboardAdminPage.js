$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#dashboardAdminPage'))	
	initAdminPageHeader()
	initAdminSettingsTOC(dashboardAdminPageContext.databaseID,"settingsTOCDashboards")
			
	initAdminDashboardSettings(dashboardAdminPageContext.databaseID)
	
})