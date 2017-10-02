function initUserSelectionRoleProps(params) {
	
	var $roleList = $('#'+ params.elemPrefix + 'AdminUserSelectionRoleList')
	var initiallySelectedRoles = new IDLookupTable(params.initialRoles)
	
	
	function roleListItem(roleInfo) {
		
		var checkboxID = "userRoleSel_" + roleInfo.roleID
		
		var checkboxHTML = '' +
			'<div class="list-group-item">' +
				'<div class="checkbox">' +
					'<input type="checkbox" id="' + checkboxID + '"></input>'+
					'<label for="' + checkboxID +  '"><span class="noselect roleNameLabel"></span></label>' +
				'</div>' + 
			'</div>'
		
		var $checkbox = $(checkboxHTML)
		$checkbox.find('.roleNameLabel').text(roleInfo.roleName)
		
		var $checkboxInput = $checkbox.find("input")
		$checkboxInput.attr('data-roleID',roleInfo.roleID)
		
		var roleEnabled = initiallySelectedRoles.hasID(roleInfo.roleID)
		initCheckboxControlChangeHandler($checkboxInput,roleEnabled,function(isEnabled) {
			
			var checkedRoles = []
			$roleList.find("input").each(function() {
				if($(this).prop("checked")===true) {
					var roleID = $(this).attr('data-roleID')
					checkedRoles.push(roleID)
				}
			})
			console.log("checked roles: " + JSON.stringify(checkedRoles))
			params.setRolesCallback(checkedRoles)
			
		})
		
		return $checkbox
		
	}
	
	var dbRolesParams = { databaseID: params.databaseID }
	jsonAPIRequest("userRole/getDatabaseRoles",dbRolesParams,function(rolesInfo) {
		
		$roleList.empty()
		

		console.log("Got roles info: " + JSON.stringify(rolesInfo))
		$roleList.empty()
		for(var roleInfoIndex = 0; roleInfoIndex<rolesInfo.length; roleInfoIndex++) {
			var currRoleInfo = rolesInfo[roleInfoIndex]
			$roleList.append(roleListItem(currRoleInfo))
		}
	})
	
}