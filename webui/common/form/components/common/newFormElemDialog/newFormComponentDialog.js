

function openNewFormComponentDialog(newComponentParams) {


	var dialogSelector = '#' + newComponentParams.elemPrefix + "NewFormComponentDialog"
	var progressSelector = '#' + newComponentParams.elemPrefix + 'WizardDialogProgress'
	
	// Use this flag to track whether or not the user canceled or closed the dialog, or 
	// continued on the actually finish creating the component. If the component is created,
	// The flag is set to true, which prevents the component's container from being removed
	// when the dialog is closed.
	var componentCreated = false

	function doneCreatingComponent($parentDialog) {
		componentCreated = true
		$parentDialog.modal("hide")
	}
	
	// Create the wizard dialog panels
	
	var newOrExistingFieldPanelParams = {
		doneIfSelectExistingField:true,
		doneFunc:doneCreatingComponent
	}
	$.extend(newOrExistingFieldPanelParams,newComponentParams)
	var newOrExistingFieldPanel = createNewOrExistingFieldPanelContextBootstrap(newOrExistingFieldPanelParams)
		
	var newFieldPanelParams = {
		doneFunc:doneCreatingComponent,
		hideCreateCalcFieldCheckbox: false
	}
	$.extend(newFieldPanelParams,newComponentParams)			
	var newFieldPanel = createNewFieldDialogPanelContextBootstrap(newFieldPanelParams)
		
		
	openWizardDialog({
		closeFunc: function() {
			if(!componentCreated) {
				console.log("Cancel new component creation: removing placeholder component")
				var $componentContainer = newComponentParams.containerParams.containerObj
				$componentContainer.remove()
				pruneComponentLayoutEmptyColsAndRows($(formDesignCanvasSelector))
			}
			
      	},
		dialogDivID: dialogSelector,
		panels: [newOrExistingFieldPanel,newFieldPanel],
		progressDivID: progressSelector,
	})
		
}