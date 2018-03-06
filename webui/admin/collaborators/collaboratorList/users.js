
function userListTableRow(pageContext,userRoleInfo) {
	
	var roles = ""
	if(userRoleInfo.isAdmin) {
		roles = roles + "Administrator"
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
		
	var $userRow = $("#collaboratorListRowTemplate").clone()
	$userRow.attr("id","")
	
	var $nameCell = $userRow.find(".userName")
	$nameCell.text(userNameDisplay)
	
	var $rolesCell = $userRow.find(".userRoles")
	$rolesCell.text(roles)
		
	if(userRoleInfo.isAdmin) {
		$userRow.find(".userButtons").empty()
	} else {
		var $deleteButton = $userRow.find(".deleteCollaboratorButton")
		initButtonControlClickHandler($deleteButton,function() {
			console.log("Remove collaborator button clicked")
			openConfirmDeleteDialog("collaborator",function() {
	
				var deleteParams = {
					collaboratorID: userRoleInfo.collaboratorID,
					databaseID: pageContext.databaseID
				}
				jsonAPIRequest("admin/deleteCollaborator",deleteParams,function(replyStatus) {
					$deleteButton.closest("tr").remove()
					console.log("Delete confirmed")
				})
				
			})
		})
		
		var $editButton = $userRow.find(".editUserButton")
		var editCollabPropsURL = '/admin/collaborator/' + pageContext.databaseID + '/' + userRoleInfo.collaboratorID
		setPageContentButtonClickHandler($editButton,editCollabPropsURL,function() {
			initCollaboratorPropsAdminSettingsPageContent(pageContext,userRoleInfo)
		})	
		
	}
	
	return $userRow
	
}

function initUserListSettings(pageContext) {	
	
	var getRoleInfoParams = { databaseID: pageContext.databaseID }
	jsonAPIRequest("admin/getUserRoleInfo",getRoleInfoParams,function(userRoleInfo) {
		console.log("Got role info: " + JSON.stringify(userRoleInfo))
		console.log("Number of roles: " + userRoleInfo.length)
		
		for (var userRoleIndex = 0; userRoleIndex < userRoleInfo.length; userRoleIndex++) {
			var currUserRole = userRoleInfo[userRoleIndex]
			console.log("appending user role")
			$('#userListTableBody').append(userListTableRow(pageContext,currUserRole))
		}
	})
	
	initButtonClickHandler('#addUserButton',function() {
		console.log("Add new user button clicked")
		openNewUserDialog(pageContext)
	})

}