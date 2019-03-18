// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function populateGlobalSelectionMenu(menuSelector, databaseID, globalTypes) {
	
	function populateSelectionMenuTableList(globalsInfo, menuSelector) {
		$(menuSelector).empty()		
		$(menuSelector).append(defaultSelectOptionPromptHTML("Select a global value"))
		$.each(globalsInfo, function(index, globalInfo) {
			$(menuSelector).append(selectFieldHTML(globalInfo.globalID, globalInfo.name))		
		})
	}
	
	var includeType = {}
	for (var globalTypeIndex = 0; globalTypeIndex < globalTypes.length; globalTypeIndex++) {
		var type = globalTypes[globalTypeIndex]
		includeType[type] = true
	}
	
	function filterByType(globalInfo) {
		if((includeType['all'] == true) || (includeType[globalInfo.type] == true)) {
			return true
		} else {
			return false
		}
	}
	
	
	
	var listParams =  { parentDatabaseID: databaseID }
	jsonAPIRequest("global/getList",listParams,function(globalsInfo) {
		
		var filteredInfo = globalsInfo.filter(filterByType)
		
		populateSelectionMenuTableList(filteredInfo,menuSelector)
	})
	
}