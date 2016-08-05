function addRoleDashboardPrivTableRow(dashboardID) {
		
	var rowHTML = '' +
		'<tr>' +
	         '<td style="vertical-align:middle;text-align:right;">' + dashboardID +  '</td>' +
	         '<td>' + dashboardRolePrivsButtonsHTML(dashboardID) +  '</td>' +
	     '</tr>'
	
	$('#roleDashboardPrivsTable').append(rowHTML)
}


function initRoleDashboardPrivSettingsTable() {
	$('#roleDashboardPrivsTable').empty()
	addRoleDashboardPrivTableRow("dashboard1")
	addRoleDashboardPrivTableRow("dashboard2")
	addRoleDashboardPrivTableRow("dashboard3")
	
}