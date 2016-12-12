$(document).ready(function() {	
	
	initUserDropdownMenu()
		
	var getDBListParams = {}
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
	
}); // document ready
