function clearNewCommentAttachmentList($commentContainer) {
	$commentContainer.data("attachmentList",[])
	var $attachList = commentAttachmentListFromContainer($commentContainer)
	$attachList.empty()
}

function getNewCommentAttachmentList($commentContainer) {
	return $commentContainer.data("attachmentList")
}

function setNewCommentAttachmentList($commentContainer,attachmentList) {
	$commentContainer.data("attachmentList",attachmentList)
	
	var $attachList = commentAttachmentListFromContainer($commentContainer)
	
	populateAttachmentList($attachList,attachmentList)
}


function loadRecordIntoCommentBox(commentElem, recordRef) {
	
	console.log("loadRecordIntoCommentBox: loading record into comment box: " + JSON.stringify(recordRef))
	
	var commentObjectRef = commentElem.data("objectRef")
	var componentContext = commentElem.data("componentContext")
	
	// Clear any previous comments - TODO: is there any way to preserve the comments
	// if the user switches items.
	clearNewCommentAttachmentList(commentElem)

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
					'<div class="formTimelineComment">' + escapeHTML(valChange.updatedValue.commentText) + '</div>' +
					'<div class="formTimelineCommentAttachments"></div>' + 
				'</div>';
				
				var $commentContainer = $(commentHTML)
				var $attachments = $commentContainer.find(".formTimelineCommentAttachments")
				populateAttachmentList($attachments,valChange.updatedValue.attachments)
		
				return $commentContainer
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
	var $commentList = commentCommentListFromContainer($commentContainer)
	
	var $commentInput = commentInputFromContainer($commentContainer)
	var $commentEntryControls = commentEntryControlsFromContainer($commentContainer)
	var $entryContainer = commentEntryContainerFromOverallCommentContainer($commentContainer)
	
	// Prevent selection of the form component when clicking on the list or 
	// comment entry controls.
	$entryContainer.click(function (e) {
		e.stopPropagation()
	})
	$commentList.click(function (e) {
		e.stopPropagation()
	})
	
	
	// Dynamically resize the comment input and comment list when the user starts and 
	// finishes entering comments.
	function resizeCommentListToEntryControls() {	
		var entryBottom = $entryContainer.position().top + $entryContainer.outerHeight(true);
		var listTopPx = (entryBottom+5) + "px"
		$commentList.css('top',listTopPx)
	}
	function resetAndMinimizeCommentEntry() {
		$commentInput.val("")
		$commentInput.height(20)
		$commentEntryControls.hide()
		resizeCommentListToEntryControls()
	}
	function expandCommentEntryAndControls() {
		$commentInput.height(60)
		$commentEntryControls.show()
		resizeCommentListToEntryControls()
	}
	
	initButtonControlClickHandler($addCommentButton,function() {
		var commentVal = $commentInput.val()
		
		var newCommentAttachments = getNewCommentAttachmentList($commentContainer)
		
		var currRecordRef = recordProxy.getRecordFunc()
				
		if(nonEmptyStringVal(commentVal) || (attachments.length > 0)) {
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
				commentText:commentVal,
				attachments:newCommentAttachments, 
				valueFormat:commentValueFormat }
		
			console.log("Setting date value: " + JSON.stringify(setRecordValParams))
		
			jsonAPIRequest("recordUpdate/setCommentFieldValue",setRecordValParams,function(updatedRecordRef) {
			
				clearNewCommentAttachmentList($commentContainer)
				
				// After updating the record, the local cache of records in currentRecordSet will
				// be out of date. So after updating the record on the server, the locally cached
				// version of the record also needs to be updated.
				recordProxy.updateRecordFunc(updatedRecordRef)
			}) // set record's text field value
						
				
		}
		resetAndMinimizeCommentEntry()
	})
	
	var $attachmentButton = commentAttachmentButtonFromContainer($commentContainer)
	
	$commentContainer.data("attachmentList",[])
	
	initButtonControlClickHandler($attachmentButton,function() {
				
		function addAttachmentsToAttachmentList(newAttachments) {
			var currAttachments = getNewCommentAttachmentList($commentContainer)
			currAttachments = $.merge(currAttachments,newAttachments)
			setNewCommentAttachmentList($commentContainer,currAttachments)
		}
		
		var manageAttachmentParams = {
			parentDatabaseID: componentContext.databaseID,
			addAttachmentsCallback: addAttachmentsToAttachmentList
		}
		openAddAttachmentsDialog(manageAttachmentParams)
	})
	
	// Resize the text area dynamically when the user starts and finishes editing.
	$commentInput.focus(function() {
		expandCommentEntryAndControls()
	})
	$commentInput.blur(function() {
		var commentInputVal = $commentInput.val()
		if (commentInputVal.length === 0) {
			resetAndMinimizeCommentEntry()
		}
	})
	// If the user resizes the entry box, dynamically resize the comment list with it.
	$commentInput.mousemove(function() {
		resizeCommentListToEntryControls()
	})
	
		
}