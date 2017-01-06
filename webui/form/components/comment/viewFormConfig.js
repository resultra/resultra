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
	var addCommentButtonID = commentAddCommentButtonIDFromContainerElemID(commentContainerID)
	
	console.log("initCommentBoxRecordEditBehavior: container ID =  " +commentContainerID + ' comment box ID = '+ commentID)
	
	var $commentContainer = $('#'+commentContainerID)
	
	$commentContainer.data("componentContext",componentContext)
	
	$commentContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoCommentBox
	})
	
	var commentFieldID = commentObjectRef.properties.fieldID
	
	var $addCommentButton = $('#' + addCommentButtonID)
	
	
	var $commentInput = $('#'+commentInputID)
	initButtonControlClickHandler($addCommentButton,function() {
		var commentVal = $commentInput.val()
		var currRecordRef = currRecordSet.currRecordRef()
				
		if(nonEmptyStringVal(commentVal)) {
			console.log("initCommentBoxRecordEditBehavior: Add comment:" + commentVal)
			
			
			var commentValueFormat = {
				context:"commentBox",
				format:"general"
			}
			
			var setRecordValParams = { 
				parentDatabaseID:viewListContext.databaseID,
				recordID:currRecordRef.recordID, 
				fieldID:commentFieldID, 
				value:commentVal,
			valueFormat:commentValueFormat }
		
			console.log("Setting date value: " + JSON.stringify(setRecordValParams))
		
			jsonAPIRequest("recordUpdate/setCommentFieldValue",setRecordValParams,function(updatedRecordRef) {
			
				// After updating the record, the local cache of records in currentRecordSet will
				// be out of date. So after updating the record on the server, the locally cached
				// version of the record also needs to be updated.
				currRecordSet.updateRecordRef(updatedRecordRef)
				// After changing the value, some of the calculated fields may have changed. For this
				// reason, it is necessary to reload the record into the layout/form, so the most
				// up to date values will be displayed.
				loadCurrRecordIntoLayout()
			}) // set record's text field value
						
				
		}
		$commentInput.val("")
	})
	
	// TODO - Initialize edit behavior
		
}