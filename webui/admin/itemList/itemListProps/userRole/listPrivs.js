// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function listRolePrivsRoleNameHTML(roleName) {
	return '' + 
		'<div class="row">' +
			'<label>' + roleName + '</label>' +
		'</div>';
}

function listPropsPrivsButtonRowHTML(roleID) {
	return '' + 
		'<div class="row formRolePrivsPrivRadioRow">' +
			formRolePrivsButtonsHTML(roleID) + 
		'</div>';
}


function listRolePrivilegeListItemHTML(listPriv) {
		
	return '' +
		'<div class="list-group-item formRolePrivListItem" id="'+listPriv.roleID+'">' +
			'<div class="container-fluid">' +
				listRolePrivsRoleNameHTML(listPriv.roleName) +
				formRolePrivsButtonsHTML(listPriv.roleID)
			'</div>' +
		'</div>';
}

function initListRolePrivProperties(listID) {
	
	jsonAPIRequest("userRole/getListRolePrivs", { listID: listID }, function(listPrivs) {
		
		var $privList = $('#listRolesPrivilegesList')
		
		$privList.empty()
		
		for(var listPrivIndex=0; listPrivIndex < listPrivs.length; listPrivIndex++) {
			var listPriv = listPrivs[listPrivIndex]
			$privList.append(listRolePrivilegeListItemHTML(listPriv))
			
			initFormRolePrivsButtons(listPriv.roleID,listPriv.privs, function(roleID,privs) {
				var setListRolePrivParams = {
					listID: listID,
					roleID: roleID,
					privs: privs
				}
				jsonAPIRequest("userRole/setListRolePrivs", setListRolePrivParams, function(listPrivs) {
					console.log("Updating list privileges: done")			
				})
			})	
		}
	})
}
