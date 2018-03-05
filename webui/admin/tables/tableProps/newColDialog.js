function openNewTableColDialog(pageContext,tableRef) {
		
	var panelConfig = {
		databaseID: tableRef.parentDatabaseID,
		tableID: tableRef.tableID
	}
	
	
	var newOrExistingPanel = createNewTableColNewOrExistingDialogPanelConfig(panelConfig)
	var newFieldPanel = createNewTableColNewFieldDialogPanelConfig(panelConfig)
	var colTypePanel = createNewTableColColTypeDialogPanelConfig(pageContext,panelConfig)
		
	openWizardDialog({
		closeFunc: function() {},
		dialogDivID: '#newTableColDialog',
		panels: [newOrExistingPanel,newFieldPanel,colTypePanel],
		progressDivID: '#tableProps_WizardDialogProgress',
	})
	
}