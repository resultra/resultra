

function initTrackerList() {
	
	var getDBListParams = {} // no parameters necessary - gets the tracker list for the currently signed in user
	var $trackerList = $("#myTrackerList")
	
	jsonAPIRequest("database/getList",getDBListParams,function(trackerList) {
		
		for (var trackerIndex=0; trackerIndex<trackerList.length; trackerIndex++) {
			
			var trackerInfo = trackerList[trackerIndex]
			
			var $listItem = $('#trackerListItemTemplate').clone()
			$listItem.attr("id","")
		
			var $nameLabel = $listItem.find(".nameLabel")
			$nameLabel.text(trackerInfo.databaseName)
			
			var $settingsLink = $listItem.find(".adminEditPropsButton")
			var editPropsLink = '/admin/' + trackerInfo.databaseID
			$settingsLink.attr('href',editPropsLink)
			
		
			$trackerList.append($listItem)
		}
			
	})
	
}