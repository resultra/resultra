
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

function getFormRolePrivRadioButtonVal(roleID) {
	
	var radioVals = {}
	
	$("input:radio").each (function() {
		var radioName = $(this).attr('name')
		if(radioName.indexOf(formRolePrivsPrefix) == 0) {
			var idVal = radioName.replace(formRolePrivsPrefix,'')
			var radioSelector = 'input[name="'+radioName+'"]:checked'
			radioVals[idVal] = $(radioSelector).val()			
		}
	})
	
	console.log("Radio vals: " + JSON.stringify(radioVals))
	
	return radioVals
}

