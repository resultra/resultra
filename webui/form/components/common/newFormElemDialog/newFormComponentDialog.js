

function openNewFormComponentDialog(newComponentParams) {


	var dialogSelector = '#' + newComponentParams.elemPrefix + "NewFormComponentDialog"
	var progressSelector = '#' + newComponentParams.elemPrefix + 'WizardDialogProgress'

	function saveNewFormComponent($parentDialog) {
		console.log("saveNewFormComponent: done handler called")
		
		var newOrExistingVals = getWizardDialogPanelVals($parentDialog,createNewOrExistingFieldDialogPanelID)
				
		console.log("saveNewFormComponent: panel values: " + JSON.stringify(newOrExistingVals))
		
		if(newOrExistingVals.newField == true) {
			// TODO Create the new field first, then create the component attached to this field
		} else {
			var newComponentAPIParams = {
				fieldParentTableID: newComponentParams.parentTableID,
				parentFormID: newComponentParams.formID,
				geometry: newComponentParams.containerParams.geometry,
				fieldID: newOrExistingVals.selectedFieldID
			}
			newComponentParams.createNewFormComponent($parentDialog,newComponentAPIParams)
		}
	}
	
	var newOrExistingFieldPanel = createNewOrExistingFieldPanelContextBootstrap({
		parentTableID: newComponentParams.parentTableID,
		elemPrefix:newComponentParams.elemPrefix,
		fieldTypes: newComponentParams.fieldTypes,
		doneIfSelectExistingField:true,
		doneFunc:saveNewFormComponent})
	var newFieldPanel = createNewFieldDialogPanelContextBootstrap(newComponentParams.elemPrefix)
		
		
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