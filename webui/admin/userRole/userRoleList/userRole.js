// Javascript for user roles

function userRoleItemButtonsHTML(roleID) {
	
	var editRoleLink = '/admin/userRole/' + roleID
	
	
return '' +
			'<div class="pull-right userListItemButtons">' + 
	
			'<a class="btn btn-xs btn-default editUserRoleButton" role="button" href="'+ editRoleLink + '">' + 
					'<span class="glyphicon glyphicon-pencil" style="padding-bottom:2px;"></span>' +
				'</a>' + 
  			'<button class="btn btn-xs btn-danger deleteUserRoleButton">' + 
				// padding-bottom: 2px makes the button image vertically line up better.
				'<span class="glyphicon glyphicon-remove" style="padding-bottom:2px;"></span>' +
			'</button>';

			'</div>'

	
}

function userRoleTableRowHTML(roleID) {
	
	var roles = "TBD"
	var privs = "TBD"
	
	var buttonsHTML = userRoleItemButtonsHTML(roleID)
	
	return '' +
		'<tr class="userListRow">' +
	         '<td>' + roleID +  '</td>' +
	         '<td>' + roles +  '</td>' +
	         '<td>' + privs +  '</td>' +
	         '<td>' + privs +  '</td>' +
	         '<td class="userListButtonCell">' + buttonsHTML + '</td>' +
	     '</tr>'
	
}

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

function addCustomRoleTableRow(customRoleInfo) {
	
	
	var roleUsersDisplay = []
	for(var userInfoIndex = 0; userInfoIndex < customRoleInfo.roleUsers.length; userInfoIndex++) {
		var userInfo = customRoleInfo.roleUsers[userInfoIndex]
		var userDisplay = '@' + userInfo.userName
		roleUsersDisplay.push(userDisplay)
	}
	

	var buttonsHTML = userRoleItemButtonsHTML(customRoleInfo.roleID)

	var privs = "Full Access"
	
	var rowHTML = '' +
		'<tr class="userListRow">' +
	         '<td>' + customRoleInfo.roleName +  '</td>' +
	         '<td>' + roleUsersDisplay.join(", ") +  '</td>' +
	         '<td class="roleListButtonCell">' + buttonsHTML + '</td>' +
	     '</tr>'
	
	$('#userRoleTableBody').append(rowHTML)
}



function initUserRoleSettings(databaseID) {
	
	var getRoleInfoParams = { databaseID: databaseID }
	jsonAPIRequest("admin/getRoleInfo",getRoleInfoParams,function(roleInfo) {
		
		console.log("Got role info: " + JSON.stringify(roleInfo))
		console.log("Number of roles: " + roleInfo.length)
		
		addAdminRoleTableRow(roleInfo.adminUsers)
		
		for(var customRoleIndex = 0; customRoleIndex < roleInfo.customRoles.length; customRoleIndex++) {
			var customRoleInfo = roleInfo.customRoles[customRoleIndex]
			addCustomRoleTableRow(customRoleInfo)
		}
		
	})
	
	initButtonClickHandler('#userRoleNewRoleButton',function() {
		console.log("New Role button clicked")
		openNewUserRoleDialog(databaseID)
	})
	
}