$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#dashboardAdminPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(dashboardAdminPageContext.databaseID,"settingsTOCDashboards")
			
	initAdminDashboardSettings(dashboardAdminPageContext.databaseID)
	
})