// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
// Javascript for user roles


function addAdminRoleTableRow(adminUsers) {
	
	var adminUserDispl = []
	for(var adminUserIndex = 0; adminUserIndex < adminUsers.length; adminUserIndex++) {
		var userInfo = adminUsers[adminUserIndex]
		var userNameDisplay = '@' + userInfo.userName + 
			" (" + userInfo.firstName + " " + userInfo.lastName + ")"
		adminUserDispl.push(userNameDisplay)
	}

	// TBD - Put controls to configure the admin user in the roleListButtonCell

	var privs = "<strong>Full Access</strong>"
	
	var rowHTML = '' +
		'<tr class="userListRow">' +
	         '<td>' + "<strong>Administrator<strong>" +  '</td>' +
	         '<td>' + adminUserDispl.join(", ") +  '</td>' +
	         '<td class="roleListButtonCell"></td>' +
	     '</tr>'
	
	$('#userRoleTableBody').append(rowHTML)
}

function addCustomRoleTableRow(pageContext,customRoleInfo) {
	
	
	var roleUsersDisplay = []
	for(var userInfoIndex = 0; userInfoIndex < customRoleInfo.roleUsers.length; userInfoIndex++) {
		var userInfo = customRoleInfo.roleUsers[userInfoIndex]
		var userDisplay = '@' + userInfo.userName
		roleUsersDisplay.push(userDisplay)
	}
	
	
	var $roleRow = $("#userListRowTemplate").clone()
	$roleRow.attr("id","")
	
	var $nameCell = $roleRow.find(".userRoleName")
	$nameCell.text(customRoleInfo.roleName)
	
	var $collabsCell = $roleRow.find(".userRoleCollaborators")
	$collabsCell.text(roleUsersDisplay.join(", "))
	
	var editRoleContentURL = '/admin/userRole/' + customRoleInfo.roleID
	
	var $editRoleButton = $roleRow.find(".editUserRoleButton")
	setPageContentButtonClickHandler($editRoleButton,editRoleContentURL,function() {
		initUserRolePropsAdminSettingsPageContent(pageContext,customRoleInfo)
	})
	
		
	$('#userRoleTableBody').append($roleRow)
}



function initUserRoleSettings(pageContext) {
	
	var getRoleInfoParams = { databaseID: pageContext.databaseID }
	jsonAPIRequest("admin/getRoleInfo",getRoleInfoParams,function(roleInfo) {
		
		console.log("Got role info: " + JSON.stringify(roleInfo))
		console.log("Number of roles: " + roleInfo.length)
		
		addAdminRoleTableRow(roleInfo.adminUsers)
		
		for(var customRoleIndex = 0; customRoleIndex < roleInfo.customRoles.length; customRoleIndex++) {
			var customRoleInfo = roleInfo.customRoles[customRoleIndex]
			addCustomRoleTableRow(pageContext,customRoleInfo)
		}
		
	})
	
	initButtonClickHandler('#userRoleNewRoleButton',function() {
		console.log("New Role button clicked")
		openNewUserRoleDialog(pageContext)
	})
	
}