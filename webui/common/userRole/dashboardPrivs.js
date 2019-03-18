// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

var dashboardRolePrivSelectionRadioPrefix = "dashboardPrivSelection_"

function dashboardRoleButtonsRadioName(roleID) { 
	return dashboardRolePrivSelectionRadioPrefix + roleID
}

function dashboardRolePrivsButtonsHTML(roleID) {
	
	var radioName = dashboardRoleButtonsRadioName(roleID)

	return '' + 
			'<div class="btn-group" data-toggle="buttons">' +
				  '<label class="btn btn-default active btn-sm">' +
				    	'<input type="radio" name="'+ radioName + '" value="none" autocomplete="off" checked>None' +
				  '</label>' +
				  '<label class="btn btn-default btn-sm">' +
				    	'<input type="radio" name="'+ radioName + '"  value = "view" autocomplete="off">View' +
				  '</label>' +
		'</div>';
}

function initDashboardRolePrivsButtons(roleID,privs,privsChangedFunc) {

	var radioName = dashboardRoleButtonsRadioName(roleID)

	// Initialize the radio selection - Using the click() function is Bootstrap specific
	$(':radio[name="'+radioName+'"][value="' + privs + '"]').click()
	
	var radioSelector = 'input[type="radio"][name="'+radioName+'"]'
	
	$(radioSelector).change(function() {
		var newPrivs = this.value
		console.log("Privilege selection changed: radio name = " + radioName + " privilages = " + newPrivs)
		privsChangedFunc(roleID,newPrivs)
	});
	
}




function getDashboardRolePrivRadioButtonVals() {

	return getGroupedRadioButtonVals(dashboardRolePrivSelectionRadioPrefix)
}
