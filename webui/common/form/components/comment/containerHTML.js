function commentElemIDFromContainerElemID(commentElemID) {
	return 	commentElemID + '_comment'
}

function commentInputIDFromContainerElemID(commentElemID) {
	return 	commentElemID + '_commentInput'
}

function commentAddCommentButtonIDFromContainerElemID(commentElemID) {
	return 	commentElemID + '_addCommentButton'
}

function commentCommentListFromContainerElemID(commentElemID) {
	return 	commentElemID + '_commentList'
}



function commentContainerHTML(elementID)
{	
	var commentID = commentElemIDFromContainerElemID(elementID)
	var commentInputID = commentInputIDFromContainerElemID(elementID)
	var addCommentButtonID = commentAddCommentButtonIDFromContainerElemID(elementID)
	var commentListID = commentCommentListFromContainerElemID(elementID)
	
	var containerHTML = ''+
	'<div class="ui-widget-content layoutContainer commentContainer  draggable resizable" id="'+elementID+'">' +
		'<div class="field">'+
			'<label>Comment Box Label</label>'+
				'<div class="form-group">' + 
					'<textarea class="form-control commentCommentEntryBox" rows="2" id="' + commentInputID + '"></textarea>' + 
					'<button class="btn btn-primary btn-xs" type="submit" id="' + addCommentButtonID + '">Add Comment</button>' +
				'</div>' + 
		'</div>'+
		'<div class="list-group" id="'+commentListID+'"></div>' +	
	
	'</div>';
		
	return containerHTML
}