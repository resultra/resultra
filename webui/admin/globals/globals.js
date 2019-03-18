// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

function adminGlobalListButtonsHTML(formInfo) {
return '' +
			'<div class="pull-right adminGlobalListButtons">' + 
	
  			'<button class="btn btn-xs btn-default editGlobalButton">' + 
				// padding-bottom: 2px makes the button image vertically line up better.
				'<span class="glyphicon glyphicon-pencil" style="padding-bottom:2px;"></span>' +
			'</button>' + 
	
  			'<button class="btn btn-xs btn-danger deleteGlobalButton">' + 
				// padding-bottom: 2px makes the button image vertically line up better.
				'<span class="glyphicon glyphicon-remove" style="padding-bottom:2px;"></span>' +
			'</button>';

			'</div>'
}


function addGlobalToAdminList(globalInfo) {
	 	
	var listItemHTML = '<li class="list-group-item" id="' + globalInfo.formID + '">' + 
		globalInfo.name +
		adminGlobalListButtonsHTML(globalInfo) +
	 '</li>'
	
	$('#adminGlobalsList').append(listItemHTML)		
}

function initAdminGlobals(databaseID) {
	
	
	var getGlobalsParams = { parentDatabaseID: databaseID }
	jsonAPIRequest("global/getList",getGlobalsParams,function(globalsInfo) {
		$('#adminGlobalsList').empty()
	
		for(var globalIndex = 0; globalIndex < globalsInfo.length; globalIndex++) {
			var globalInfo = globalsInfo[globalIndex]
			addGlobalToAdminList(globalInfo)
		}
	})
	
	
	initButtonClickHandler('#adminGlobalsNewGlobalButton',function() {
		console.log("New Global button clicked")
		openNewGlobalDialog(databaseID)
	})
	
}