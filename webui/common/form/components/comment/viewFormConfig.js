function loadRecordIntoCommentBox(commentElem, recordRef) {
	
	console.log("loadRecordIntoCommentBox: loading record into comment box: " + JSON.stringify(recordRef))
	
	var commentObjectRef = commentElem.data("objectRef")
	var componentContext = commentElem.data("componentContext")

	var $commentInput = commentInputFromContainer(commentElem)
		
	var changeInfoParams = {
		recordID: recordRef.recordID,
		fieldID: commentObjectRef.properties.fieldID
	}
	
	jsonAPIRequest("record/getFieldValChangeInfo",changeInfoParams,function(valChanges) {
		
		console.log("loadRecordIntoCommentBox: retrieved comment info: " + JSON.stringify(valChanges))
		
		var $commentList = commentCommentListFromContainer(commentElem)
		$commentList.empty()
		
		for(var valChangeIter = 0; valChangeIter < valChanges.length; valChangeIter++) {
			
			function createOneCommentValDisplay(valChange) {
		
				var formattedUserName = "@" + valChange.userName
				if(valChange.isCurrentUser) {
						formattedUserName = formattedUserName + ' (you)'
				}
		
				var formattedCreateDate = moment(valChange.updateTime).calendar()

				var commentHTML =  '<div class="list-group-item">' +
					'<div><small>' + formattedUserName  + ' - ' + formattedCreateDate + '</small></div>' +
					'<div class="formTimelineComment">' + escapeHTML(valChange.updatedValue) + '</div>' +
				'</div>';		
		
				return $(commentHTML)
			}

			var valChange = valChanges[valChangeIter]
			
			$commentList.append(createOneCommentValDisplay(valChange))
			
		}
		
	}) // set record's text field value
	
	
}


function initCommentBoxRecordEditBehavior($commentContainer, componentContext,recordProxy, commentObjectRef) {
				
	$commentContainer.data("componentContext",componentContext)
	
	$commentContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoCommentBox
	})
	
	var commentFieldID = commentObjectRef.properties.fieldID
	
	var $addCommentButton = commentAddCommentButtonFromContainer($commentContainer)
	var $commentInput = commentInputFromContainer($commentContainer)
	
	
	initButtonControlClickHandler($addCommentButton,function() {
		var commentVal = $commentInput.val()
		
		var currRecordRef = recordProxy.getRecordFunc()
				
		if(nonEmptyStringVal(commentVal)) {
			console.log("initCommentBoxRecordEditBehavior: Add comment:" + commentVal)
			
			
			var commentValueFormat = {
				context:"commentBox",
				format:"general"
			}
			
			var setRecordValParams = { 
				parentDatabaseID:currRecordRef.parentDatabaseID,
				recordID:currRecordRef.recordID,
				changeSetID: recordProxy.changeSetID,
				fieldID:commentFieldID, 
				value:commentVal,
			valueFormat:commentValueFormat }
		
			console.log("Setting date value: " + JSON.stringify(setRecordValParams))
		
			jsonAPIRequest("recordUpdate/setCommentFieldValue",setRecordValParams,function(updatedRecordRef) {
			
				// After updating the record, the local cache of records in currentRecordSet will
				// be out of date. So after updating the record on the server, the locally cached
				// version of the record also needs to be updated.
				recordProxy.updateRecordFunc(updatedRecordRef)
			}) // set record's text field value
						
				
		}
		$commentInput.val("")
	})
	
	// TODO - Initialize edit behavior
		
}