
function addFormToFormTOCList(formInfo) {
	var formListItemHTML = '<a href="/viewForm/' + formInfo.formID + '" class="list-group-item">' + formInfo.name + '</a>'
	$('#tocFormList').append(formListItemHTML)		
}

function addDashboardLinkToTOCList(dashboardInfo) {
	// TODO - Link to "dashboard view" page instead of dashboard design page (view page isn't implemented yet)
	var dashboardListItemHTML = '<a href="/designDashboard/' + dashboardInfo.dashboardID 
			+ '" class="list-group-item">' + dashboardInfo.name + '</a>'
	$('#tocDashboardList').append(dashboardListItemHTML)		
}


function initDatabaseTOC(databaseID) {
	
	var getDBInfoParams = { databaseID: databaseID }
	jsonAPIRequest("databaseInfo/getInfo",getDBInfoParams,function(dbInfo) {
		console.log("Got database info: " + JSON.stringify(dbInfo))
		
		$('#tocFormList').empty()
		for (var formInfoIndex = 0; formInfoIndex < dbInfo.formsInfo.length; formInfoIndex++) {
			var formInfo = dbInfo.formsInfo[formInfoIndex]
			addFormToFormTOCList(formInfo)
		}

		$('#tocDashboardList').empty()
		for (var dashboardInfoIndex = 0; dashboardInfoIndex < dbInfo.dashboardsInfo.length; dashboardInfoIndex++) {
			var dashboardInfo = dbInfo.dashboardsInfo[dashboardInfoIndex]
			addDashboardLinkToTOCList(dashboardInfo)
		}

		
	}) // getRecord
	
}