// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initRoleCollaborators(databaseID,roleID) {
	
	function roleCollaboratorListItem(userRoleInfo) {
		var privCheckboxHTML = '' +
			'<div class="list-group-item">' +
				'<div class="checkbox">' +
					'<input type="checkbox" id="' + userRoleInfo.collaboratorID + '"></input>'+
					'<label for="' + userRoleInfo.collaboratorID +  '"><span class="noselect nameLabel"></span></label>' +
				'</div>' + 
			'</div>'
		
		
		var userInfo = userRoleInfo.userInfo
		var userNameDisplay = '@' + userInfo.userName + 
			" (" + userInfo.firstName + " " + userInfo.lastName + ")"
	
		
		var $privCheckbox = $(privCheckboxHTML)
		$privCheckbox.find('.nameLabel').text(userNameDisplay)
		
		var $checkboxInput = $privCheckbox.find("input")
		
		initCheckboxControlChangeHandler($checkboxInput,userRoleInfo.isMemberOfRole,function(alertEnabled) {	
			var roleParams = {
				userID: userInfo.userID,
				databaseID: databaseID,
				collaboratorID: userRoleInfo.collaboratorID,
				roleID: roleID,
				memberOfRole: $checkboxInput.prop("checked")
			}			
			jsonAPIRequest("admin/setUserRoleInfo",roleParams,function(userRoles) { })
		})
		
		return $privCheckbox
		
	}
	
	
	var getRoleInfoParams = { 
		databaseID: databaseID,
		roleID: roleID }
	jsonAPIRequest("admin/getRoleCollaborators",getRoleInfoParams,function(userRoleInfo) {
				
		var $privList = $('#adminRoleCollaboratorList')
		$privList.empty()
		
		for (var userRoleIndex = 0; userRoleIndex < userRoleInfo.length; userRoleIndex++) {
			var currUserRole = userRoleInfo[userRoleIndex]
			$privList.append(roleCollaboratorListItem(currUserRole))
		}
	})
		

}