
function initRatingTooltipProperties(ratingRef) {
	
	function getTooltipText() {
		
		var tooltipText = []
		
		$(".ratingTooltipText").each(function(index,val) {
			tooltipText.push($(this).val())
		})
		
		return tooltipText
	}
	
	// Populate the property panel with text boxes for each rating and initialize the text boxes
	// with any tooltips already in the properties of ratingRef.
	var numRatings = 5
	$('#ratingTooltipPropertiesFormGroup').empty()
	for(var ratingIndex = 0; ratingIndex < numRatings; ratingIndex++) {
		var tooltipInputHTML = '<textarea class="form-control ratingTooltipText" rows="2"></textarea>'
		var $tooltipInput = $(tooltipInputHTML)
		if(ratingRef.properties.tooltips[ratingIndex] != undefined) {
			var ratingText = ratingRef.properties.tooltips[ratingIndex]
			$tooltipInput.val(ratingText)
		}
		$('#ratingTooltipPropertiesFormGroup').append($tooltipInput)
	}
	
	$(".ratingTooltipText").blur(function() {
		console.log("Tooltip text changed: " + $(this).val())
		
		var updatedTooltips = getTooltipText()
		console.log("Tooltip text changed: " + updatedTooltips)
		
		var tooltipParams = {
			parentFormID: ratingRef.parentFormID,
			ratingID: ratingRef.ratingID,
			tooltips: updatedTooltips
		}
			
		jsonAPIRequest("frm/rating/setTooltips", tooltipParams, function(updateRating) {
			setElemObjectRef(updateRating.ratingID,updateRating)
		})	
		
	})
}