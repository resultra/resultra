var newRoleFormPrivsDialogPanelID = "formPrivs"


function createNewRoleFormPrivsPanelContext(formsInfo) {
	
	var panelSelector = "#newUserRoleDialogFormPrivsPanel"
	
	var newFieldPanelConfig = {
		panelID: newRoleFormPrivsDialogPanelID,
		divID: panelSelector,
		progressPerc:60,
		initPanel: function ($parentDialog) {
			
			initButtonClickHandler('#newRoleFormPrivsPrevButton',function() {
				console.log("Prev button clicked")
				transitionToPrevWizardDlgPanelByPanelID($parentDialog,newRoleRoleNameDialogPanelID)
			})
			
			
			initButtonClickHandler('#newRoleFormPrivsNextButton',function() {
				console.log("Next button clicked")
				transitionToNextWizardDlgPanelByID($parentDialog,newRoleDashboardPrivsDialogPanelID)
			})
			
			initRoleFormPrivSettingsTable(formsInfo)
			
					
		}, // init panel
		transitionIntoPanel: function ($dialog) { 
			
			setWizardDialogButtonSet('newRoleFormPrivsButtons')
			
			var $newRoleRoleNamePanelForm = $('#newUserRoleDialogFormPrivsForm')
						
		},
		getPanelVals: function () {
			return getFormRolePrivRadioButtonVals()
		}
	} // wizard dialog configuration for panel to create new field
	
	return newFieldPanelConfig;
	
}
