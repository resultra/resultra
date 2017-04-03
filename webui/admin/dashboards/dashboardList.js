var adminDashboardListElemPrefix = "adminDashboardList_"


function dashboardDesignPageHTMLLink(dashboardID) {
	return '/admin/dashboard/' + dashboardID
}

function adminDashboardListButtonsHTML(dashboardInfo) {
return '' +
			'<div class="pull-right adminDashboardListButtons">' + 
	
			'<a class="btn btn-xs btn-default" href="' + dashboardDesignPageHTMLLink(dashboardInfo.dashboardID) + 
					'" role="button">' + 
				'<span class="glyphicon glyphicon-pencil" style="padding-bottom:2px;"></span>' +
			'</a>' + 
  			'<button class="btn btn-xs btn-danger deleteDashboardButton">' + 
				// padding-bottom: 2px makes the button image vertically line up better.
				'<span class="glyphicon glyphicon-remove" style="padding-bottom:2px;"></span>' +
			'</button>';

			'</div>'
}


function addDashboardToAdminDashboardList(dashboardInfo) {
	 
	var dashboardListDashboardID = adminDashboardListElemPrefix + dashboardInfo.dashboardID
	
	var dashboardListItemHTML = '<li class="list-group-item dashboardListItem" id="' + dashboardListDashboardID + '">' + 
		dashboardInfo.name +
		adminDashboardListButtonsHTML(dashboardInfo) +
	 '</li>'
	
	var $dashboardListItem = $(dashboardListItemHTML)
	$dashboardListItem.attr('data-dashboardID',dashboardInfo.dashboardID)
	
	$('#adminDashboardList').append($dashboardListItem)		
}

function initAdminDashboardSettings(databaseID) {
	
	var $dashboardList =  $("#adminDashboardList")
	
    $dashboardList.sortable({
		placeholder: "ui-state-highlight",
		cursor:"move",
		update: function( event, ui ) {
			
			var dashboardOrder = []
			$dashboardList.find(".dashboardListItem").each( function() {
				var dashboardID = $(this).attr('data-dashboardID')
				dashboardOrder.push(dashboardID)
			})
			var setOrderParams = {
				databaseID:databaseID,
				dashboardOrder: dashboardOrder
			}
			console.log("New dashboard sort order:" + JSON.stringify(dashboardOrder))
			jsonAPIRequest("database/setDashboardOrder",setOrderParams,function(dbInfo) {
				console.log("Done changing database dashboardOrder")
			})
			
		}
    });
	
	
	var getDBInfoParams = { databaseID: databaseID }
	jsonAPIRequest("database/getInfo",getDBInfoParams,function(dbInfo) {
		console.log("Got database info: " + JSON.stringify(dbInfo))
		
		$dashboardList.empty()
		for (var dashboardInfoIndex = 0; dashboardInfoIndex < dbInfo.dashboardsInfo.length; dashboardInfoIndex++) {
			var dashboardInfo = dbInfo.dashboardsInfo[dashboardInfoIndex]
			addDashboardToAdminDashboardList(dashboardInfo)
		}
		
	})
	
	
	initButtonClickHandler('#adminNewDashboardButton',function() {
		console.log("New form button clicked")
		openNewDashboardDialog(databaseID)
	})
	
	
	
}