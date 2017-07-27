$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#collabPropsPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(collabPropsContext.databaseID)
	
	
	function addRoleToRoleCheckboxList(roleInfo) {
			
		var roleCheckboxHTML = '' +
			'<div class="list-group-item">' +
				'<div class="checkbox">' +
					'<input type="checkbox" id="' + roleInfo.roleID + '"></input>'+
					'<label for="' + roleInfo.roleID +  '"><span class="noselect roleNameLabel"></span></label>' +
				'</div>' + 
			'</div>'
		
		var $roleCheckbox = $(roleCheckboxHTML)
		$roleCheckbox.find('.roleNameLabel').text(roleInfo.roleName)
		
	
		$('#adminCollabRolesList').append($roleCheckbox)	
		
		$roleCheckbox.data("roleInfo",roleInfo)
	
	}
	
	
	
	var dbRolesParams = {
		databaseID: collabPropsContext.databaseID
	}
	jsonAPIRequest("userRole/getDatabaseRoles",dbRolesParams,function(rolesInfo) {
		
		console.log("Got roles info: " + JSON.stringify(rolesInfo))
		$('#adminNewUserRoleList').empty()
		for(var roleInfoIndex = 0; roleInfoIndex<rolesInfo.length; roleInfoIndex++) {
			var currRoleInfo = rolesInfo[roleInfoIndex]
			addRoleToRoleCheckboxList(currRoleInfo)
		}
		
	})
	
				
})