

function loadRecordIntoButton(buttonElem, recordRef) {	
	// no-op
}

function initFormButtonRecordEditBehavior(componentContext,getRecordFunc, updateRecordFunc,buttonObjectRef,
		loadFormViewComponentFunc,loadRecordIntoFormLayoutFunc) {
	
	
	var $popupFormDialog = $('#formButtonPopupFormDialog')
	
	var buttonElemID = buttonIDFromContainerElemID(buttonObjectRef.buttonID)
	initButtonClickHandler('#' + buttonElemID, function() {
		console.log("Form button clicked: " + JSON.stringify(buttonObjectRef))
		
		var currRecord = getRecordFunc()

		var canvasSelector = '#formButtonPopupFormCanvas'
		var $viewFormCanvas = $(canvasSelector)
		$viewFormCanvas.empty()

		
		function getPopupFormRecordFunc() { return currRecord }
		function updatePopupFormRecordFunc(updatedRecordRef) {
			loadRecordIntoFormLayoutFunc(canvasSelector,updatedRecordRef)
			
			// Propagate the record update the parent form, allowing an update
			// to the parent layout.
			updateRecordFunc(updatedRecordRef)
		}
		
		function showDialogAfterFormComponentLoaded() {
			loadRecordIntoFormLayoutFunc(canvasSelector,currRecord)
			$popupFormDialog.modal('show')
		}
			
		var viewFormContext = {
			databaseID: componentContext.databaseID,
			formID: buttonObjectRef.properties.linkedFormID
		}
				
		var popupMode = buttonObjectRef.properties.popupBehavior.popupMode
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
					})
					$popupFormDialog.modal('hide')
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
			
			var changeSetID = ""
			loadFormViewComponentFunc(canvasSelector, viewFormContext, changeSetID,
				 getPopupFormRecordFunc, updatePopupFormRecordFunc,
				 showDialogAfterFormComponentLoaded)
		}
			
		
	})
	
	
	var $buttonContainer = $('#'+buttonObjectRef.buttonID)
	$buttonContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoButton
	})
	
}
