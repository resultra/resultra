


function addTrackerListItem(trackerInfo) {

	var $trackerList = $("#myTrackerList")

	var $listItem = $('#trackerListItemTemplate').clone()
	$listItem.attr("id","")

	var $nameLabel = $listItem.find(".trackerLinkNameLabel")
	$nameLabel.text(trackerInfo.databaseName)
	
	
	// Only enable the link to open the tracker if the tracker is  active.
	if(trackerInfo.isActive) {
		
		$nameLabel.click(function() {
		 	   console.log("tracker link clicked")
			navigateToTracker(trackerInfo)
		})
		
	} else {
		$nameLabel.addClass("disabledTrackerLink")
	}

	var $settingsLink = $listItem.find(".adminEditPropsButton")

	if (trackerInfo.isAdmin) {
		var editPropsLink = '/admin/' + trackerInfo.databaseID
		$settingsLink.attr('href',editPropsLink)
		$settingsLink.tooltip()
	} else {
		$settingsLink.hide()
	}

	$trackerList.append($listItem)

}



function initTrackerList() {
	
	var $trackerList = $("#myTrackerList")

		
	function reloadTrackerList(includeInactive) {
		var getDBListParams = {
			includeInactive:includeInactive
		}
		jsonAPIRequest("database/getList",getDBListParams,function(trackerList) {
			$trackerList.empty()
			for (var trackerIndex=0; trackerIndex<trackerList.length; trackerIndex++) {	
				var trackerInfo = trackerList[trackerIndex]
				addTrackerListItem(trackerInfo)
			}
		})
	}
	reloadTrackerList(false)
	
	
	initButtonClickHandler('#newTrackerButton',function() {
		console.log("New form button clicked")
		openNewTrackerDialog()
	})
	
	initCheckboxChangeHandler('#showInactiveTrackers', false, function(includeInactive) {
		reloadTrackerList(includeInactive)
	})

	
}