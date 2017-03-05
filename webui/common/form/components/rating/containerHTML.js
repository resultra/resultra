
function getRatingControlFromRatingContainer($ratingContainer) {
	var $ratingControl = $ratingContainer.find(".ratingFormComponentControl")
	assert($ratingControl !== undefined, "getRatingControlFromRatingContainer: Can't get control")
	return $ratingControl
}

function ratingContainerHTML(elementID)
{	
	var containerHTML = ''+
		'<div class=" layoutContainer ratingFormContainer">' +
			'<label class="marginBottom0">Rating</label>' +
			'<div class="formRatingControl">' +
				'<input type="hidden" class="ratingFormComponentControl"/>' + // Rating control from Bootstrap Rating plugin
			'</div>' +
		'</div><';
										
	return containerHTML
}