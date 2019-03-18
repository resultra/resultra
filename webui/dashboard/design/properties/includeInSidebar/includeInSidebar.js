// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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