var newTableColNewFieldDialogPanelID = "newField"

function createNewTableColNewFieldDialogPanelConfig() {
	
	
	function initPanel($parentDialog) {
		
		var $panelForm = $('#newColNewFieldPanelForm')
		
		var validator = $panelForm.validate({
			rules: {} 
		})
		
		validator.resetForm()
		
		
		initButtonClickHandler('#newTableColNewFieldNextButton',function() {
			if ($panelForm.valid()) {
				$parentDialog.modal("hide")
//				transitionToNextWizardDlgPanelByID($parentDialog,newFieldDialogPanelID)
			} // if validate form
		})
		
	}
	
	function getPanelValues() {
		return {}
	}
	
	var panelConfig = {
		panelID: newTableColNewFieldDialogPanelID,
		divID: '#newColNewFieldPanel',
		progressPerc: 60,
		dlgButtons: null, // dialog buttons - TODO - reimplement with Bootstrap buttons
		initPanel: initPanel, // init panel
		getPanelVals: getPanelValues,
		transitionIntoPanel: function ($dialog) {
			setWizardDialogButtonSet("newFieldButtons")	
		}
	} // wizard dialog configuration for panel to create new field

	return panelConfig
	
}