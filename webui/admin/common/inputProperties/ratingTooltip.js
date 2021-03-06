// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.


function initRatingTooltipProperties(params) {
	
	function getTooltipText() {
		
		var tooltipText = []
		
		$(".ratingTooltipText").each(function(index,val) {
			tooltipText.push($(this).val())
		})
		
		return tooltipText
	}
	
	// Populate the property panel with text boxes for each rating and initialize the text boxes
	// with any tooltips already in the properties of ratingRef.
	var numRatings = params.maxVal-params.minVal
	$('#ratingTooltipPropertiesFormGroup').empty()
	for(var ratingIndex = 0; ratingIndex < numRatings; ratingIndex++) {
		var tooltipInputHTML = '<textarea class="form-control ratingTooltipText" rows="2"></textarea>'
		var $tooltipInput = $(tooltipInputHTML)
		if(params.initialTooltips[ratingIndex] != undefined) {
			var ratingText = params.initialTooltips[ratingIndex]
			$tooltipInput.text(ratingText)
		}
		$('#ratingTooltipPropertiesFormGroup').append($tooltipInput)
	}
	
	$(".ratingTooltipText").blur(function() {
		console.log("Tooltip text changed: " + $(this).val())
		
		var updatedTooltips = getTooltipText()
		console.log("Tooltip text changed: " + updatedTooltips)
		params.setTooltips(updatedTooltips)
		
	})
}