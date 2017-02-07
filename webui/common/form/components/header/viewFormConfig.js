

function loadRecordIntoHeader(headerElem, recordRef) {
	// no-op
}

function initHeaderRecordEditBehavior($headerContainer,componentContext,headerObjectRef) {	
	$headerContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoHeader
	})
	
}
