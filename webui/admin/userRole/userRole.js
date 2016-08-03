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




function initUserRoleSettings() {
	$('#userRoleTableBody').append(userRoleTableRowHTML("Role1"))
	$('#userRoleTableBody').append(userRoleTableRowHTML("Role2"))
	$('#userRoleTableBody').append(userRoleTableRowHTML("Role3"))

}