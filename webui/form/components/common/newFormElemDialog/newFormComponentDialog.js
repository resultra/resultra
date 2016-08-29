

function openNewFormComponentDialog(newComponentParams) {


	var dialogSelector = '#' + newComponentParams.elemPrefix + "NewFormComponentDialog"
	var progressSelector = '#' + newComponentParams.elemPrefix + 'WizardDialogProgress'
	
	// Use this flag to track whether or not the user canceled or closed the dialog, or 
	// continued on the actually finish creating the component. If the component is created,
	// The flag is set to true, which prevents the component's container from being removed
	// when the dialog is closed.
	var componentCreated = false

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
			componentCreated = true
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
		
		
	openWizardDialog({
		closeFunc: function() {
			if(!componentCreated) {
				var newComponentPlaceholderSelector = '#' + newComponentParams.containerParams.containerID
				console.log("Cancel new component creation: removing placholder component = " 
									+ newComponentPlaceholderSelector)
				$(newComponentPlaceholderSelector).remove()				
			}
			
      	},
		dialogDivID: dialogSelector,
		panels: [newOrExistingFieldPanel,newFieldPanel],
		progressDivID: progressSelector,
	})
		
}