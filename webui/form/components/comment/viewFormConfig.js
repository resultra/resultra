function loadRecordIntoCommentBox(commentElem, recordRef) {
	
	console.log("loadRecordIntoTextBox: loading record into text box: " + JSON.stringify(recordRef))
	
	var commentObjectRef = commentElem.data("objectRef")
	var $commentInput = commentElem.find('input')
	var componentContext = commentElem.data("componentContext")
	
	var fieldID = commentObjectRef.properties.fieldID
	
	// TODO - Load existing comments into comment box
	
	
}


function initCommentBoxRecordEditBehavior(componentContext,commentObjectRef) {
	
	var commentContainerID = commentObjectRef.commentID
	var commentID = commentElemIDFromContainerElemID(commentContainerID)
	
	var commentInputID = commentInputIDFromContainerElemID(commentContainerID)
	
	
	
	
	console.log("initCommentBoxRecordEditBehavior: container ID =  " +commentContainerID + ' comment box ID = '+ commentID)
	
	var $commentContainer = $('#'+commentContainerID)
	
	$commentContainer.data("componentContext",componentContext)
	

	$commentContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoCommentBox
	})
	
	var fieldID = commentObjectRef.properties.fieldID
	
	// TODO - Initialize edit behavior
		
}