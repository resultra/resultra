
var dashboardRolePrivSelectionRadioPrefix = "dashboardPrivSelection_"

function dashboardRolePrivsButtonsHTML(roleID) {
	
	var radioName = dashboardRolePrivSelectionRadioPrefix + roleID

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


function getDashboardRolePrivRadioButtonVals() {

	return getGroupedRadioButtonVals(dashboardRolePrivSelectionRadioPrefix)
}
