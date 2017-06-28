$(document).ready(function() {	
	 
	initSubmitFormUILayoutPanes()
				
	initUserDropdownMenu()
	
	function initRecordFormView(recordRef,changeSetID) {
		
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
			
			if(viewFormPageContext.srcColumnID.length > 0) {
				// Get the default values from the column used to open the form. 
				var getButtonParams = {
					buttonID: viewFormPageContext.srcColumnID
				}
				jsonAPIRequest("tableView/formButton/getFromButtonID",getButtonParams,function(buttonRef) {
					var defaultVals = buttonRef.properties.defaultValues
					loadRecordWithDefaultVals(defaultVals)
				})
			} else if(viewFormPageContext.srcFrmButtonID.length > 0) {
				var getButtonParams = {
					buttonID: viewFormPageContext.srcFrmButtonID
				}
				jsonAPIRequest("frm/formButton/get",getButtonParams,function(buttonRef) {
					var defaultVals = buttonRef.properties.defaultValues
					loadRecordWithDefaultVals(defaultVals)
				})
					
			} else {
				// Load without default values.
				loadRecordIntoFormLayout($formViewCanvas,recordRef)		
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
	var formViewMode = FormViewModeModeless
	
	function getRecordRefAndChangeSetID(doneCallback) {
		var recordRef
		var changeSetID		
		var callsRemaining = 2
		function processOneCall() {
			callsRemaining--
			if (callsRemaining <= 0) {
				doneCallback(recordRef,changeSetID)
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
		
		if (formViewMode === FormViewModeSave) {					
			jsonAPIRequest("record/allocateChangeSetID",{},function(changeSetIDResp) {
				changeSetID = changeSetIDResp.changeSetID
				
				initButtonClickHandler('#viewFormPageSaveButton', function() {
					console.log("Modal Save changes button clicked: " + JSON.stringify(buttonObjectRef))
					// TODO - Remove the temporary changes set ID for any changes made while editing the record.
					
					var commitChangeParams = {
						recordID: getFormRecordFunc().recordID,
						changeSetID: changeSetIDResp.changeSetID }
					jsonAPIRequest("recordUpdate/commitChangeSet",commitChangeParams,function(updatedRecordRef) {
						// If the popup form is modal, the parent form's record is not updated until the "Save Changes" button
						// is pressed.
						console.log("Form changes saved")
					})
				})
				processOneCall()
			})
		} else {
			var immediatelyCommitChangesChangeSetID = ""
			changeSetID = immediatelyCommitChangesChangeSetID
			$("#viewFormPageSaveButton").hide()
			
			processOneCall()
		}
	}
	
	getRecordRefAndChangeSetID(initRecordFormView)
					
}); // document ready