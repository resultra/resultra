// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initAdminValueListListSettings(databaseID) {
	
	var $valueListList = $('#adminValueListList')
	
	
	function addValueListToValueListList(valueListInfo) {
 
		var $listItem = $('#valueListItemTemplate').clone()
		$listItem.attr("id","")
	
		$listItem.attr("data-listID",valueListInfo.valueListID)
	
		var $editPropsButton = $listItem.find(".editValueListPropsButton")
		
		$editPropsButton.click(function(e) {
			e.preventDefault()
			$editPropsButton.blur()
			
			var editPropsContentURL = '/admin/valueList/' + valueListInfo.valueListID
			setSettingsPageContent(editPropsContentURL,function() {
				initValueListSettingsPageContent(valueListInfo)
			})
		})
				
		var $nameLabel = $listItem.find(".adminValueListNameLabel")
		$nameLabel.text(valueListInfo.name)
		 	
		$valueListList.append($listItem)		
	}
	
	var getValueListsParams = { parentDatabaseID: databaseID }
	jsonAPIRequest("valueList/getList",getValueListsParams,function(valueListList) {
		$valueListList.empty()
		for(var valueListIndex = 0; valueListIndex < valueListList.length; valueListIndex++) {
			var currValList = valueListList[valueListIndex]
			addValueListToValueListList(currValList)
		}
	})
	
	
	
	initButtonClickHandler('#adminNewValueListButton',function() {
		console.log("New value list button clicked")
		openNewValueListDialog(databaseID)
	})
	
	
}