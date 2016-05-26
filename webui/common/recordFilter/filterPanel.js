function initRecordFilterPanel(tableID) {
	$('#filterRecordsManageFiltersButton').click(function(e) {
	    console.log("Filter dropdown: Manage filters selected")
		openRecordFilterManageFiltersDialog(tableID)
	    e.preventDefault();// prevent the default anchor functionality
	});
	
}