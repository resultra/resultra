// Javascript for user roles

function userRoleItemButtonsHTML() {
return '' +
			'<div class="pull-right userListItemButtons">' + 
	
  			'<button class="btn btn-xs editUserRoleButton">' + 
				// padding-bottom: 2px makes the button image vertically line up better.
				'<span class="glyphicon glyphicon-pencil" style="padding-bottom:2px;"></span>' +
			'</button>' +
  			'<button class="btn btn-xs btn-danger deleteUserRoleButton">' + 
				// padding-bottom: 2px makes the button image vertically line up better.
				'<span class="glyphicon glyphicon-remove" style="padding-bottom:2px;"></span>' +
			'</button>';

			'</div>'

	
}

function userRoleTableRowHTML(roleID) {
	
	var roles = "TBD"
	var privs = "TBD"
	
	var buttonsHTML = userRoleItemButtonsHTML()
	
	return '' +
		'<tr class="userListRow">' +
	         '<td>' + roleID +  '</td>' +
	         '<td>' + roles +  '</td>' +
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

	var buttonsHTML = userRoleItemButtonsHTML()

	var privs = "Full Access"
	
	var rowHTML = '' +
		'<tr class="userListRow">' +
	         '<td>' + "<strong>Administrator<strong>" +  '</td>' +
	         '<td>' + adminUserDispl.join(", ") +  '</td>' +
	         '<td>' + privs +  '</td>' +
	         '<td class="userListButtonCell">' + buttonsHTML + '</td>' +
	     '</tr>'
	
	$('#userRoleTableBody').append(rowHTML)
}


function initUserRoleSettings(databaseID) {
	
	var getRoleInfoParams = { databaseID: databaseID }
	jsonAPIRequest("admin/getRoleInfo",getRoleInfoParams,function(roleInfo) {
		
		console.log("Got role info: " + JSON.stringify(roleInfo))
		console.log("Number of roles: " + roleInfo.length)
		
		addAdminRoleTableRow(roleInfo.adminUsers)
		
	})
	
	initButtonClickHandler('#userRoleNewRoleButton',function() {
		console.log("New Role button clicked")
		openNewUserRoleDialog()
	})
	
}