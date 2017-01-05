function commentElemIDFromContainerElemID(commentElemID) {
	return 	commentElemID + '_comment'
}

function commentInputIDFromContainerElemID(commentElemID) {
	return 	commentElemID + '_commentInput'
}


function commentContainerHTML(elementID)
{	
	var commentID = commentElemIDFromContainerElemID(elementID)
	var commentInputID = commentInputIDFromContainerElemID(elementID)
	
	var containerHTML = ''+
	'<div class="ui-widget-content layoutContainer commentContainer  draggable resizable" id="'+elementID+'">' +
		'<div class="field">'+
			'<label>Comment Box Label</label>'+
				'<input type="text" name="symbol" class="layoutInput form-control" placeholder="Enter" id="' + commentInputID + '">'+
		'</div>'+
	'</div>';
		
	return containerHTML
}