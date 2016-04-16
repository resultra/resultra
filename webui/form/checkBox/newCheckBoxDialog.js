
var checkboxDialogSelector = "#newCheckbox"

function openNewCheckboxDialog(containerParams)
{
	
	// Must be the same as designForm.go - this is the common prefix on all DOM element IDs to distinguish
	// checkbox related elements from other form elements.
	var checkboxElemPrefix = "checkbox_"
	var dialogProgressSelector = '#' + checkboxElemPrefix + 'NewFormElemDialogProgress'
	
	newCheckBoxParams = {
		containerParams: containerParams,
		containerCreated: false,
		placeholderID: containerParams.containerID,
		dialogBox: $( checkboxDialogSelector )
	}
	
	// Enable Semantic UI checkboxes and popups
	$('.ui.checkbox').checkbox();
	$('.ui.radio.checkbox').checkbox();	
	
	
	var newOrExistingFieldPanel = createNewOrExistingFieldPanelConfig(checkboxElemPrefix)
	
	openWizardDialog({
		closeFunc: function() {
			console.log("Close dialog")
			if(!newCheckBoxParams.containerCreated)
			{
			  // If the the text box creation is not complete, remove the placeholder
			  // from the canvas.
				$('#'+newCheckBoxParams.placeholderID).remove()
			}
      	},
		width: 500, height: 500,
		dialogDivID: checkboxDialogSelector,
		panels: [newOrExistingFieldPanel],
		progressDivID: dialogProgressSelector,
	})
		
} // newLayoutContainer

function initNewCheckBoxDialog() {
	// Initialize the newTextBox dialog with the minimum parameters. This is necessary
	// to hide the dialog from view when the document is initially loaded. The
	// dialog is fully re-initialized just prior to it being opened.
	initWizardDialog(checkboxDialogSelector)
}