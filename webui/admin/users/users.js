function userListItemButtonsHTML() {
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

function userListTableRowHTML(userID) {
	
	var roles = "TBD"
	
	var buttonsHTML = userListItemButtonsHTML()
	
	return '' +
		'<tr class="userListRow">' +
	         '<td>' + userID +  '</td>' +
	         '<td>' + roles +  '</td>' +
	         '<td class="userListButtonCell">' + buttonsHTML + '</td>' +
	     '</tr>'
	
}

function initUserListSettings() {	
	$('#userListTableBody').append(userListTableRowHTML('user1'))
	$('#userListTableBody').append(userListTableRowHTML('user2'))
	$('#userListTableBody').append(userListTableRowHTML('user3'))

}