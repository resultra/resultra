

function initFormComponentTimelinePane(timelineParams) {
	
	var commentAddButtonSelector = createPrefixedSelector(timelineParams.elemPrefix,'TimelineCommentAddButton')
	var commentTextSelector = createPrefixedSelector(timelineParams.elemPrefix,'TimelineCommentText')
	var timelineListSelector = createPrefixedSelector(timelineParams.elemPrefix,'TimelineList')

	$(commentTextSelector).val("")
	
	function populateOneTimelineComment(comment) {
		
		var formattedUserName = "@" + comment.UserName
		if(comment.isCurrentUser) {
				formattedUserName = formattedUserName + ' (you)'
		}
		
		var formattedCreateDate = moment(comment.commentDate).calendar()

		var commentHTML =  '<div class="list-group-item">' +
			'<div><small>' + formattedUserName  + ' - ' + formattedCreateDate + '</small></div>' +
			'<div class="formTimelineComment">' + escapeHTML(comment.comment) + '</div>' +
		'</div>';		
		
		$(timelineListSelector).prepend(commentHTML)
	}
	
	function initTimelineWithComments() {
		
		$(timelineListSelector).empty()
		if(timelineParams.componentLink.linkedValType == linkedComponentValTypeField) {
			var getCommentParams = {
				fieldID: timelineParams.componentLink.fieldID,
				recordID:timelineParams.recordID }
			jsonAPIRequest("timeline/getFieldComments", getCommentParams, function(comments) {
				for (var commentIndex = 0; commentIndex < comments.length; commentIndex++) {
					var currComment = comments[commentIndex]
					populateOneTimelineComment(currComment)
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
				populateOneTimelineComment(newComment)
			})
			
		}

	})
}