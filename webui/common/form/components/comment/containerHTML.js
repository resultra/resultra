
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

function commentContainerHTML(elementID)
{	
	var containerHTML = ''+
	'<div class=" layoutContainer commentContainer">' +
		'<div class="field">'+
			'<label>Comment Box Label</label>'+
				'<div class="form-group">' + 
					'<textarea class="form-control commentCommentEntryBox" rows="2"></textarea>' + 
					'<div class="newCommentAttachmentList"></div>' +
					'<button class="btn btn-primary btn-xs commentComponentAddCommentButton" type="submit">Add Comment</button>' +
							attachmentButtonHTML("commentComponentAttachmentButton") + 
				'</div>' +
				
		'</div>'+
		'<div class="list-group commentComponentCommentList"></div>' +	
	
	'</div>';
		
	return containerHTML
}