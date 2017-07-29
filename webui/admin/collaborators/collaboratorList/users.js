function userListItemButtonsHTML(databaseID, collaboratorID) {
	
	
		var editCollabPropsURL = '/admin/collaborator/' + databaseID + '/' + collaboratorID
	
return '' +
			'<div class="pull-right userListItemButtons">' + 
	
  			'<a class="btn btn-xs editUserRoleButton" href="'+  editCollabPropsURL + '" role="button">' + 
				// padding-bottom: 2px makes the button image vertically line up better.
				'<span class="glyphicon glyphicon-pencil" style="padding-bottom:2px;"></span>' +
			'</a>' +
  			'<button class="btn btn-xs btn-danger deleteUserRoleButton">' + 
				// padding-bottom: 2px makes the button image vertically line up better.
				'<span class="glyphicon glyphicon-remove" style="padding-bottom:2px;"></span>' +
			'</button>';

			'</div>'

	
}

function userListTableRowHTML(databaseID,userRoleInfo) {
	
	var roles = ""
	if(userRoleInfo.isAdmin) {
		roles = roles + "<strong>Administrator</strong>"
	} else {
		var roleNames = []
		for (var roleIndex = 0; roleIndex < userRoleInfo.customRoles.length; roleIndex++) {
			var currRoleInfo = userRoleInfo.customRoles[roleIndex]
			roleNames.push(currRoleInfo.roleName)
		}
		roles = roleNames.join(", ")
	}
	
	var userInfo = userRoleInfo.userInfo
	var userNameDisplay = '@' + userInfo.userName + 
		" (" + userInfo.firstName + " " + userInfo.lastName + ")"
	
	var buttonsHTML = userListItemButtonsHTML(databaseID,userRoleInfo.collaboratorID)
	
	return '' +
		'<tr class="userListRow">' +
	         '<td>' + userNameDisplay +  '</td>' +
	         '<td>' + roles +  '</td>' +
	         '<td class="userListButtonCell">' + buttonsHTML + '</td>' +
	     '</tr>'
	
}

function initUserListSettings(databaseID) {	
	
	var getRoleInfoParams = { databaseID: databaseID }
	jsonAPIRequest("admin/getUserRoleInfo",getRoleInfoParams,function(userRoleInfo) {
		console.log("Got role info: " + JSON.stringify(userRoleInfo))
		console.log("Number of roles: " + userRoleInfo.length)
		
		for (var userRoleIndex = 0; userRoleIndex < userRoleInfo.length; userRoleIndex++) {
			var currUserRole = userRoleInfo[userRoleIndex]
			console.log("appending user role")
			$('#userListTableBody').append(userListTableRowHTML(databaseID,currUserRole))
		}

		
	})
	
	initButtonClickHandler('#addUserButton',function() {
		console.log("Add new user button clicked")
		openNewUserDialog(databaseID)
	})

}