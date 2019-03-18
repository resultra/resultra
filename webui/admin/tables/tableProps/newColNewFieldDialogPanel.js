// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
var newTableColNewFieldDialogPanelID = "newField"

function createNewTableColNewFieldDialogPanelConfig(panelParams) {
	
	var newFieldPanel
	
	function initPanel($parentDialog) {
		
		var $panelForm = $('#newColNewFieldPanelForm')
		newFieldPanel = new NewFieldPanel(panelParams.databaseID,$panelForm)
		
		initButtonClickHandler('#newTableColNewFieldNextButton',function() {
			if (newFieldPanel.validateNewFieldParams()) {
				transitionToNextWizardDlgPanelByID($parentDialog,newTableColColTypeDialogPanelID)
			} // if validate form
		})
		initButtonClickHandler('#newTableColNewFieldPrevButton',function() {
			transitionToPrevWizardDlgPanelByPanelID(
				$parentDialog,newTableColCreateNewOrExistingFieldDialogPanelID)
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