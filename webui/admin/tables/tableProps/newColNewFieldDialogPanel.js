var newTableColNewFieldDialogPanelID = "newField"

function createNewTableColNewFieldDialogPanelConfig(panelParams) {
	
	var newFieldPanel
	
	function initPanel($parentDialog) {
		
		var $panelForm = $('#newColNewFieldPanelForm')
		newFieldPanel = new NewFieldPanel(panelParams.databaseID,$panelForm)
		
		initButtonClickHandler('#newTableColNewFieldNextButton',function() {
			if (newFieldPanel.validateNewFieldParams()) {
				$parentDialog.modal("hide")
//				transitionToNextWizardDlgPanelByID($parentDialog,newFieldDialogPanelID)
			} // if validate form
		})
		
	}
	
	function getPanelValues() {
		return {
			newFieldPanel: newFieldPanel
		}
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