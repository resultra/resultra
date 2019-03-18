// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function addRoleFormPrivTableRow(formInfo) {
	
	var buttonsHTML = userRoleItemButtonsHTML()

	var privs = "Full Access"
	
	var rowHTML = '' +
		'<tr>' +
	         '<td style="vertical-align:middle;text-align:right;">' + formInfo.name +  '</td>' +
	         '<td>' + formRolePrivsButtonsHTML(formInfo.formID) +  '</td>' +
	     '</tr>'
	
	$('#roleFormPrivsTable').append(rowHTML)
}


function initRoleFormPrivSettingsTable(formsInfo) {
	
	$('#roleFormPrivsTable').empty()
	for(var formIndex = 0; formIndex < formsInfo.length; formIndex++) {
		var formInfo = formsInfo[formIndex]
		addRoleFormPrivTableRow(formInfo)
		
	}	
}