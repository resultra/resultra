

function initFormComponentTimelinePane(timelineParams) {
	
	var commentAddButtonSelector = createPrefixedSelector(timelineParams.elemPrefix,'TimelineCommentAddButton')
	var commentTextSelector = createPrefixedSelector(timelineParams.elemPrefix,'TimelineCommentText')
	var timelineListSelector = createPrefixedSelector(timelineParams.elemPrefix,'TimelineList')

	$(commentTextSelector).val("")
	
	function createOneTimelineComment(comment) {
		
		var formattedUserName = "@" + comment.userName
		if(comment.isCurrentUser) {
				formattedUserName = formattedUserName + ' (you)'
		}
		
		var formattedCreateDate = moment(comment.commentDate).calendar()

		var commentHTML =  '<div class="list-group-item">' +
			'<div><small>' + formattedUserName  + ' - ' + formattedCreateDate + '</small></div>' +
			'<div class="formTimelineComment">' + escapeHTML(comment.comment) + '</div>' +
		'</div>';		
		
		return $(commentHTML)
	}
	
	function createOneTimelineFieldValChange(fieldValChange) {
		
		var formattedUserName = "@" + fieldValChange.userName
		if(fieldValChange.isCurrentUser) {
				formattedUserName = formattedUserName + ' (you)'			
		}
		
		var formattedCreateDate = moment(fieldValChange.updateTime).calendar()

		var valueUpdateLabel = "Changed value to " + fieldValChange.updatedValue
		
		var updateHTML =  '<div class="list-group-item">' +
			'<div><small>' + formattedUserName  + ' - ' + formattedCreateDate + '</small></div>' +
			'<div class="formTimelineComment">' + valueUpdateLabel + '</div>' +
		'</div>';		
		
		return $(updateHTML)
		
	}
	
	function initTimelineWithComments() {
		
		$(timelineListSelector).empty()
		if(timelineParams.componentLink.linkedValType == linkedComponentValTypeField) {
			var getCommentParams = {
				parentTableID: timelineParams.tableID,
				fieldID: timelineParams.componentLink.fieldID,
				recordID:timelineParams.recordID }
			jsonAPIRequest("timeline/getTimelineInfo", getCommentParams, function(timelineInfo) {
				for (var infoIndex = 0; infoIndex < timelineInfo.length; infoIndex++) {
					var currInfo = timelineInfo[infoIndex]
					if(currInfo.hasOwnProperty('commentInfo')) {
						$(timelineListSelector).append(createOneTimelineComment(currInfo.commentInfo))
					} else if (currInfo.hasOwnProperty('fieldValChangeInfo')) {
						$(timelineListSelector).append(createOneTimelineFieldValChange(currInfo.fieldValChangeInfo))
					}
				}
			})
		}
	}
	
	initTimelineWithComments()
	
	initButtonClickHandler(commentAddButtonSelector,function() {
		
		var commentText = $(commentTextSelector).val()
		console.log("Add comment button clicked: comment= " + $(commentTextSelector).val())
		$(commentTextSelector).val("")
		
		if(timelineParams.componentLink.linkedValType == linkedComponentValTypeField) {
			var saveCommentParams = {
				fieldID: timelineParams.componentLink.fieldID,
				recordID:timelineParams.recordID, 
				comment: commentText,
			}
		
			jsonAPIRequest("timeline/saveFieldComment", saveCommentParams, function(newComment) {
				$(timelineListSelector).prepend(createOneTimelineComment(newComment))
			})
			
		}

	})
}