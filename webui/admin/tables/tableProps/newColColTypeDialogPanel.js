var newTableColColTypeDialogPanelID = "colType"

function createNewTableColColTypeDialogPanelConfig(panelParams) {
	
	var $panelForm = $('#newColColTypePanelForm')
	var $colTypeSelection = $panelForm.find('select[name=colTypeSelection]')
		
	function initPanel($parentDialog) {
		
		var validator = $panelForm.validate({
			rules: {
				colTypeSelection: {
					required: true
				}
			},
			messages: {
				colTypeSelection: {
					required: "Column type is required"
				}
			}
		})
		validator.resetForm()
		
		
		
		initButtonClickHandler('#newTableColColTypeSaveButton',function() {
			
			function createNewColumn(fieldInfo) {
				
				function createNumberInput(fieldInfo) {
					var params = {
						parentTableID: panelParams.tableID,
						fieldID: fieldInfo.fieldID 
					}
					jsonAPIRequest("tableView/numberInput/new",params,function(numberInput) {
						console.log("Number input column created: " + JSON.stringify(numberInput))
					})
					
				}
				function createTextInput(fieldInfo) {
					var params = {
						parentTableID: panelParams.tableID,
						fieldID: fieldInfo.fieldID 
					}
					jsonAPIRequest("tableView/textInput/new",params,function(textInput) {
						console.log("Text input column created: " + JSON.stringify(textInput))
					})
					
				}
				
				console.log("Creating new column for field: " + JSON.stringify(fieldInfo))
				
				switch (fieldInfo.type) {
				case fieldTypeNumber:
					createNumberInput(fieldInfo)
					break
				case fieldTypeText:
					createTextInput(fieldInfo)
					break
				}
			}
			
			if ($panelForm.valid()) {
				var newOrSelectedFieldPanelVals = getWizardDialogPanelVals(
						$parentDialog,newTableColCreateNewOrExistingFieldDialogPanelID)
					if(newOrSelectedFieldPanelVals.isNewField) {
						var newFieldPanelVals = getWizardDialogPanelVals(
							$parentDialog,newTableColNewFieldDialogPanelID)
						newFieldPanelVals.newFieldPanel.createNewField(function(newFieldInfo) {
							if(newFieldInfo !== null) {
								createNewColumn(newFieldInfo)			
							}
						})
					} else {
						var selectedFieldID = newOrSelectedFieldPanelVals.selectedField				
						var getFieldParams = { fieldID: selectedFieldID }
						jsonAPIRequest("field/get",getFieldParams,function(existingFieldInfo) {
							createNewColumn(existingFieldInfo)
						})
					}
				
				
				
				$parentDialog.modal("hide")
			} // if validate form
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
	
	function transitionIntoPanel($dialog) {
		setWizardDialogButtonSet("colTypeButtons")
		
		function populateColTypeSelectionByFieldType(fieldType) {
			$colTypeSelection.empty()
			$colTypeSelection.append(defaultSelectOptionPromptHTML('Select a column type'))
			
			switch (fieldType) {
			case fieldTypeNumber:
				$colTypeSelection.append(selectOptionHTML('numberInput','Number input'))
				$colTypeSelection.append(selectOptionHTML('rating','Rating'))
				break
			case fieldTypeText:
				$colTypeSelection.append(selectOptionHTML('textInput','Text input'))
				break
			case fieldTypeBool:
				$colTypeSelection.append(selectOptionHTML('checkBox','Checkbox'))
				$colTypeSelection.append(selectOptionHTML('toggle','Toggle'))
				break
			}
		}
		
		// Populate the column type selection, depending on what type of field
		// the column is being linked to.
		var newOrSelectedFieldPanelVals = getWizardDialogPanelVals(
				$dialog,newTableColCreateNewOrExistingFieldDialogPanelID)
		if(newOrSelectedFieldPanelVals.isNewField) {
			var newFieldPanelVals = getWizardDialogPanelVals(
				$dialog,newTableColNewFieldDialogPanelID)
			var newFieldType = newFieldPanelVals.newFieldPanel.newFieldParams().type
			console.log("Configuring column type panel for new field: type = " + newFieldType)
			populateColTypeSelectionByFieldType(newFieldType)
		} else {
			var selectedFieldID = newOrSelectedFieldPanelVals.selectedField
			var getFieldParams = { fieldID: selectedFieldID }
			jsonAPIRequest("field/get",getFieldParams,function(fieldInfo) {
				var existingFieldType = fieldInfo.type
				console.log("Configuring column type panel for existing field: type = " + existingFieldType)
				populateColTypeSelectionByFieldType(existingFieldType)
			})
		}		
		
	}
	
	
	var panelConfig = {
		panelID: newTableColColTypeDialogPanelID,
		divID: '#newColColTypePanel',
		progressPerc: 90,
		dlgButtons: null, // dialog buttons - TODO - reimplement with Bootstrap buttons
		initPanel: initPanel, // init panel
		getPanelVals: getPanelValues,
		transitionIntoPanel: transitionIntoPanel
	} // wizard dialog configuration for panel to create new field

	return panelConfig
	
}