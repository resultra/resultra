// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function addRoleDashboardPrivTableRow(dashboardInfo) {
		
	var rowHTML = '' +
		'<tr>' +
	         '<td style="vertical-align:middle;text-align:right;">' + dashboardInfo.name +  '</td>' +
	         '<td>' + dashboardRolePrivsButtonsHTML(dashboardInfo.dashboardID) +  '</td>' +
	     '</tr>'
	
	$('#roleDashboardPrivsTable').append(rowHTML)
}


function initRoleDashboardPrivSettingsTable(dashboardsInfo) {
	
	
	$('#roleDashboardPrivsTable').empty()
	
	for(var dashboardIndex = 0; dashboardIndex < dashboardsInfo.length; dashboardIndex++) {
		var dashboardInfo = dashboardsInfo[dashboardIndex]
		addRoleDashboardPrivTableRow(dashboardInfo)
	}
	
}

