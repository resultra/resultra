

function loadRecordIntoHeader(headerElem, recordRef) {
	// no-op
}

function initHeaderRecordEditBehavior(componentContext,headerObjectRef) {	
	var $headerContainer = $('#'+headerObjectRef.headerID)
	$headerContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoHeader
	})
	
}
