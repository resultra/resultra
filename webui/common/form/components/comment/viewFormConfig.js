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

function resizeEntryAreaForCommentComponentPermissions($commentContainer,commentObjRef) {
	var $entryContainer = commentEntryContainerFromOverallCommentContainer($commentContainer)
	if(formComponentIsReadOnly(commentObjRef.properties.permissions)) {
		$entryContainer.hide()
	} else {
		$entryContainer.show()
		
	}
	// Set the maximum height of the comment area to be the remainder of the comment components
	var entryBottom = $entryContainer.position().top + $entryContainer.outerHeight(true);
	var listMaxHeightPx = (commentObjRef.properties.geometry.sizeHeight - (entryBottom+5)) + "px"
	var $commentList = commentCommentListFromContainer($commentContainer)
	$commentList.css('max-height',listMaxHeightPx)
	
}


function loadRecordIntoCommentBox(commentElem, recordRef) {
	
	console.log("loadRecordIntoCommentBox: loading record into comment box: " + JSON.stringify(recordRef))
	
	var commentObjectRef = commentElem.data("objectRef")
	var componentContext = commentElem.data("componentContext")
	
	resizeEntryAreaForCommentComponentPermissions(commentElem,commentObjectRef)
	
	
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

				var commentDisplayHTML = formatInlineContentHTMLDisplay(valChange.updatedValue.commentText)
				
				var commentHTML =  '<div class="list-group-item">' +
					'<div><small>' + formattedUserName  + ' - ' + formattedCreateDate + '</small></div>' +
					'<div class="inlineContent">' + commentDisplayHTML + '</div>' +
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
	
	// In view mode, the height will be flexible. However, by placing a maximum width
	// on the comment list (see below), the overall height set in the form designer
	// will be preserved when the user is not editing a comment.
	setElemFixedWidthFlexibleHeight($commentContainer,commentObjectRef.properties.geometry.sizeWidth)
	
	resizeEntryAreaForCommentComponentPermissions($commentContainer,commentObjectRef)

	function resetAndMinimizeCommentEntry() {
		$commentInput.html('<p class="commentPlaceholder">Enter a comment...</p>')
		$commentInput.height(40)
		$commentEntryControls.hide()
	}
	function expandCommentEntryAndControls() {
		$commentInput.find(".commentPlaceholder").remove()
		$commentInput.height(80)
		$commentEntryControls.show()
	}
	
	
	
	$commentContainer.data("attachmentList",[])
	
	var $attachmentButton = commentAttachmentButtonFromContainer($commentContainer)
	initButtonControlClickHandler($attachmentButton,function() {
		
		console.log("Attachment button clicked for comments")
				
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
	
	$commentInput.click(function() {
		if (!inlineCKEditorEnabled($commentInput)) {
	
			expandCommentEntryAndControls()
			
			var editor = enableInlineCKEditor($commentInput)
			
			$commentInput.focus()
				
			initButtonControlClickHandler($addCommentButton,function() {
		
				// Save the comment
				var commentVal = editor.getData();
				resetAndMinimizeCommentEntry()
		
				var newCommentAttachments = getNewCommentAttachmentList($commentContainer)
		
				var currRecordRef = recordProxy.getRecordFunc()
				
				if(nonEmptyStringVal(commentVal) || (newCommentAttachments.length > 0)) {
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
				disableInlineCKEditor($commentInput,editor)
				resetAndMinimizeCommentEntry()
			})
			
			var $cancelButton = commentCancelCommentButtonFromContainer($commentContainer)
			initButtonControlClickHandler($cancelButton,function() {
				disableInlineCKEditor($commentInput,editor)
				clearNewCommentAttachmentList($commentContainer)
				resetAndMinimizeCommentEntry()
			})
			
		}
	})		
}