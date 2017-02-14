
var newUserRoleDialogProgressSelector = "#" + "userRole" + "WizardDialogProgress"

function openNewUserRoleDialog(databaseID) {


	function saveNewUserRole($dialog) {
		
		var newUserRoleParams = {
			databaseID: databaseID,
			roleName: getWizardDialogPanelVals($dialog,newRoleRoleNameDialogPanelID),
			formPrivs: getWizardDialogPanelVals($dialog,newRoleFormPrivsDialogPanelID),
			dashboardPrivs: getWizardDialogPanelVals($dialog,newRoleDashboardPrivsDialogPanelID)
		} 
		console.log("Saving new user role: params=" + JSON.stringify(newUserRoleParams))
		
		jsonAPIRequest("userRole/newRole",newUserRoleParams,function(response) {
		
			$('#newUserRoleDialog').modal('hide')	
		
		})
		
		
	}

	var dbInfoParams = { databaseID: databaseID }
	jsonAPIRequest("database/getInfo",dbInfoParams,function(databaseInfo) {
	
		var dialogSelector = '#newUserRoleDialog'
		var roleNamePanel = createNewRoleRoleNamePanelContext()
		var formPrivsPanel = createNewRoleFormPrivsPanelContext(databaseInfo.formsInfo)
		var dashboardPrivsPanel = createNewRoleDashboardPrivsPanelContext(saveNewUserRole,databaseInfo.dashboardsInfo)
	
		openWizardDialog({
			closeFunc: function() {
				console.log("Close dialog")
	      	},
			dialogDivID: dialogSelector,
			panels: [roleNamePanel,formPrivsPanel,dashboardPrivsPanel],
			progressDivID: newUserRoleDialogProgressSelector,
			minBodyHeight:'350px'
		})
	
	})
	
		
}