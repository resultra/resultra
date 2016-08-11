function formRolePrivsRoleNameHTML(roleName) {
	return '' + 
		'<div class="row">' +
			'<label>' + roleName + '</label>' +
		'</div>';
}

function formPropsPrivsButtonRowHTML(roleID) {
	return '' + 
		'<div class="row formRolePrivsPrivRadioRow">' +
			formRolePrivsButtonsHTML(roleID) + 
		'</div>';
}


function formRolePrivilegeListItemHTML(roleID) {
		
	return '' +
		'<div class="list-group-item formRolePrivListItem" id="'+roleID+'">' +
			'<div class="container-fluid">' +
				formRolePrivsRoleNameHTML(roleID) +
				formPropsPrivsButtonRowHTML(roleID)
			'</div>' +
		'</div>';
}

function initDesignFormRolePrivProperties() {
	$('#formRolesPrivilegesList').empty()
	$('#formRolesPrivilegesList').append(formRolePrivilegeListItemHTML("role1"))
	$('#formRolesPrivilegesList').append(formRolePrivilegeListItemHTML("role2"))
	$('#formRolesPrivilegesList').append(formRolePrivilegeListItemHTML("role3"))
}
