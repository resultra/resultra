


function initFilterRecordsElems(parentTableID) {
	
	console.log("initFilterRecordsElems: initializing filter panel")
	
	$('#filterRecordsAddFilterButton').click(function(e){
		e.preventDefault();
		console.log("add filter button clicked")
		openAddFilterDialog(parentTableID)
	})

	$( "#filterRecordsAddFilterDialog" ).dialog({ autoOpen: false })
	
	// Populate the filter panel using a JSON call to retrieve the list of filtering rules. This will
	// get more elaborate once record filtering is actually implemented, but this suffices for 
	// prototyping.
	populateFilterPanel()
	
	console.log("initFilterRecordsElems: done initializing filter panel")	

}