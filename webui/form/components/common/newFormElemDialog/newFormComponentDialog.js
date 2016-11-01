

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
		
		if(newOrExistingVals.componentValSelection == "newField") {
			var newFieldParams = getWizardDialogPanelVals($parentDialog,newFieldDialogPanelID)
			console.log("creating new field: params= " + JSON.stringify(newFieldParams))
			$parentDialog.modal("hide")
			// TODO Create the new field first, then create the component attached to this field
		} else if (newOrExistingVals.componentValSelection == "newGlobal") {
			// TODO Create the new global first, then create the component attached to this field
		} else if (newOrExistingVals.componentValSelection == "existingField") {
			var newComponentAPIParams = {
				parentFormID: newComponentParams.formID,
				geometry: newComponentParams.containerParams.geometry,
				componentLink: {
					linkedValType: "field",
					fieldID: newOrExistingVals.selectedFieldID
				}
			}
			componentCreated = true
			newComponentParams.createNewFormComponent($parentDialog,newComponentAPIParams)
		} else {
			assert(newOrExistingVals.componentValSelection == "existingGlobal")
			var newComponentAPIParams = {
				parentFormID: newComponentParams.formID,
				geometry: newComponentParams.containerParams.geometry,
				componentLink: {
					linkedValType: "global",
					globalID: newOrExistingVals.selectedGlobalID					
				}
			}
			componentCreated = true
			console.log("New Component params (existing global):" + JSON.stringify(newComponentAPIParams))
			newComponentParams.createNewFormComponent($parentDialog,newComponentAPIParams)
		}
	}
	
	var newOrExistingFieldPanel = createNewOrExistingFieldPanelContextBootstrap({
		parentTableID: newComponentParams.parentTableID,
		databaseID: newComponentParams.databaseID,
		elemPrefix:newComponentParams.elemPrefix,
		fieldTypes: newComponentParams.fieldTypes,
		globalTypes: newComponentParams.globalTypes,
		doneIfSelectExistingField:true,
		doneFunc:saveNewFormComponent})
	var newFieldPanel = createNewFieldDialogPanelContextBootstrap(newComponentParams)
		
		
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