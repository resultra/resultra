
function ratingControlIDFromElemID(elementID) {
	return "rating_"+elementID
}

function ratingContainerHTML(elementID)
{
	var controlID = ratingControlIDFromElemID(elementID)
	
	var containerHTML = ''+
		'<div class="ui-widget-content layoutContainer ratingFormContainer" id="'+elementID+'">' +
			'<label class="marginBottom0">Rating</label>' +
			'<div class="formRatingControl">' +
				'<input type="hidden" id="'+controlID+'"/>' + // Rating control from Bootstrap Rating plugin
			'</div>' +
		'</div><';
		
//	var controlSelector = '#' + ratingControlIDFromElemID(elementID)
//	$(controlSelector).rating()
								
	return containerHTML
}