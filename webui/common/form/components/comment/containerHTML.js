
function commentInputFromContainer($commentContainer) {
	return 	$commentContainer.find(".commentCommentEntryBox")
}

function commentAddCommentButtonFromContainer($commentContainer) {
	return 	$commentContainer.find(".commentComponentAddCommentButton")
}

function commentCommentListFromContainer($commentContainer) {
	return 	$commentContainer.find(".commentComponentCommentList")
}


function commentContainerHTML(elementID)
{	
	var containerHTML = ''+
	'<div class="ui-widget-content layoutContainer commentContainer  draggable resizable">' +
		'<div class="field">'+
			'<label>Comment Box Label</label>'+
				'<div class="form-group">' + 
					'<textarea class="form-control commentCommentEntryBox" rows="2"></textarea>' + 
					'<button class="btn btn-primary btn-xs commentComponentAddCommentButton" type="submit">Add Comment</button>' +
				'</div>' + 
		'</div>'+
		'<div class="list-group commentComponentCommentList"></div>' +	
	
	'</div>';
		
	return containerHTML
}