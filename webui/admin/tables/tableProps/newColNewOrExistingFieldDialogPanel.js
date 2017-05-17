var newTableColCreateNewOrExistingFieldDialogPanelID = "newOrExistingField"

function createNewTableColNewOrExistingDialogPanelConfig(panelParams) {
	
	
	function initPanel($parentDialog) {
		
		var $panelForm = $('#newColNewOrExistingFieldPanelForm')
		var $newOrExistingRadio = $panelForm.find("input[name=newOrExistingRadio]")
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
				transitionToNextWizardDlgPanelByID($parentDialog,newTableColNewFieldDialogPanelID)
			} // if validate form
		})
		
	}
	
	function getPanelValues() {
		return {
			isNewField: isNewField(),
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