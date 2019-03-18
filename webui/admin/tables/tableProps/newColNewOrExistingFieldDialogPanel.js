// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
var newTableColCreateNewOrExistingFieldDialogPanelID = "newOrExistingField"

function createNewTableColNewOrExistingDialogPanelConfig(panelParams) {
	
	var $panelForm = $('#newColNewOrExistingFieldPanelForm')
	var $selectField = $panelForm.find("select[name=existingFieldSelection]")
	
	function newFieldSelected() {
		var $checkedNewOrExistingRadio = $panelForm.find("input[name=newOrExistingRadio]:checked")
		var checkedVal = $checkedNewOrExistingRadio.val()
		if(checkedVal === 'newField') {
			return true
		} else {
			return false
		}	
	}

	function initPanel($parentDialog) {
		
		var $newOrExistingRadio = $panelForm.find("input[name=newOrExistingRadio]")
			
		var validator = $panelForm.validate({
			rules: {
				existingFieldSelection: {
					required: {
						depends: function(element) {
							return !newFieldSelected()
						}					
					}
				}
			},
			messages: {
				existingFieldSelection: {
					required: "Select an existing field"
				}
			}
		})
		
		validator.resetForm()
		
			
		$newOrExistingRadio.change(function() {
			if(newFieldSelected()) {
				$selectField.attr('disabled',true)
			} else {
				$selectField.attr('disabled',false)
			}
		})
		
		$selectField.dropdown()
		loadSortedFieldInfo(panelParams.databaseID,[fieldTypeAll],function(sortedFields) {
			populateSortedFieldSelectionMenu($selectField,sortedFields)
		})
		
		
		
		initButtonClickHandler('#newTableColNextButton',function() {
			if ($panelForm.valid()) {
				if(newFieldSelected()) {
					transitionToNextWizardDlgPanelByID($parentDialog,newTableColNewFieldDialogPanelID)
				} else {
					transitionToNextWizardDlgPanelByID($parentDialog,newTableColColTypeDialogPanelID)
				}
			} // if validate form
		})
		
	}
	
	function getPanelValues() {
		return {
			isNewField: newFieldSelected(),
			selectedField: $selectField.val()
		}
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