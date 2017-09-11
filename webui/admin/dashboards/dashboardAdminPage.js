$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#dashboardAdminPage'))	
	initUserDropdownMenu()
	initAlertHeader(dashboardAdminPageContext.databaseID)
	initAdminSettingsTOC(dashboardAdminPageContext.databaseID)
			
	initAdminDashboardSettings(dashboardAdminPageContext.databaseID)
	
})