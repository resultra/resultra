
function commentInputFromContainer($commentContainer) {
	return 	$commentContainer.find(".commentCommentEntryBox")
}

function commentAddCommentButtonFromContainer($commentContainer) {
	return 	$commentContainer.find(".commentComponentAddCommentButton")
}

function commentCommentListFromContainer($commentContainer) {
	return 	$commentContainer.find(".commentComponentCommentList")
}

function commentAttachmentButtonFromContainer($commentContainer) {
	return 	$commentContainer.find(".commentComponentAttachmentButton")
	
}

function commentAttachmentListFromContainer($commentContainer) {
	return 	$commentContainer.find(".newCommentAttachmentList")
	
}

function commentEntryControlsFromContainer($commentContainer) {
	return $commentContainer.find(".commentEntryControls")
}

function commentEntryContainerFromOverallCommentContainer($commentContainer) {
	return $commentContainer.find(".commentEntryContainer")
}

function commentCancelCommentButtonFromContainer($commentContainer) {
	return 	$commentContainer.find(".commentComponentCancelCommentButton")
}

function commentContainerHTML(elementID)
{	
	var containerHTML = ''+
	'<div class=" layoutContainer commentContainer">' +
	'<label>Comment Box Label</label>'+
		'<div class="form-group commentEntryContainer">' + 
			'<div class="commentCommentEntryBox inlineContent lightGreyBorder">'+
				'<p class="commentPlaceholder">Enter a comment ...</p>' +
			'</div>' +
			'<div class="commentEntryControls initiallyHidden">' +
				'<div class="newCommentAttachmentList"></div>' +
					'<button class="btn btn-default btn-xs commentComponentCancelCommentButton" type="submit">Cancel</button>' +
					'<button class="btn btn-primary btn-xs marginLeft5 commentComponentAddCommentButton" type="submit">Add Comment</button>' +
					attachmentButtonHTML("commentComponentAttachmentButton") + 
			'</div>' +
		'</div>' +				
		'<div class="list-group commentComponentCommentList lightGreyBorder"></div>' +	
	'</div>';
		
	return containerHTML
}