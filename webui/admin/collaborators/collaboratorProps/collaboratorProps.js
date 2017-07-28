$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#collabPropsPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(collabPropsContext.databaseID)
	
	
	function addRoleToRoleCheckboxList(roleInfo, isMemberOfRole) {
			
		var roleCheckboxHTML = '' +
			'<div class="list-group-item">' +
				'<div class="checkbox">' +
					'<input type="checkbox" id="' + roleInfo.roleID + '"></input>'+
					'<label for="' + roleInfo.roleID +  '"><span class="noselect roleNameLabel"></span></label>' +
				'</div>' + 
			'</div>'
		
		var $roleCheckbox = $(roleCheckboxHTML)
		$roleCheckbox.find('.roleNameLabel').text(roleInfo.roleName)
		if(isMemberOfRole) {
			$roleCheckbox.find("input").prop("checked",true)
		}
		
	
		$('#adminCollabRolesList').append($roleCheckbox)	
		
		$roleCheckbox.data("roleInfo",roleInfo)
	
	}
	
	function getRoleAndUserRoleInfo(roleInfoCallback) {
		var getsRemaining = 2
		
		var allRolesInfo
		var userRoleInfo 
		function processOneGet() {
			getsRemaining--
			if (getsRemaining <= 0) {
				roleInfoCallback(allRolesInfo,userRoleInfo)
			}
		}
		
		var dbRolesParams = {
			databaseID: collabPropsContext.databaseID
		}
		jsonAPIRequest("userRole/getDatabaseRoles",dbRolesParams,function(rolesInfo) {
			allRolesInfo = rolesInfo
			processOneGet()
		})
		
		var userRolesParams = {
			userID: collabPropsContext.userID,
			databaseID: collabPropsContext.databaseID
		}
		jsonAPIRequest("admin/getSingleUserRoleInfo",userRolesParams,function(userRoles) {
			userRoleInfo = userRoles
			processOneGet()
		})
	}
	
	getRoleAndUserRoleInfo(function (rolesInfo,userRoleInfo) {
		console.log("Got roles info: " + JSON.stringify(rolesInfo))
		
		var memberRoles = []
		for(var roleIndex = 0; roleIndex < userRoleInfo.roleInfo.length; roleIndex++) {
			var roleID = userRoleInfo.roleInfo[roleIndex].roleID
			memberRoles.push(roleID)
		}
		var roleMemberLookup = new IDLookupTable(memberRoles)
		
		$('#adminNewUserRoleList').empty()
		for(var roleInfoIndex = 0; roleInfoIndex<rolesInfo.length; roleInfoIndex++) {
			var currRoleInfo = rolesInfo[roleInfoIndex]
			addRoleToRoleCheckboxList(currRoleInfo,roleMemberLookup.hasID(currRoleInfo.roleID))
		}
		
	})	
				
})