
var newUserRoleDialogProgressSelector = "#" + "userRole" + "WizardDialogProgress"

function openNewUserRoleDialog() {


	function saveNewUserRole($dialog) {
		
		var newUserRoleParams = {
			roleName: getWizardDialogPanelVals($dialog,newRoleRoleNameDialogPanelID),
			formPrivs: getWizardDialogPanelVals($dialog,newRoleFormPrivsDialogPanelID),
			dashboardPrivs: getWizardDialogPanelVals($dialog,newRoleDashboardPrivsDialogPanelID)
		} 
		console.log("Saving new user role: params=" + JSON.stringify(newUserRoleParams))
		
		$('#newUserRoleDialog').modal('hide')	
	}

	var dialogSelector = '#newUserRoleDialog'
	var roleNamePanel = createNewRoleRoleNamePanelContext()
	var formPrivsPanel = createNewRoleFormPrivsPanelContext()
	var dashboardPrivsPanel = createNewRoleDashboardPrivsPanelContext(saveNewUserRole)
	
	openWizardDialog({
		closeFunc: function() {
			console.log("Close dialog")
      	},
		dialogDivID: dialogSelector,
		panels: [roleNamePanel,formPrivsPanel,dashboardPrivsPanel],
		progressDivID: newUserRoleDialogProgressSelector,
		minBodyHeight:'350px'
	})
	
	var $newRoleForm = $('#newUserRoleDialogForm')		
}