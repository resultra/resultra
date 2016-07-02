

function openNewFormComponentDialog(elemPrefix, formID,parentTableID,containerParams) {

	var dialogSelector = '#' + elemPrefix + "NewFormComponentDialog"
	var progressSelector = '#' + elemPrefix + 'WizardDialogProgress'

	function saveNewFormComponent($parentDialog) {
		console.log("saveNewFormComponent: done handler called")
	}
	
	var newOrExistingFieldPanel = createNewOrExistingFieldPanelContextBootstrap({
		parentTableID: parentTableID,
		elemPrefix:elemPrefix,
		fieldTypes: [fieldTypeBool],
		doneIfSelectExistingField:true,
		doneFunc:saveNewFormComponent})
	var newFieldPanel = createNewFieldDialogPanelContextBootstrap(elemPrefix)
		
		
	openWizardDialogBootstrap({
		closeFunc: function() {
			console.log("Close dialog")
			if(!newCheckBoxParams.containerCreated)
			{
			  // If the the text box creation is not complete, remove the placeholder
			  // from the canvas.
				$('#'+newCheckBoxParams.placeholderID).remove()
			}
      	},
		dialogDivID: dialogSelector,
		panels: [newOrExistingFieldPanel,newFieldPanel],
		progressDivID: progressSelector,
	})
		
}