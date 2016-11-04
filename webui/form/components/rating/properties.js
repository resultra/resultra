function loadRatingProperties(ratingRef) {
	console.log("Loading rating properties")
	
	initRatingTooltipProperties(ratingRef)
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#ratingProps')
		
	toggleFormulaEditorForComponent(ratingRef.properties.componentLink)
	
}