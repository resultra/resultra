function openNewTableColDialog(databaseID) {
	
	function saveNewTableCol($dialog) {
		
/*		var newTableColParams = { ... }
		console.log("Saving new table column: params=" + JSON.stringify(newTableColParams))
		
		jsonAPIRequest("tableView/newColumn",newTableColParams,function(response) {
		
			$dialog.modal('hide')	
		
		})
		*/
		
	}
	
	var panelConfig = {
		databaseID: databaseID
	}
	
	
	var newOrExistingPanel = createNewTableColNewOrExistingDialogPanelConfig(panelConfig)
	var newFieldPanel = createNewTableColNewFieldDialogPanelConfig(panelConfig)
	var colTypePanel = createNewTableColColTypeDialogPanelConfig(panelConfig)
		
	openWizardDialog({
		closeFunc: function() {},
		dialogDivID: '#newTableColDialog',
		panels: [newOrExistingPanel,newFieldPanel,colTypePanel],
		progressDivID: '#tableProps_WizardDialogProgress',
	})
	
}