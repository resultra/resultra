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



function initCommentBoxRecordEditBehavior($commentContainer, componentContext,recordProxy, commentObjectRef,commentBoxHeight) {
				
	$commentContainer.data("componentContext",componentContext)


	function resizeEntryAreaForCommentComponentPermissions() {
	
		var $entryContainer = commentEntryContainerFromOverallCommentContainer($commentContainer)
		if(formComponentIsReadOnly(commentObjectRef.properties.permissions)) {
			$entryContainer.hide()
		} else {
			$entryContainer.show()
		
		}
	
		// Set the maximum height of the comment area to be the remainder of the comment components
		var entryBottom = $entryContainer.position().top + $entryContainer.outerHeight(true);
		var listMaxHeightPx = (commentBoxHeight - (entryBottom+5)) + "px"
		var $commentList = commentCommentListFromContainer($commentContainer)
		$commentList.css('max-height',listMaxHeightPx)
	
	}


	function loadRecordIntoCommentBox(commentElem, recordRef) {
	
		console.log("loadRecordIntoCommentBox: loading record into comment box: " + JSON.stringify(recordRef))
		
		resizeEntryAreaForCommentComponentPermissions()
	
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
	
	
	resizeEntryAreaForCommentComponentPermissions()

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

function initCommentBoxFormRecordEditBehavior($commentContainer, componentContext,recordProxy, commentObjectRef) {
	// In view mode, the height will be flexible. However, by placing a maximum width
	// on the comment list (see below), the overall height set in the form designer
	// will be preserved when the user is not editing a comment.
	setElemFixedWidthFlexibleHeight($commentContainer,commentObjectRef.properties.geometry.sizeWidth)
	
	var commentBoxHeight = commentObjectRef.properties.geometry.sizeHeight

	initCommentBoxRecordEditBehavior($commentContainer, componentContext,recordProxy, commentObjectRef,commentBoxHeight)
	
}


function initCommentBoxTableViewRecordEditBehavior($commentContainer, componentContext,recordProxy, commentObjectRef) {
	
	
	var $commentPopupLink = $commentContainer.find(".commentEditPopop")
	
	
	// TBD - Needs a popup to display the editor.
	var validateInput = function(validationCompleteCallback) {
			validationCompleteCallback(true)
	}
	
	function formatCommentLink(recordRef) {
		var changeInfoParams = {
			recordID: recordRef.recordID,
			fieldID: commentObjectRef.properties.fieldID
		}
	
		jsonAPIRequest("record/getFieldValChangeInfo",changeInfoParams,function(valChanges) {
			if(formComponentIsReadOnly(commentObjectRef.properties.permissions)) {
				if (valChanges.length > 0) {
					$commentPopupLink.css("display","")
					$commentPopupLink.text("View comments ("+valChanges.length+")")
				} else {
					$commentPopupLink.text("")
					$commentPopupLink.css("display","none")
				}
				
			} else {
				$commentPopupLink.css("display","")
				if (valChanges.length > 0) {
					$commentPopupLink.text("Edit comments ("+valChanges.length+")")
				} else {
					$commentPopupLink.text("Add comment")
				}
				
			}
		})
		
	}
	
	var currRecordRef = null
	var loadRecordIntoPopupCommentEditor = null
	function loadRecordIntoCommentEditor($commentContainer, recordRef) {
		currRecordRef = recordRef
		if(loadRecordIntoPopupCommentEditor != null) {
			loadRecordIntoPopupCommentEditor()
		}
		formatCommentLink(recordRef)
	}
		
	
	console.log("Comment table view cell: " + $commentContainer.html())
	
	$commentPopupLink.popover({
		html: 'true',
		content: function() { return commentBoxTableViewEditContainerHTML() },
		trigger: 'click',
		placement: 'auto left',
		container: "body"
	})
	
	$commentPopupLink.on('shown.bs.popover', function()
	{
	    //get the actual shown popover
	    var $popover = $(this).data('bs.popover').tip();
		
		// By default the popover takes on the maximum size of it's containing
		// element. Overridding this size allows the size to grow as needed.
		$popover.css("max-width","300px")
		// The max-height needs to be large enough to allow the comment box to
		// expand somewhat.
		$popover.css("max-height","600px")
		console.log("Popover html: " + $popover.html())
	
		// Override the popover's default z-index to be less than the popup used to display the
		// attachments and the dialog box used to add new attachments.
		$popover.css("z-index","550")
	
		
		var $commentEditorContainer = $popover.find(".commentEditorPopupContainer")
		
		var commentEditorWidth = 250
		setElemFixedWidthFlexibleHeight($commentEditorContainer,commentEditorWidth)
				
		var $closePopupButton = $commentEditorContainer.find(".closeEditorPopup")
		initButtonControlClickHandler($closePopupButton,function() {
			$commentPopupLink.popover('hide')
			loadRecordIntoPopupCommentEditor = null
		})
		
		
		console.log("Popover html: " + $commentEditorContainer.html())
		
		function loadCurrentRecordIntoPopup() {
			if(currRecordRef != null) {
				var viewConfig = $commentEditorContainer.data("viewFormConfig")
				viewConfig.loadRecord($commentEditorContainer,currRecordRef)
			}			
		}
				
		var commentBoxHeight = 250
		initCommentBoxRecordEditBehavior($commentEditorContainer, componentContext,recordProxy, commentObjectRef,commentBoxHeight)
		loadCurrentRecordIntoPopup()
		
		// Save the function pointer to load the record into the popup. If the comment is updated, this is needed, so to 
		// list of comments can be indirectly updated in the popup. 
		loadRecordIntoPopupCommentEditor = loadCurrentRecordIntoPopup

	});
	
	$commentContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoCommentEditor,
		validateValue: validateInput
	})
	
	
}