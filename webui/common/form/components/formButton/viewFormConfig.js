

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
				
		// loadFormViewComponentFunc is passed in as a parameter, since the loadFormViewComponents function is in 
		// a package which has a dependency on this package. 
		loadFormViewComponentFunc(canvasSelector, viewFormContext, 
			 getPopupFormRecordFunc, updatePopupFormRecordFunc,
			 showDialogAfterFormComponentLoaded)
		
	})
	
	initButtonClickHandler('#formButtonPopupFormDialogDoneButton', function() {
		console.log("Form button clicked: " + JSON.stringify(buttonObjectRef))
		$popupFormDialog.modal('hide')
	})
	
	var $buttonContainer = $('#'+buttonObjectRef.buttonID)
	$buttonContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoButton
	})
	
}
