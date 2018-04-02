function initDashboardIncludeInSidebarProperty(dashboardInfo) {
	
	initCheckboxChangeHandler('#adminDashboardIncludeInSidebar', 
				dashboardInfo.properties.includeInSidebar, function(newVal) {
		var setIncludeSidebarParams = {
			dashboardID: dashboardInfo.dashboardID,
			includeInSidebar: newVal
		}
		jsonAPIRequest("dashboard/setIncludeInSidebar",setIncludeSidebarParams,function(updatedLinkInfo) {
		})			
		
	})

}