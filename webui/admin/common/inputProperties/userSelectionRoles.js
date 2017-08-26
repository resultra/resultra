function initUserSelectionRoleProps(databaseID, userSelection) {
	
	function roleListItem(roleInfo) {
		
		var checkboxHTML = '' +
			'<div class="list-group-item">' +
				'<div class="checkbox">' +
					'<input type="checkbox" id="' + roleInfo.roleID + '"></input>'+
					'<label for="' + roleInfo.roleID +  '"><span class="noselect roleNameLabel"></span></label>' +
				'</div>' + 
			'</div>'
		
		var $checkbox = $(checkboxHTML)
		$checkbox.find('.roleNameLabel').text(roleInfo.roleName)
		
		var $checkboxInput = $checkbox.find("input")
		
		var roleEnabled = false
		initCheckboxControlChangeHandler($checkboxInput,roleEnabled,function(isEnabled) {
			
/*			var setSelectableRoleParams = {
				roleID: rolePriv.roleID,
				alertID: alertInfo.alertID,
				alertEnabled: alertEnabled
			}			
			jsonAPIRequest("userRole/setAlertRolePrivs",setSelectableRoleParams,function(setPrivsStatus) {
			})
*/
		})
		
		return $checkbox
		
	}
	
	var dbRolesParams = { databaseID: databaseID }
	jsonAPIRequest("userRole/getDatabaseRoles",dbRolesParams,function(rolesInfo) {
		
		var $roleList = $('#adminUserSelectionRoleList')
		$roleList.empty()

		console.log("Got roles info: " + JSON.stringify(rolesInfo))
		$('#adminNewUserRoleList').empty()
		for(var roleInfoIndex = 0; roleInfoIndex<rolesInfo.length; roleInfoIndex++) {
			var currRoleInfo = rolesInfo[roleInfoIndex]
			$roleList.append(roleListItem(currRoleInfo))
		}
	})
	
}