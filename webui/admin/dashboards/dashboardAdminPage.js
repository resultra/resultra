$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#dashboardAdminPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(dashboardAdminPageContext.databaseID)
			
	initAdminDashboardSettings(dashboardAdminPageContext.databaseID)
	
})