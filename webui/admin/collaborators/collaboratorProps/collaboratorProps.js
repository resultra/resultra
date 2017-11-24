$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#collabPropsPage'))	
	initAdminPageHeader()
	initAdminSettingsTOC(collabPropsContext.databaseID,"settingsTOCUsers")
	
	
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
		
		var $checkboxInput = $roleCheckbox.find("input")
		initCheckboxControlChangeHandler($checkboxInput,isMemberOfRole,function(newVal) {
			
			var roleParams = {
				userID: collabPropsContext.userID,
				databaseID: collabPropsContext.databaseID,
				collaboratorID: collabPropsContext.collaboratorID,
				roleID: roleInfo.roleID,
				memberOfRole: $checkboxInput.prop("checked")
			}			
			jsonAPIRequest("admin/setUserRoleInfo",roleParams,function(userRoles) {
			})
		})
		
	
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
			collaboratorID: collabPropsContext.collaboratorID,
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
		
		var userInfo = userRoleInfo.userInfo
		var userNameDisplay = '@' + userInfo.userName + 
			" (" + userInfo.firstName + " " + userInfo.lastName + ")"
		$("#collabNameTableCell").text(userNameDisplay)
		
	})	
	
	
	
				
})