function openNewUserRoleDialog() {
	
	initRoleFormPrivSettingsTable()
	initRoleDashboardPrivSettingsTable()
	$('#newUserRoleDialog').modal('show')
	
	initButtonClickHandler('#newUserRoleSaveButton',function() {
		console.log("Save button clicked")
		
		var formPrivs = getFormPrivSettingVals()
				
		$('#newUserRoleDialog').modal('hide')
	})
	
}