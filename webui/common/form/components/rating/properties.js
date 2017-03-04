function loadRatingProperties($rating,ratingRef) {
	console.log("Loading rating properties")
	
	initRatingTooltipProperties($rating,ratingRef)
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#ratingProps')
		
	toggleFormulaEditorForField(ratingRef.properties.fieldID)
	
}