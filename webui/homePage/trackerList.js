// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.



function addTrackerListItem(pageContext,trackerInfo) {

	var $trackerList = $("#myTrackerList")

	var $listItem = $('#trackerListItemTemplate').clone()
	$listItem.attr("id","")

	var $nameLabel = $listItem.find(".trackerLinkNameLabel")
	$nameLabel.text(trackerInfo.databaseName)
	
	
	// Only enable the link to open the tracker if the tracker is  active.
	if(trackerInfo.isActive) {
		
		$nameLabel.click(function() {
		 	   console.log("tracker link clicked")
			navigateToTracker(pageContext,trackerInfo)
		})
		
	} else {
		$nameLabel.addClass("disabledTrackerLink")
	}
	
	var $settingsButton = $listItem.find(".editTrackerPropsButton")
	if(trackerInfo.isAdmin) {
		$settingsButton.click(function(e) {
			e.preventDefault()
			$settingsButton.blur()
			navigateToAdminSettingsPageContent(pageContext,trackerInfo)
		})
	} else {
		$settingsButton.hide()
	}

	$trackerList.append($listItem)

}



function initTrackerList(pageContext) {
	
	var $trackerList = $("#myTrackerList")

		
	function reloadTrackerList(includeInactive) {
		var getDBListParams = {
			includeInactive:includeInactive
		}
		jsonAPIRequest("database/getList",getDBListParams,function(trackerList) {
			$trackerList.empty()
			for (var trackerIndex=0; trackerIndex<trackerList.length; trackerIndex++) {	
				var trackerInfo = trackerList[trackerIndex]
				addTrackerListItem(pageContext,trackerInfo)
			}
		})
	}
	reloadTrackerList(false)
	
	
	initButtonClickHandler('#newTrackerButton',function() {
		console.log("New form button clicked")
		openNewTrackerDialog(pageContext)
	})
	
	initCheckboxChangeHandler('#showInactiveTrackers', false, function(includeInactive) {
		reloadTrackerList(includeInactive)
	})
	
	var $viewTemplatesLink = $('#viewTemplatesButton')
	$viewTemplatesLink.click(function() {
		
		navigateToMainWindowContent("workspaceTemplates")
	})
	

	
}