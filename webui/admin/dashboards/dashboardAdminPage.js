$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#dashboardAdminPage'))	
	initAdminPageHeader(dashboardAdminPageContext.isSingleUserWorkspace)
	initAdminSettingsTOC(dashboardAdminPageContext.databaseID,"settingsTOCDashboards",dashboardAdminPageContext.isSingleUserWorkspace)
			
	initAdminDashboardSettings(dashboardAdminPageContext.databaseID)
	
	appendPageSpecificBreadcrumbHeader("/admin/dashboards/"+dashboardAdminPageContext.databaseID,"Dashboards")
	
	
})