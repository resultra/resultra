$(document).ready(function() {	
	 
	initSubmitFormUILayoutPanes()
				
	initUserDropdownMenu()
	
	function initRecordFormView(recordRef) {
		
		var $formViewCanvas = $('#viewFormPageLayoutCanvas')
		$formViewCanvas.empty()
		
		var currRecord = recordRef
		function getFormRecordFunc() { return currRecord }
		function updateFormRecordFunc(updatedRecordRef) {
			currRecord = updatedRecordRef
			loadRecordIntoFormLayout($formViewCanvas,updatedRecordRef)
		}
		function loadRecordIntoFormViewAfterFormComponentsLoaded() {
			loadRecordIntoFormLayout($formViewCanvas,recordRef)
		}		
		var viewFormContext = {
			databaseID: viewFormPageContext.databaseID,
			formID: viewFormPageContext.formID
		}
		
		var FormViewModeModeless = "modeless"
		var FormViewModeSave = "modal"
		var formViewMode = FormViewModeModeless
				
		if (formViewMode === FormViewModeSave) {			
			
			jsonAPIRequest("record/allocateChangeSetID",{},function(changeSetIDResp) {
				
				var recordProxy = {
					changeSetID: changeSetIDResp.changeSetID,
					getRecordFunc: getFormRecordFunc,
					updateRecordFunc: updateFormRecordFunc
				}
				loadFormViewComponents($formViewCanvas,viewFormContext,recordProxy,
						loadRecordIntoFormViewAfterFormComponentsLoaded)
	
	/*			
				var defaultVals = buttonObjectRef.properties.popupBehavior.defaultValues
				
				if (defaultVals.length > 0) {
					// Apply the default values before loading the form.
					var defaultValRecord = getPopupFormRecordFunc()
					var defaultValParams = {
						parentDatabaseID: componentContext.databaseID,
						recordID: defaultValRecord.recordID,
						changeSetID: changeSetIDResp.changeSetID,
						defaultVals: defaultVals }
					jsonAPIRequest("recordUpdate/setDefaultValues",defaultValParams,function(updatedRecordRef) {
						
						currRecord = updatedRecordRef
						
						// loadFormViewComponentFunc is passed in as a parameter, since the loadFormViewComponents 
						// function is in 
						// a package which has a dependency on this package.
									 
						loadFormViewComponentFunc($popupFormViewCanvas, viewFormContext, recordProxy,
							 showDialogAfterFormComponentLoaded)
					})
					
					
				} else {
					loadFormViewComponentFunc($popupFormViewCanvas, viewFormContext, recordProxy,
						 showDialogAfterFormComponentLoaded)
				}
	   */		
					
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
			
			})
			
		} else { // Form shown in modeless mode
			
			
			$("#viewFormPageSaveButton").hide()
			
			var immediatelyCommitChangesChangeSetID = ""
			var recordProxy = {
				changeSetID: immediatelyCommitChangesChangeSetID,
				getRecordFunc: getFormRecordFunc,
				updateRecordFunc: updateFormRecordFunc
			}
	
			loadFormViewComponents($formViewCanvas,viewFormContext,recordProxy,
					loadRecordIntoFormViewAfterFormComponentsLoaded)
			
		}
		
	}
	
	var getRecordParams = {
		parentDatabaseID: viewFormPageContext.databaseID,
		recordID: viewFormPageContext.recordID
	}
	jsonAPIRequest("recordRead/getRecordValueResults",getRecordParams,function(recordRef) {
		initRecordFormView(recordRef)
	})
	
/*	var submitFormParams = {
		databaseID: submitFormPageContext.databaseID,
		$parentFormCanvas: $('#submitFormPageLayoutCanvas'),
		formLinkID: submitFormPageContext.formLinkID,
		formID: submitFormPageContext.formID
	}
	
	var $addAnotherButton = $('#newItemPageAddAnotherButton')
	initButtonControlClickHandler($addAnotherButton, function() {
		initFormPageSubmitForm(submitFormParams)
	})
	
	
	initFormPageSubmitForm(submitFormParams)
	*/
					
}); // document ready