function openAddFilterDialog(params) {
	var $dialog = $(createPrefixedSelector(params.elemPrefix,'RecordFilterAddFilterModal'))
	
	// Dialogs need to be appended to the body element for a proper z-index. Otherwise, they 
	// might be improperly positioned underneath other elements.
	$dialog.appendTo("body").modal('show');
}