var newTableColCreateNewOrExistingFieldDialogPanelID = "newOrExistingField"

function createNewTableColNewOrExistingDialogPanelConfig() {
	
	
	function initPanel($parentDialog) {
		
		var $panelForm = $('#newColNewOrExistingFieldPanelForm')
		
		var validator = $panelForm.validate({
			rules: {} 
		})
		
		validator.resetForm()
		
		
		initButtonClickHandler('#newTableColNextButton',function() {
			if ($panelForm.valid()) {
				transitionToNextWizardDlgPanelByID($parentDialog,newTableColNewFieldDialogPanelID)
			} // if validate form
		})
		
	}
	
	function getPanelValues() {
		return {}
	}
	
	var panelConfig = {
		panelID: newTableColCreateNewOrExistingFieldDialogPanelID,
		divID: '#newColNewOrExistingFieldPanel',
		progressPerc: 20,
		dlgButtons: null, // dialog buttons - TODO - reimplement with Bootstrap buttons
		initPanel: initPanel, // init panel
		getPanelVals: getPanelValues,
		transitionIntoPanel: function ($dialog) {
			setWizardDialogButtonSet("newExistingFieldButtons")	
		}
	} // wizard dialog configuration for panel to create new field

	return panelConfig;
	
}