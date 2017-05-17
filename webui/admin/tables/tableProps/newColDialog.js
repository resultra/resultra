function openNewTableColDialog() {
	
	var newOrExistingPanel = createNewTableColNewOrExistingDialogPanelConfig()
	var newFieldPanel = createNewTableColNewFieldDialogPanelConfig()
		
	openWizardDialog({
		closeFunc: function() {},
		dialogDivID: '#newTableColDialog',
		panels: [newOrExistingPanel,newFieldPanel],
		progressDivID: '#tableProps_WizardDialogProgress',
	})
	
}