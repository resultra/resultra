

function loadRecordIntoButton(buttonElem, recordRef) {
	// no-op
}

function initButtonRecordEditBehavior(componentContext,buttonObjectRef) {	
	var $buttonContainer = $('#'+buttonObjectRef.buttonID)
	$headerContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoHeader
	})
	
}
