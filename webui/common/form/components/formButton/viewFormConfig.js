

function loadRecordIntoButton(buttonElem, recordRef) {	
	// no-op
}

function initFormButtonRecordEditBehavior(componentContext,
			parentFormGetRecordFunc, parentFormUpdateRecordFunc,buttonObjectRef,
		loadFormViewComponentFunc,loadRecordIntoFormLayoutFunc) {
	
	
	var $popupFormDialog = $('#formButtonPopupFormDialog')
			
	var popupMode = buttonObjectRef.properties.popupBehavior.popupMode
	
	var buttonElemID = buttonIDFromContainerElemID(buttonObjectRef.buttonID)
	initButtonClickHandler('#' + buttonElemID, function() {
		console.log("Form button clicked: " + JSON.stringify(buttonObjectRef))
		
		// Editing of the record in the popup is done with the parent form's current record.
		var currRecord = parentFormGetRecordFunc()

		var canvasSelector = '#formButtonPopupFormCanvas'
		var $viewFormCanvas = $(canvasSelector)
		$viewFormCanvas.empty()

		
		function getPopupFormRecordFunc() { return currRecord }
		function updatePopupFormRecordFunc(updatedRecordRef) {
			loadRecordIntoFormLayoutFunc(canvasSelector,updatedRecordRef)
			
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
			loadRecordIntoFormLayoutFunc(canvasSelector,currRecord)
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
				loadFormViewComponentFunc(canvasSelector, viewFormContext, changeSetIDResp.changeSetID,
					 getPopupFormRecordFunc, updatePopupFormRecordFunc,
					 showDialogAfterFormComponentLoaded)
					
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
	
	
	var $buttonContainer = $('#'+buttonObjectRef.buttonID)
	$buttonContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoButton
	})
	
}
