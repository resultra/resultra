
function openNewUserRoleDialog(pageContext) {

	var $roleNameInput = $('#newRoleNameInput')
	var $newRoleDialog = $('#newUserRoleDialog')
	
	var $newRoleRoleNamePanelForm = $('#newUserRoleDialogRoleNameForm')
	
	var validator = $newRoleRoleNamePanelForm.validate({
		rules: {
			newRoleNameInput: {
				minlength: 3,
				required: true,
				remote: {
					url: '/api/userRole/validateRoleName',
					data: {
						roleName: function() { return $roleNameInput.val(); }
					}
				} // remote
			} // newRoleNameInput
		},
	})
	
	$newRoleDialog.modal('show')
	
	initButtonClickHandler('#newRoleSaveButton',function() {
		console.log("Save button clicked")
		if($newRoleRoleNamePanelForm.valid()) {
			var newRoleParams = {
				databaseID: pageContext.databaseID,
				roleName: $roleNameInput.val()
			}
			console.log("Saving new user role: params=" + JSON.stringify(newRoleParams))
		
			jsonAPIRequest("userRole/newRole",newRoleParams,function(newRoleInfo) {
		
				$newRoleDialog.modal('hide')	
				
				var userRoleContentURL = '/admin/userRole/' + newRoleInfo.roleID
				setSettingsPageContent(userRoleContentURL, function() {
					initUserRolePropsAdminSettingsPageContent(pageContext,newRoleInfo)
				})
		
			})
		}
	})
		
}