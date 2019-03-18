// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function roleListRolePrivsListNameHTML(listName) {
	return '' + 
		'<div class="row">' +
			'<label>' + listName + '</label>' +
		'</div>';
}

function roleListPropsPrivsButtonRowHTML(roleID) {
	return '' + 
		'<div class="row formRolePrivsPrivRadioRow">' +
			formRolePrivsButtonsHTML(roleID) + 
		'</div>';
}


function roleListPrivilegeListItemHTML(roleID,listPriv) {
		
	return '' +
		'<div class="list-group-item formRolePrivListItem maxWidth300" id="'+listPriv.roleID+'">' +
			'<div class="container-fluid">' +
				roleListRolePrivsListNameHTML(listPriv.listName) +
				formRolePrivsButtonsHTML(listPriv.listID)
			'</div>' +
		'</div>';
}

function initRoleListPrivProperties(roleID) {
	
	jsonAPIRequest("userRole/getRoleListPrivs", { roleID: roleID }, function(roleListPrivs) {
		
		var $privList = $('#adminRoleListPrivilegesList')
		
		$privList.empty()
		
		for(var privIndex=0; privIndex < roleListPrivs.length; privIndex++) {
			var listPriv = roleListPrivs[privIndex]
			
			$privList.append(roleListPrivilegeListItemHTML(roleID,listPriv))
			
			
			initFormRolePrivsButtons(listPriv.listID,listPriv.privs, function(listID,privs) {
				var setListRolePrivParams = {
					listID: listID,
					roleID: roleID,
					privs: privs }
				jsonAPIRequest("userRole/setListRolePrivs", setListRolePrivParams, function(listPrivs) {
					console.log("Updating list privileges: done")			
				})
			})	

		}
	})
}