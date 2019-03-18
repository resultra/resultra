// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function dashboardRolePrivsRoleNameHTML(roleName) {
	return '' + 
		'<div class="row">' +
			'<label>' + roleName + '</label>' +
		'</div>';
}

function dashboardPropsPrivsButtonRowHTML(roleID) {
	return '' + 
		'<div class="row dashboardRolePrivsPrivRadioRow">' +
			dashboardRolePrivsButtonsHTML(roleID) + 
		'</div>';
}


function dashboardRolePrivilegeListItemHTML(dashboardPriv) {
		
	return '' +
		'<div class="list-group-item formRolePrivListItem" id="'+dashboardPriv.roleID+'">' +
			'<div class="container-fluid">' +
				dashboardRolePrivsRoleNameHTML(dashboardPriv.roleName) +
				dashboardPropsPrivsButtonRowHTML(dashboardPriv.roleID)
			'</div>' +
		'</div>';
}

function initDesignDashboardRolePrivProperties(dashboardID) {
	
	jsonAPIRequest("userRole/getDashboardRolePrivs", { dashboardID: dashboardID }, function(dashboardPrivs) {
		$('#dashboardRolesPrivilegesList').empty()
		
		for(var dashboardPrivIndex=0; dashboardPrivIndex < dashboardPrivs.length; dashboardPrivIndex++) {
			var dashboardPriv = dashboardPrivs[dashboardPrivIndex]
			$('#dashboardRolesPrivilegesList').append(dashboardRolePrivilegeListItemHTML(dashboardPriv))
			
			initDashboardRolePrivsButtons(dashboardPriv.roleID,dashboardPriv.privs, function(roleID,privs) {
				var setDashboardRolePrivParams = {
					dashboardID: dashboardID,
					roleID: roleID,
					privs: privs
				}
				console.log("Updating dashboard privileges: " + JSON.stringify(setDashboardRolePrivParams))
				
				jsonAPIRequest("userRole/setDashboardRolePrivs", setDashboardRolePrivParams, function(dashboardPrivs) {
					console.log("Updating dashwobard privileges: done")			
				})
			})	
			
		}
	})
}