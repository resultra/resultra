
function openNewUserRoleDialog(databaseID) {

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
				databaseID: databaseID,
				roleName: $roleNameInput.val()
			}
			console.log("Saving new user role: params=" + JSON.stringify(newRoleParams))
		
			jsonAPIRequest("userRole/newRole",newRoleParams,function(newRoleInfo) {
		
				$newRoleDialog.modal('hide')	
				navigateToURL('/admin/userRole/' + newRoleInfo.roleID)
		
			})
		}
	})
		
}