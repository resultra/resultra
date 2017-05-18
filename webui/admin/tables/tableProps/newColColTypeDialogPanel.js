var newTableColColTypeDialogPanelID = "colType"

function createNewTableColColTypeDialogPanelConfig(panelParams) {
		
	function initPanel($parentDialog) {
		
		var $panelForm = $('#newColColTypePanelForm')
		
		initButtonClickHandler('#newTableColColTypeSaveButton',function() {
			$parentDialog.modal("hide")
		})
		initButtonClickHandler('#newTableColColTypePrevButton',function() {
			var newOrSelectedFieldPanelVals = getWizardDialogPanelVals(
					$parentDialog,newTableColCreateNewOrExistingFieldDialogPanelID)
				if(newOrSelectedFieldPanelVals.isNewField) {
					transitionToPrevWizardDlgPanelByPanelID(
						$parentDialog,newTableColNewFieldDialogPanelID)
				} else {
					transitionToPrevWizardDlgPanelByPanelID(
							$parentDialog,newTableColCreateNewOrExistingFieldDialogPanelID)
					
				}
		})
		
		
		
	}
	
	function getPanelValues() {
		return {}
	}
	
	var panelConfig = {
		panelID: newTableColColTypeDialogPanelID,
		divID: '#newColColTypePanel',
		progressPerc: 90,
		dlgButtons: null, // dialog buttons - TODO - reimplement with Bootstrap buttons
		initPanel: initPanel, // init panel
		getPanelVals: getPanelValues,
		transitionIntoPanel: function ($dialog) {
			setWizardDialogButtonSet("colTypeButtons")	
		}
	} // wizard dialog configuration for panel to create new field

	return panelConfig
	
}