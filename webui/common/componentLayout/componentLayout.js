
// Reads the component layout from the DOM elements, using parentComponentLayoutSelector
// as the parent div of the layout elements.
function getComponentLayout(parentComponentLayoutSelector) {
	var componentRows = []
	$(parentComponentLayoutSelector).children('.componentRow').each(function() { 
		var rowComponents = []
		$(this).children('.layoutContainer').each(function() {
			var componentID = $(this).attr("id")
			rowComponents.push(componentID)
		})
		if (rowComponents.length > 0) {
			// Skip over empty/placeholder rows
			componentRows.push({componentIDs: rowComponents } )	
		}
	});
	
	return componentRows

}
