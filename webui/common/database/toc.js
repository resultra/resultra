

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

function addFormLinkToTOCList(tocConfig, linkInfo) {
	
	var formLinkListItemHTML = '<a href="/newItem/' + linkInfo.linkID 
			+ '" class="list-group-item">' + linkInfo.name + '</a>'
	$('#tocFormList').append(formLinkListItemHTML)
}


function initDatabaseTOC(tocConfig) {
	
	
	$('#tocFormList').empty()
	var linkParams = { parentDatabaseID: tocConfig.databaseID }
	jsonAPIRequest("formLink/getList",linkParams,function(linkList) {
		for(var linkIndex = 0; linkIndex < linkList.length; linkIndex++) {
			var currLink = linkList[linkIndex]
			if(currLink.includeInSidebar) {
				addFormLinkToTOCList(tocConfig,currLink)		
			}
		}
	})
	
	
	var getDBInfoParams = { databaseID: tocConfig.databaseID }
	jsonAPIRequest("database/getInfo",getDBInfoParams,function(dbInfo) {
		console.log("Got database info: " + JSON.stringify(dbInfo))		
		
		$('#tocListList').empty()
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