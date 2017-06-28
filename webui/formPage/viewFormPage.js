$(document).ready(function() {	
	 
	initSubmitFormUILayoutPanes()
				
	initUserDropdownMenu()
	
	function initRecordFormView(pageConfig,recordRef,changeSetID) {
		
		var $formViewCanvas = $('#viewFormPageLayoutCanvas')
		$formViewCanvas.empty()
	
		var currRecord = recordRef
		function getFormRecordFunc() { return currRecord }
		function updateFormRecordFunc(updatedRecordRef) {
			currRecord = updatedRecordRef
			loadRecordIntoFormLayout($formViewCanvas,updatedRecordRef)
		}
	
		var viewFormContext = {
			databaseID: viewFormPageContext.databaseID,
			formID: viewFormPageContext.formID
		}
		
		function loadRecordIntoFormViewAfterFormComponentsLoaded() {
			
			function loadRecordWithDefaultVals(defaultVals) {
				if (defaultVals.length > 0) {
					// Apply the default values before loading the form.
					var defaultValRecord = getFormRecordFunc()
					var defaultValParams = {
						parentDatabaseID: componentContext.databaseID,
						recordID: defaultValRecord.recordID,
						changeSetID: changeSetID,
						defaultVals: defaultVals }
					jsonAPIRequest("recordUpdate/setDefaultValues",defaultValParams,function(updatedRecordRef) {
			
						// Update the current record to include default values
						currRecord = updatedRecordRef
											 
						loadRecordIntoFormLayout($formViewCanvas,recordRef)
					})			
				} else {
					// load without default values
					loadRecordIntoFormLayout($formViewCanvas,recordRef)
				}
			}
			loadRecordWithDefaultVals(pageConfig.defaultVals)
			
			var $saveButton = $("#viewFormPageSaveButton")
			if(pageConfig.saveMode === FormViewModeSave) {
				$saveButton.show()
				initButtonControlClickHandler($saveButton, function() {					
					var commitChangeParams = {
						recordID: getFormRecordFunc().recordID,
						changeSetID: changeSetID }
					jsonAPIRequest("recordUpdate/commitChangeSet",commitChangeParams,function(updatedRecordRef) {
						// If the popup form is modal, the parent form's record is not updated until the "Save Changes" button
						// is pressed.
						updateFormRecordFunc(updatedRecordRef)
					})
				})
			} else {
				$saveButton.hide()
			}
			
	
		}

		var recordProxy = {
			changeSetID: changeSetID,
			getRecordFunc: getFormRecordFunc,
			updateRecordFunc: updateFormRecordFunc
		}
		loadFormViewComponents($formViewCanvas,viewFormContext,recordProxy,
				loadRecordIntoFormViewAfterFormComponentsLoaded)
				
		
	}
	
	var FormViewModeModeless = "modeless"
	var FormViewModeSave = "modal"
	
	function getRecordRefAndChangeSetID(pageConfig,doneCallback) {
		
		var recordRef
		var changeSetID		
		var callsRemaining = 2
		function processOneCall() {
			callsRemaining--
			if (callsRemaining <= 0) {
				doneCallback(pageConfig,recordRef,changeSetID)
			}
		}
		
		var getRecordParams = {
			parentDatabaseID: viewFormPageContext.databaseID,
			recordID: viewFormPageContext.recordID
		}
		jsonAPIRequest("recordRead/getRecordValueResults",getRecordParams,function(existingRecordRef) {
			recordRef = existingRecordRef
			processOneCall()
		})
		
		if (pageConfig.saveMode === FormViewModeSave) {					
			jsonAPIRequest("record/allocateChangeSetID",{},function(changeSetIDResp) {
				changeSetID = changeSetIDResp.changeSetID
				processOneCall()
			})
		} else {
			var immediatelyCommitChangesChangeSetID = ""
			changeSetID = immediatelyCommitChangesChangeSetID
			$("#viewFormPageSaveButton").hide()
			
			processOneCall()
		}
	}
	
	function getPageConfig(pageConfigDoneCallback) {
		
		var pageConfig = {}
		
		if(viewFormPageContext.srcColumnID.length > 0) {
			// Get the default values from the column used to open the form. 
			var getButtonParams = {
				buttonID: viewFormPageContext.srcColumnID
			}
			jsonAPIRequest("tableView/formButton/getFromButtonID",getButtonParams,function(buttonRef) {
				pageConfig.defaultVals = buttonRef.properties.defaultValues
				pageConfig.saveMode = buttonRef.properties.popupBehavior.popupMode
				pageConfigDoneCallback(pageConfig)	
			})
		} else if(viewFormPageContext.srcFrmButtonID.length > 0) {
			var getButtonParams = {
				buttonID: viewFormPageContext.srcFrmButtonID
			}
			jsonAPIRequest("frm/formButton/get",getButtonParams,function(buttonRef) {
				pageConfig.defaultVals = buttonRef.properties.defaultValues
				pageConfig.saveMode = buttonRef.properties.popupBehavior.popupMode
				pageConfigDoneCallback(pageConfig)	
			})
				
		} else {
			// Load without default values.
			pageConfig.defaultVals = []
			pageConfig.saveMode = FormViewModeSave
			pageConfigDoneCallback(pageConfig)	
		}
		
	}
	
	getPageConfig(function(pageConfig) {
		getRecordRefAndChangeSetID(pageConfig,initRecordFormView)
	})
	
					
}); // document ready