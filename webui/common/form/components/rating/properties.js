function loadRatingProperties($rating,ratingRef) {
	console.log("Loading rating properties")
	
	function initIconProps() {
		var $iconSelection = $('#adminRatingComponentIconSelection')
		$iconSelection.val(ratingRef.properties.icon)
		initSelectControlChangeHandler($iconSelection,function(newIcon) {
		
			var iconParams = {
				parentFormID: ratingRef.parentFormID,
				ratingID: ratingRef.ratingID,
				icon: newIcon
			}
			jsonAPIRequest("frm/rating/setIcon",iconParams,function(updatedRating) {
				setContainerComponentInfo($rating,updatedRating,updatedRating.ratingID)
				reInitRatingFormComponentControl($rating,updatedRating)
			})
		
		})
		
	}
	
	initRatingTooltipProperties($rating,ratingRef)
	initIconProps()
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#ratingProps')
		
	toggleFormulaEditorForField(ratingRef.properties.fieldID)
	
}