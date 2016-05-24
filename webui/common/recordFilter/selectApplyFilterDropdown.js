

function initRecordFilterSelectAndApplyFilterDropdown(tableID) {
	console.log("initRecordFilterSelectAndApplyFilterDropdown: Initializing dropdown with table ID = " + tableID)
	
	$('#filterRecordsDropdown').dropdown()
	
	$('#recordFilterNewFilterMenuItem').click(function(e) {
	    console.log("Filter dropdown: New filter item selected")
	    e.preventDefault();// prevent the default anchor functionality
	});

	
	$('#recordFilterManageFiltersMenuItem').click(function(e) {
	    console.log("Filter dropdown: Manage filters selected")
	    e.preventDefault();// prevent the default anchor functionality
	});
	
	$('#recordFilterNoFilterMenuItem').click(function(e) {
	    console.log("Filter dropdown: no filter selected")
	    e.preventDefault();// prevent the default anchor functionality
	});


	$('.recordFilterFilterMenuItem').click(function(e) {
		var filterID = event.target.id
	    console.log("Filter dropdown: specific filter selected: filter ID =" + filterID)
	    e.preventDefault();// prevent the default anchor functionality
	});

	$('#recordFilterRefilterButton').click(function(e) {
	    console.log("Filter dropdown: refilter button clicked")
	    e.preventDefault();// prevent the default anchor functionality
	});



	
}
