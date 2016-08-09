
var newUserRoleDialogProgressSelector = "#" + "userRole" + "WizardDialogProgress"

function openNewUserRoleDialog() {
	
	initRoleFormPrivSettingsTable()
	initRoleDashboardPrivSettingsTable()


	var dialogSelector = '#newUserRoleDialog'
	var roleNamePanel = createNewRoleRoleNamePanelContext()
	var formPrivsPanel = createNewRoleFormPrivsPanelContext()
	var dashboardPrivsPanel = createNewRoleDashboardPrivsPanelContext()
	
	openWizardDialog({
		closeFunc: function() {
			console.log("Close dialog")
      	},
		dialogDivID: dialogSelector,
		panels: [roleNamePanel,formPrivsPanel,dashboardPrivsPanel],
		progressDivID: newUserRoleDialogProgressSelector,
	})
	
	var $newRoleForm = $('#newUserRoleDialogForm')		
}