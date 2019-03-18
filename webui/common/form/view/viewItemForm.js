// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
var FormViewModeModeless = "modeless"
var FormViewModeSave = "modal"


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
		databaseID: pageConfig.databaseID,
		formID: pageConfig.formID
	}
	
	function loadRecordIntoFormViewAfterFormComponentsLoaded() {
		
		function loadRecordWithDefaultVals(defaultVals) {
			if (defaultVals.length > 0) {
				// Apply the default values before loading the form.
				var defaultValRecord = getFormRecordFunc()
				var defaultValParams = {
					parentDatabaseID: pageConfig.databaseID,
					recordID: pageConfig.recordID,
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
		var $saveAndCloseButton = $('#viewFormPageSaveCloseButton')
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
			
			$saveAndCloseButton.show()
			initButtonControlClickHandler($saveAndCloseButton, function() {					
				var commitChangeParams = {
					recordID: getFormRecordFunc().recordID,
					changeSetID: changeSetID }
				jsonAPIRequest("recordUpdate/commitChangeSet",commitChangeParams,function(updatedRecordRef) {
					// If the popup form is modal, the parent form's record is not updated until the "Save Changes" button
					// is pressed.
				//	updateFormRecordFunc(updatedRecordRef)
					window.close()
				})
			})
			
		} else {
			$saveButton.hide()
			$saveAndCloseButton.hide()
		}
		
		// If a call-back is passed in for going back to the previous view, then call it when the back
		// button is pressed.
		var $backButton = $("#viewFormPageBackButton")
		if(pageConfig.loadLastViewCallback !== null) {
			$backButton.show()
			initButtonControlClickHandler($backButton, function() {
				pageConfig.loadLastViewCallback()					
			})
		} else {
			$backButton.hide()
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
		parentDatabaseID: pageConfig.databaseID,
		recordID: pageConfig.recordID
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
