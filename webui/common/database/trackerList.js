
function addTrackerListItem(trackerInfo) {

	var $trackerList = $("#myTrackerList")

	var $listItem = $('#trackerListItemTemplate').clone()
	$listItem.attr("id","")

	var $nameLabel = $listItem.find(".nameLabel")
	$nameLabel.text(trackerInfo.databaseName)
	
	var $settingsLink = $listItem.find(".adminEditPropsButton")
	var editPropsLink = '/admin/' + trackerInfo.databaseID
	$settingsLink.attr('href',editPropsLink)
	

	$trackerList.append($listItem)
	
}


function initTrackerList() {
	
	var getDBListParams = {} // no parameters necessary - gets the tracker list for the currently signed in user
	
	jsonAPIRequest("database/getList",getDBListParams,function(trackerList) {
		
		for (var trackerIndex=0; trackerIndex<trackerList.length; trackerIndex++) {	
			var trackerInfo = trackerList[trackerIndex]
			addTrackerListItem(trackerInfo)
		}
			
	})
	
	initButtonClickHandler('#newTrackerButton',function() {
		console.log("New form button clicked")
		openNewTrackerDialog()
	})
	
}