var newRoleRoleNameDialogPanelID = "roleName"


function createNewRoleRoleNamePanelContext() {
	
	var panelSelector = "#newUserRoleDialogRoleNamePanel"
		
	var newFieldPanelConfig = {
		panelID: newRoleRoleNameDialogPanelID,
		divID: panelSelector,
		progressPerc:20,
		initPanel: function ($parentDialog) {
			
			initButtonClickHandler('#newRoleRoleNameNextButton',function() {
				console.log("Next button clicked")
				var $newRoleRoleNamePanelForm = $('#newUserRoleDialogRoleNameForm')
				if($newRoleRoleNamePanelForm.valid()) {
					var formPrivs = getFormPrivSettingVals()
					transitionToNextWizardDlgPanelByID($parentDialog,newRoleFormPrivsDialogPanelID)
				}
			})			
		}, // init panel
		transitionIntoPanel: function ($dialog) { 
			
			setWizardDialogButtonSet('newRoleRoleNameButtons')
			
			var $newRoleRoleNamePanelForm = $('#newUserRoleDialogRoleNameForm')
			
			var validator = $newRoleRoleNamePanelForm.validate({
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
	} // wizard dialog configuration for panel to create new field
	
	return newFieldPanelConfig;
	
}
