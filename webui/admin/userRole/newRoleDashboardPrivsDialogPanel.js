var newRoleDashboardPrivsDialogPanelID = "dashboardPrivs"


function createNewRoleDashboardPrivsPanelContext(saveUserRoleFunc, dashboardsInfo) {
	
	var panelSelector = "#newUserRoleDialogDashboardPrivsPanel"
	
	var newFieldPanelConfig = {
		panelID: newRoleDashboardPrivsDialogPanelID,
		divID: panelSelector,
		progressPerc:80,
		initPanel: function ($parentDialog) {
			
			initButtonClickHandler('#newRoleDashboardPrivsPrevButton',function() {
				console.log("Prev button clicked")
				transitionToPrevWizardDlgPanelByPanelID($parentDialog,newRoleRoleNameDialogPanelID)
			})
			
			
			initButtonClickHandler('#newRoleDashboardPrivsSaveButton',function() {
				console.log("Save button clicked")
				saveUserRoleFunc($parentDialog)	
			})
			
			initRoleDashboardPrivSettingsTable(dashboardsInfo)
						
		}, // init panel
		transitionIntoPanel: function ($dialog) { 
			
			setWizardDialogButtonSet('newRoleDashboardPrivsButtons')
			
			var $newRoleRoleNamePanelForm = $('#newUserRoleDialoggDashboardPrivsForm')
						
		},
		getPanelVals: function () {
			return getDashboardRolePrivRadioButtonVals()
		}

	} // wizard dialog configuration for panel to create new field
	
	return newFieldPanelConfig;
	
}
