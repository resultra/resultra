

function loadRecordIntoButton(buttonElem, recordRef) {
	
	var buttonObjectRef = buttonElem.data("objectRef")
	
	var buttonElemID = buttonIDFromContainerElemID(buttonObjectRef.buttonID)
	
	initButtonClickHandler('#' + buttonElemID, function() {
		console.log("Form button clicked: " + JSON.stringify(buttonObjectRef))		
	})
	// no-op
}

function initFormButtonRecordEditBehavior(componentContext,getRecordFunc, updateRecordFunc,buttonObjectRef) {	
	var $buttonContainer = $('#'+buttonObjectRef.buttonID)
	$buttonContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoButton
	})
	
}
