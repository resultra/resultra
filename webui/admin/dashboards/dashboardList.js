var adminDashboardListElemPrefix = "adminDashboardList_"


function dashboardDesignPageHTMLLink(dashboardID) {
	return '/designDashboard/' + dashboardID
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
	
	var dashboardListItemHTML = '<li class="list-group-item" id="' + dashboardListDashboardID + '">' + 
		dashboardInfo.name +
		adminDashboardListButtonsHTML(dashboardInfo) +
	 '</li>'
	
	$('#adminDashboardList').append(dashboardListItemHTML)		
}

function initAdminDashboardSettings(databaseID) {
	
    $("#adminDashboardList").sortable({
		placeholder: "ui-state-highlight",
		cursor:"move",
		update: function( event, ui ) {
			// Get the new sorted list of form IDs. The prefix needs to be stripped from the ID.
			var prefixRegexp = new RegExp('^' + adminDashboardListElemPrefix)
			var sortedIDs =  $("#adminDashboardList").sortable("toArray").map(function(elem) {
				return elem.replace(prefixRegexp,'')
			})
			console.log("New sort order:" + JSON.stringify(sortedIDs))
		}
    });
	
	
	var getDBInfoParams = { databaseID: databaseID }
	jsonAPIRequest("database/getInfo",getDBInfoParams,function(dbInfo) {
		console.log("Got database info: " + JSON.stringify(dbInfo))
		
		$('#adminDashboardList').empty()
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