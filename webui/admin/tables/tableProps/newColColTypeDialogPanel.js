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

				function createRating(fieldInfo) {
					var params = {
						parentTableID: panelParams.tableID,
						fieldID: fieldInfo.fieldID 
					}
					jsonAPIRequest("tableView/rating/new",params,function(rating) {
						console.log("Number rating column created: " + JSON.stringify(rating))
					})
					
				}
				
				function createProgress(fieldInfo) {
					var params = {
						parentTableID: panelParams.tableID,
						fieldID: fieldInfo.fieldID 
					}
					jsonAPIRequest("tableView/progress/new",params,function(progress) {
						console.log("Progress indicator column created: " + JSON.stringify(progress))
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

				function createCommentInput(fieldInfo) {
					var params = {
						parentTableID: panelParams.tableID,
						fieldID: fieldInfo.fieldID 
					}
					jsonAPIRequest("tableView/comment/new",params,function(commentInput) {
						console.log("Comment input column created: " + JSON.stringify(commentInput))
					})
					
				}
				
				function createNoteEditorInput(fieldInfo) {
					var params = {
						parentTableID: panelParams.tableID,
						fieldID: fieldInfo.fieldID 
					}
					jsonAPIRequest("tableView/note/new",params,function(noteEditor) {
						console.log("Note editor input column created: " + JSON.stringify(noteEditor))
					})
					
				}
				
				function createAttachmentInput(fieldInfo) {
					var params = {
						parentTableID: panelParams.tableID,
						fieldID: fieldInfo.fieldID 
					}
					jsonAPIRequest("tableView/attachment/new",params,function(attachmentRef) {
						console.log("Attachment column created: " + JSON.stringify(attachmentRef))
					})
					
				}
				
				function createUserInput(fieldInfo) {
					var params = {
						parentTableID: panelParams.tableID,
						fieldID: fieldInfo.fieldID 
					}
					jsonAPIRequest("tableView/userSelection/new",params,function(userSelection) {
						console.log("User selection input column created: " + JSON.stringify(userSelection))
					})
					
				}
				function createSocialButton(fieldInfo) {
					var params = {
						parentTableID: panelParams.tableID,
						fieldID: fieldInfo.fieldID 
					}
					jsonAPIRequest("tableView/socialButton/new",params,function(socialButton) {
						console.log("Social button column created: " + JSON.stringify(userSelection))
					})
					
				}
				
				function createDatePickerInput(fieldInfo) {
					var params = {
						parentTableID: panelParams.tableID,
						fieldID: fieldInfo.fieldID 
					}
					jsonAPIRequest("tableView/datePicker/new",params,function(datePicker) {
						console.log("Date picker column created: " + JSON.stringify(datePicker))
					})
					
				}

				function createCheckBoxInput(fieldInfo) {
					var params = {
						parentTableID: panelParams.tableID,
						fieldID: fieldInfo.fieldID 
					}
					jsonAPIRequest("tableView/checkBox/new",params,function(checkBox) {
						console.log("Check box column created: " + JSON.stringify(checkBox))
					})
				}
				
				function createToggle(fieldInfo) {
					var params = {
						parentTableID: panelParams.tableID,
						fieldID: fieldInfo.fieldID 
					}
					jsonAPIRequest("tableView/toggle/new",params,function(toggle) {
						console.log("Toggle column created: " + JSON.stringify(toggle))
					})
				}

				console.log("Creating new column for field: " + JSON.stringify(fieldInfo))
				
				var colType = $colTypeSelection.val()
				
				switch (fieldInfo.type) {
				case fieldTypeNumber:
					if (colType==='numberInput') {
						createNumberInput(fieldInfo)	
					} else if (colType === 'rating'){
						createRating(fieldInfo)
					} else if (colType === 'progress'){
						createProgress(fieldInfo)
					} else {
						console.log("Unknown column type for number field : " + colType)
					}
					break
				case fieldTypeUser:
					
					if (colType==='userSelection') {
						createUserInput(fieldInfo)
					} else if (colType === 'socialButton'){
						createSocialButton(fieldInfo)
					} else {
						console.log("Unknown column type for number field : " + colType)
					}
					break
				case fieldTypeText:
					createTextInput(fieldInfo)
					break
				case fieldTypeComment:
					createCommentInput(fieldInfo)
					break
				case fieldTypeAttachment:
					createAttachmentInput(fieldInfo)
					break
				case fieldTypeLongText:
					createNoteEditorInput(fieldInfo)
					break
				case fieldTypeTime:
					createDatePickerInput(fieldInfo)
					break
				case fieldTypeBool:
					if(colType==='checkbox') {
						createCheckBoxInput(fieldInfo)
					} else if (colType==="toggle") {
						createToggle(fieldInfo)
					} else {
						console.log("Unknown column type for boolean field : " + colType)
					}
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
				$colTypeSelection.append(selectOptionHTML('progress','Progress Indicator'))
				break
			case fieldTypeText:
				$colTypeSelection.append(selectOptionHTML('textInput','Text input'))
				break
			case fieldTypeComment:
				$colTypeSelection.append(selectOptionHTML('comment','Comment box'))
				break
			case fieldTypeAttachment:
				$colTypeSelection.append(selectOptionHTML('attachment','Attachments'))
				break
			case fieldTypeLongText:
				$colTypeSelection.append(selectOptionHTML('noteEditor','Note editor'))
				break
			case fieldTypeTime:
				$colTypeSelection.append(selectOptionHTML('datePicker','Date picker'))
				break
			case fieldTypeUser:
				$colTypeSelection.append(selectOptionHTML('userSelection','User selection'))
				$colTypeSelection.append(selectOptionHTML('socialButton','Social button'))
				break
			case fieldTypeBool:
				$colTypeSelection.append(selectOptionHTML('checkbox','Checkbox'))
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