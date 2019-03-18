// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initRoleNewItemPrivs(roleID) {
	
	function roleNewItemPrivilegeListItem(newItemPriv) {
		var privCheckboxHTML = '' +
			'<div class="list-group-item">' +
				'<div class="checkbox">' +
					'<input type="checkbox" id="' + newItemPriv.linkID + '"></input>'+
					'<label for="' + newItemPriv.linkID +  '"><span class="noselect linkNameLabel"></span></label>' +
				'</div>' + 
			'</div>'
		
		var $privCheckbox = $(privCheckboxHTML)
		$privCheckbox.find('.linkNameLabel').text(newItemPriv.linkName)
		
		var $checkboxInput = $privCheckbox.find("input")
		
		initCheckboxControlChangeHandler($checkboxInput,newItemPriv.linkEnabled,function(linkEnabled) {
			
			var params = {
				roleID: roleID,
				linkID: newItemPriv.linkID,
				linkEnabled: linkEnabled
			}			
			jsonAPIRequest("userRole/setNewItemRolePrivs",params,function(setPrivsStatus) {
			})
		})
		
		return $privCheckbox
		
	}
	
	jsonAPIRequest("userRole/getRoleNewItemPrivs", { roleID: roleID }, function(roleNewItemPrivs) {
		
		var $privList = $('#adminNewItemLinkRolesPrivilegesList')
		
		$privList.empty()
		
		for(var privIndex=0; privIndex < roleNewItemPrivs.length; privIndex++) {
			var newItemPriv = roleNewItemPrivs[privIndex]
			$privList.append(roleNewItemPrivilegeListItem(newItemPriv))
		}
	})
	
}