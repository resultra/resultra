function openNewUserRoleDialog() {
	
	initRoleFormPrivSettingsTable()
	initRoleDashboardPrivSettingsTable()
	$('#newUserRoleDialog').modal('show')
	
	var $newRoleForm = $('#newUserRoleDialogForm')
	
	initButtonClickHandler('#newUserRoleSaveButton',function() {
		console.log("Save button clicked")
		
		if($newRoleForm.valid()) {
			var formPrivs = getFormPrivSettingVals()
			$('#newUserRoleDialog').modal('hide')		
		}
				
	})
	
	var validator = $newRoleForm.validate({
		rules: {
			newRoleNameInput: {
				minlength: 3,
				required: true,
				remote: {
					url: '/api/userRole/validateRoleName',
					data: {
						roleName: function() { return $('#newRoleNameInput').val(); }
					}
				} // remote
			} // newRoleNameInput
		},
	})
	
	validator.resetForm()
	
}