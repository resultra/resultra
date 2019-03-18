// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.



function initValueListSelectionPropertyPanel(panelParams) {
	var $valListSelection = $('#' + panelParams.elemPrefix + 'ValueListPropertyValueListSelection')
	
	var getValueListsParams = {
		parentDatabaseID: panelParams.databaseID
	}
	
	function initSelectionChangeHandler() {
		initSelectControlChangeHandler($valListSelection,function(selectedValueList) {
			if(selectedValueList.length == 0) {
				panelParams.saveValueListCallback(null)
			} else {
				console.log("Selected value list: " + selectedValueList)
				panelParams.saveValueListCallback(selectedValueList)
			}
		})			
	}
	
	jsonAPIRequest("valueList/getList", getValueListsParams, function(valueLists) {
		$valListSelection.empty()
		$valListSelection.append(selectOptionHTML("","Don't select from a value list"))
		for (var valListIndex = 0; valListIndex < valueLists.length; valListIndex++) {
			var currValList = valueLists[valListIndex]
			$valListSelection.append(selectOptionHTML(currValList.valueListID,currValList.name))
		}
		if(panelParams.defaultValueListID != null) {
			$valListSelection.val(panelParams.defaultValueListID)
		}
		initSelectionChangeHandler()
	})
	
}