function formRolePrivsSelectionHTML(roleID) {
	
	var selectionID = "privSel_" + roleID
		
	return '' + 
		'<div class="row">' +
			'<select class="form-control input-sm" id="'+ selectionID + '">' + 
				selectOptionHTML("none", "None") +
				selectOptionHTML("view", "View") +
				selectOptionHTML("edit", "Edit") +
			'</select>' +
		'</div>';
}

function formRolePrivsRoleNameHTML(roleName) {
	return '' + 
		'<div class="row">' +
				'<div class="col-sm-12">' +
					'<label>' + roleName + '</label>' +
				'</div>' +
		'</div>';
}

function formRolePrivsButtonsHTML(roleID) {
	
		var radioName = "privSelection_" + roleID
	
		return '' + 
			'<div class="row formRolePrivsPrivRadioRow">' +
				'<div class="col-sm-12">' +
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
					'</div>' +
				'</div>' +
			'</div>';
}


function formRolePrivilegeListItemHTML(roleID) {
		
	return '' +
		'<div class="list-group-item formRolePrivListItem" id="'+roleID+'">' +
			'<div class="container-fluid">' +
				formRolePrivsRoleNameHTML(roleID) +
				formRolePrivsButtonsHTML(roleID)
			'</div>' +
		'</div>';
}


function initDesignFormProperties() {
	$('#formRolesPrivilegesList').empty()
	$('#formRolesPrivilegesList').append(formRolePrivilegeListItemHTML("role1"))
	$('#formRolesPrivilegesList').append(formRolePrivilegeListItemHTML("role2"))
	$('#formRolesPrivilegesList').append(formRolePrivilegeListItemHTML("role3"))
	
}