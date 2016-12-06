
function addFormToFormTOCList(formInfo) {
	var formListItemHTML = '<a href="/viewForm/' + formInfo.formID + '" class="list-group-item">' + formInfo.name + '</a>'
	$('#tocListList').append(formListItemHTML)		
}

function addDashboardLinkToTOCList(dashboardInfo) {
	// TODO - Link to "dashboard view" page instead of dashboard design page (view page isn't implemented yet)
	var dashboardListItemHTML = '<a href="/viewDashboard/' + dashboardInfo.dashboardID 
			+ '" class="list-group-item">' + dashboardInfo.name + '</a>'
	$('#tocDashboardList').append(dashboardListItemHTML)		
}

function addItemListLinkToTOCList(listInfo) {
	var itemListItemHTML = '<a href="/viewList/' + listInfo.listID 
			+ '" class="list-group-item">' + listInfo.name + '</a>'
	$('#tocListList').append(itemListItemHTML)		
	
}


function initDatabaseTOC(databaseID) {
	
	var getDBInfoParams = { databaseID: databaseID }
	jsonAPIRequest("database/getInfo",getDBInfoParams,function(dbInfo) {
		console.log("Got database info: " + JSON.stringify(dbInfo))
		
		$('#tocListList').empty()
		for (var formInfoIndex = 0; formInfoIndex < dbInfo.formsInfo.length; formInfoIndex++) {
			var formInfo = dbInfo.formsInfo[formInfoIndex]
			addFormToFormTOCList(formInfo)
		}
		for(var listInfoIndex = 0; listInfoIndex < dbInfo.listsInfo.length; listInfoIndex++) {
			var listInfo = dbInfo.listsInfo[listInfoIndex]
			addItemListLinkToTOCList(listInfo)
		}

		$('#tocDashboardList').empty()
		for (var dashboardInfoIndex = 0; dashboardInfoIndex < dbInfo.dashboardsInfo.length; dashboardInfoIndex++) {
			var dashboardInfo = dbInfo.dashboardsInfo[dashboardInfoIndex]
			addDashboardLinkToTOCList(dashboardInfo)
		}

		
	}) // getRecord
	
}