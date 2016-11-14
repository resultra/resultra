
function initRecordFilterPanel(filterPaneParams) {
		
		
	var refilterButtonSelector = createPrefixedSelector(filterPaneParams.elemPrefix,'FilterRecordsPanelRefilterButton')
	$(refilterButtonSelector).unbind("click")
	$(refilterButtonSelector).click(function(e) {
		filterPaneParams.refilterCallbackFunc()
		$(this).blur();
	    e.preventDefault();// prevent the default anchor functionality
	});
	
	
	var clearFiltersButtonSelector = createPrefixedSelector(filterPaneParams.elemPrefix,'FilterRecordsClearFiltersButton')
	$(clearFiltersButtonSelector).unbind("click")
	$(clearFiltersButtonSelector).click(function(e) {
		// TODO - Revert filters back to defaults
		filterPaneParams.refilterCallbackFunc()
		$(this).blur();
	    e.preventDefault();// prevent the default anchor functionality
	});
	
}