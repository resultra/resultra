function initDesignDashboardProperties(dashboardID) {
	jsonAPIRequest("dashboard/getProperties",{dashboardID:dashboardID},function(dashboardInfo) {
		initDashboardPropertiesDashboardName(dashboardInfo)
		initDesignDashboardRolePrivProperties(dashboardID)
		
		initDashboardIncludeInSidebarProperty(dashboardInfo)
		
	})
}