
var formRolePrivsPrefix = "formPrivSel_"

function formRoleButtonsRadioName(roleID) {
	return formRolePrivsPrefix + roleID
}

function formRolePrivsButtonsHTML(roleID) {
	
	var radioName = formRoleButtonsRadioName(roleID)

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

function initFormRolePrivsButtons(roleID,privs,privsChangedFunc) {

	var radioName = formRoleButtonsRadioName(roleID)

	// Initialize the radio selection - Using the click() function is Bootstrap specific
	$(':radio[name="'+radioName+'"][value="' + privs + '"]').click()
	
	var radioSelector = 'input[type="radio"][name="'+radioName+'"]'
	
	$(radioSelector).change(function() {
		var newPrivs = this.value
		console.log("Privilege selection changed: radio name = " + radioName + " privilages = " + newPrivs)
		privsChangedFunc(roleID,newPrivs)
	});
	
}


function getFormRolePrivRadioButtonVals() {

	return getGroupedRadioButtonVals(formRolePrivsPrefix)
}

