
var formRolePrivsPrefix = "formPrivSel_"

function formRolePrivsButtonsHTML(roleID) {
	
	var radioName = formRolePrivsPrefix + roleID

	return '' + 
			'<div class="btn-group" data-toggle="buttons">' +
				  '<label class="btn btn-default active btn-sm">' +
				    	'<input type="radio" name="'+ radioName + '" value="none" autocomplete="off" checked>None' +
				  '</label>' +
				  '<label class="btn btn-default btn-sm">' +
				    	'<input type="radio" name="'+ radioName + '"  value = "view" autocomplete="off">View' +
				  '</label>' +
				  '<label class="btn btn-default btn-sm">' +
				    	'<input type="radio" name="'+ radioName + '"  value = "edit" autocomplete="off">Edit' +
				  '</label>' +
		'</div>';
}


function getFormRolePrivRadioButtonVals() {

	return getGroupedRadioButtonVals(formRolePrivsPrefix)
}

