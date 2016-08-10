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

