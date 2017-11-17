



function initTrackerList() {
	
	var $trackerList = $("#myTrackerList")

	function addTrackerListItem(trackerInfo) {


		var $listItem = $('#trackerListItemTemplate').clone()
		$listItem.attr("id","")

		var $nameLabel = $listItem.find(".nameLabel")
		$nameLabel.text(trackerInfo.databaseName)
		var openTrackerLink = '/main/' + trackerInfo.databaseID
	
		// Only enable the link to open the tracker if the tracker is  active.
		if(trackerInfo.isActive) {
			$nameLabel.attr('href',openTrackerLink)	
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