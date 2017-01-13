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


function formRolePrivilegeListItemHTML(formPriv) {
		
	return '' +
		'<div class="list-group-item formRolePrivListItem" id="'+formPriv.roleID+'">' +
			'<div class="container-fluid">' +
				formRolePrivsRoleNameHTML(formPriv.roleName) +
				formPropsPrivsButtonRowHTML(formPriv.roleID)
			'</div>' +
		'</div>';
}

function initDesignFormRolePrivProperties(formID) {
	
	jsonAPIRequest("userRole/getFormRolePrivs", { formID: formID }, function(formPrivs) {
		$('#formRolesPrivilegesList').empty()
		
		for(var formPrivIndex=0; formPrivIndex < formPrivs.length; formPrivIndex++) {
			var formPriv = formPrivs[formPrivIndex]
			$('#formRolesPrivilegesList').append(formRolePrivilegeListItemHTML(formPriv))
			initFormRolePrivsButtons(formPriv.roleID,formPriv.privs, function(roleID,privs) {
				var setFormRolePrivParams = {
					formID: formID,
					roleID: roleID,
					privs: privs
				}
				console.log("Updating form privileges: " + JSON.stringify(setFormRolePrivParams))
				jsonAPIRequest("userRole/setFormRolePrivs", setFormRolePrivParams, function(formPrivs) {
					console.log("Updating form privileges: done")			
				})
			})	
		}
	})
}
