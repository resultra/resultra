

function loadRecordIntoButton(buttonElem, recordRef) {	
	// no-op
}

function initFormButtonRecordEditBehavior($buttonContainer,componentContext,
			parentFormGetRecordFunc, parentFormUpdateRecordFunc,buttonObjectRef,
		loadFormViewComponentFunc,loadRecordIntoFormLayoutFunc) {
	
	
	$buttonContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoButton
	})
	
	
	var $popupFormDialog = $('#formButtonPopupFormDialog')
			
	var popupMode = buttonObjectRef.properties.popupBehavior.popupMode
	
	var $formButton = buttonFromFormButtonContainer($buttonContainer)
			
	initButtonControlClickHandler($formButton, function() {
		console.log("Form button clicked: " + JSON.stringify(buttonObjectRef))
		
		// Editing of the record in the popup is done with the parent form's current record.
		var currRecord = parentFormGetRecordFunc()

		var canvasSelector = '#formButtonPopupFormCanvas'
		var $viewFormCanvas = $(canvasSelector)
		$viewFormCanvas.empty()

		
		function getPopupFormRecordFunc() { return currRecord }
		function updatePopupFormRecordFunc(updatedRecordRef) {
			
			var $parentFormLayout = $(canvasSelector)
			loadRecordIntoFormLayoutFunc($parentFormLayout,updatedRecordRef)
			
			// Propagate the record update the parent form, allowing an update
			// to the parent layout. The changes are only propagated when the popup
			// form is not in a modal state and the changes are committed immediately.
			// If the popup is in a modal state, the changes are linked to a temporary
			// changeSetID and are not propagated until the "Save Changes" button is
			// pressed; however, if the "Cancel" button is pressed all these changes are
			// rolled back.
			if (popupMode !== FormButtonPopupBehaviorModal) {
				parentFormUpdateRecordFunc(updatedRecordRef)			
			}
		}
		
		function showDialogAfterFormComponentLoaded() {
			var $parentFormLayout = $(canvasSelector)
			loadRecordIntoFormLayoutFunc($parentFormLayout ,currRecord)
			$popupFormDialog.modal('show')
		}
			
		var viewFormContext = {
			databaseID: componentContext.databaseID,
			formID: buttonObjectRef.properties.linkedFormID
		}
				
		if (popupMode === FormButtonPopupBehaviorModal) {
			
			$(".formButtonPopupModalPopupDialogButton").show()
			$(".formButtonPopupModelessPopupDialogButton").hide()
			
			
			jsonAPIRequest("record/allocateChangeSetID",{},function(changeSetIDResp) {
				
				// loadFormViewComponentFunc is passed in as a parameter, since the loadFormViewComponents function is in 
				// a package which has a dependency on this package. 
				
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
						
						loadFormViewComponentFunc(canvasSelector, viewFormContext, changeSetIDResp.changeSetID,
							 getPopupFormRecordFunc, updatePopupFormRecordFunc,
							 showDialogAfterFormComponentLoaded)
					})
					
					
				} else {
					loadFormViewComponentFunc(canvasSelector, viewFormContext, changeSetIDResp.changeSetID,
						 getPopupFormRecordFunc, updatePopupFormRecordFunc,
						 showDialogAfterFormComponentLoaded)
				}
				
					
				initButtonClickHandler('#formButtonPopupFormDialogSaveChangesButton', function() {
					console.log("Modal Save changes button clicked: " + JSON.stringify(buttonObjectRef))
					// TODO - Remove the temporary changes set ID for any changes made while editing the record.
					
					var commitChangeParams = {
						recordID: getPopupFormRecordFunc().recordID,
						changeSetID: changeSetIDResp.changeSetID }
					jsonAPIRequest("recordUpdate/commitChangeSet",commitChangeParams,function(updatedRecordRef) {
						// If the popup form is modal, the parent form's record is not updated until the "Save Changes" button
						// is pressed.
						parentFormUpdateRecordFunc(updatedRecordRef)
						$popupFormDialog.modal('hide')
					})
				})
				initButtonClickHandler('#formButtonPopupFormDialogCancelChangesButton', function() {
					console.log("Cancel button clicked: " + JSON.stringify(buttonObjectRef))
					$popupFormDialog.modal('hide')
				})
				
			
			})
			
		} else { // Popup shown in modeless mode
			
			// TBD - Should the default value be set when the popup form is modeless? It wouldn't
			// seem to make sense to set the default values in this case.
			
			$(".formButtonPopupModalPopupDialogButton").hide()
			$(".formButtonPopupModelessPopupDialogButton").show()
			initButtonClickHandler('#formButtonPopupFormDialogDoneButton', function() {
				console.log("Form button clicked: " + JSON.stringify(buttonObjectRef))
				$popupFormDialog.modal('hide')
			})
			
			var immediatelyCommitChangesChangeSetID = ""
			loadFormViewComponentFunc(canvasSelector, viewFormContext, immediatelyCommitChangesChangeSetID,
				 getPopupFormRecordFunc, updatePopupFormRecordFunc,
				 showDialogAfterFormComponentLoaded)
		}
			
		
	})
	
	
	
}
