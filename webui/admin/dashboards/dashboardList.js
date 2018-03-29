var adminDashboardListElemPrefix = "adminDashboardList_"



function addDashboardToAdminDashboardList(pageContext,dashboardInfo) {
	 
	var dashboardListDashboardID = adminDashboardListElemPrefix + dashboardInfo.dashboardID
		
	var $listItem = $('#adminDashboardListItemTemplate').clone()
	$listItem.attr("id",dashboardListDashboardID)
	$listItem.attr('data-dashboardID',dashboardInfo.dashboardID)
	
	$listItem.find(".adminDashboardListDashboardName").text(dashboardInfo.name)
	
	var $editButton = $listItem.find(".adminDashboardListEditDashboardButton")
	$editButton.click(function(e) {
		e.preventDefault()
		$editButton.blur()
		navigateToDashboardDesignerPageContent(pageContext,dashboardInfo)
	})
	// TODO - initialize button
		
	$('#adminDashboardList').append($listItem)		
}

function initAdminDashboardSettings(pageContext) {
	
	var $dashboardList =  $("#adminDashboardList")
	
    $dashboardList.sortable({
		placeholder: "ui-state-highlight",
		cursor:"move",
		update: function( event, ui ) {
			
			var dashboardOrder = []
			$dashboardList.find(".list-group-item").each( function() {
				var dashboardID = $(this).attr('data-dashboardID')
				dashboardOrder.push(dashboardID)
			})
			var setOrderParams = {
				databaseID:pageContext.databaseID,
				dashboardOrder: dashboardOrder
			}
			console.log("New dashboard sort order:" + JSON.stringify(dashboardOrder))
			jsonAPIRequest("database/setDashboardOrder",setOrderParams,function(dbInfo) {
				console.log("Done changing database dashboardOrder")
			})
			
		}
    });
	
	
	var getDBInfoParams = { databaseID: pageContext.databaseID }
	jsonAPIRequest("database/getInfo",getDBInfoParams,function(dbInfo) {
		console.log("Got database info: " + JSON.stringify(dbInfo))
		
		$dashboardList.empty()
		for (var dashboardInfoIndex = 0; dashboardInfoIndex < dbInfo.dashboardsInfo.length; dashboardInfoIndex++) {
			var dashboardInfo = dbInfo.dashboardsInfo[dashboardInfoIndex]
			addDashboardToAdminDashboardList(pageContext,dashboardInfo)
		}
		
	})
	
	
	initButtonClickHandler('#adminNewDashboardButton',function() {
		console.log("New form button clicked")
		openNewDashboardDialog(pageContext)
	})
	
	
	
}