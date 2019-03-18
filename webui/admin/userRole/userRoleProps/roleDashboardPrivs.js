// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function roleDashboardPrivsDashboardNameHTML(dashboardName) {
	return '' + 
		'<div class="row">' +
			'<label>' + dashboardName + '</label>' +
		'</div>';
}

function roleDashboardPropsPrivsButtonRowHTML(dashboardID) {
	return '' + 
		'<div class="row dashboardRolePrivsPrivRadioRow">' +
			dashboardRolePrivsButtonsHTML(dashboardID) + 
		'</div>';
}


function roleDashboardPrivilegeListItemHTML(dashboardPriv) {
		
	return '' +
		'<div class="list-group-item formRolePrivListItem maxWidth300" id="'+dashboardPriv.DashboardID+'">' +
			'<div class="container-fluid">' +
				roleDashboardPrivsDashboardNameHTML(dashboardPriv.dashboardName) +
				roleDashboardPropsPrivsButtonRowHTML(dashboardPriv.dashboardID)
			'</div>' +
		'</div>';
}

function initRoleDashboardPrivProperties(roleID) {
	
	jsonAPIRequest("userRole/getRoleDashboardPrivs", { roleID: roleID }, function(roleDashboardPrivs) {
		
		$('#adminDashboardRolesPrivilegesList').empty()
		
		for(var privIndex=0; privIndex < roleDashboardPrivs.length; privIndex++) {
			var dashboardPriv = roleDashboardPrivs[privIndex]
			
			$('#adminDashboardRolesPrivilegesList').append(roleDashboardPrivilegeListItemHTML(dashboardPriv))
			
			initDashboardRolePrivsButtons(dashboardPriv.dashboardID,dashboardPriv.privs, function(dashboardID,privs) {
				
				var setDashboardRolePrivParams = {
					dashboardID: dashboardID,
					roleID: roleID,
					privs: privs
				}
				console.log("Updating dashboard privileges: " + JSON.stringify(setDashboardRolePrivParams))
				
				jsonAPIRequest("userRole/setDashboardRolePrivs", setDashboardRolePrivParams, function(dashboardPrivs) {
					console.log("Updating dashboard privileges: done")			
				})
			})	
			
		}
	})
}