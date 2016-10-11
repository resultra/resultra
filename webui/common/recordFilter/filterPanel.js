
function getSelectedFilterPanelFilterIDs(filterListSelector) {
	// Iterate over checkboxes which are descendants of #filterRecordsPanelFilterList
	// and build a list of currently selected filters.
	var selectedFilters = []
	var selectedFilterIDs = []
	$(filterListSelector + ' input[type=checkbox]:checked').each(function() {
		var filterRef = $(this).data("filterRef")
		selectedFilters.push(filterRef.name)
		selectedFilterIDs.push(filterRef.filterID)
	});
	console.log("Selected filters: " + JSON.stringify(selectedFilters))
	console.log("Selected filterIDs: " + JSON.stringify(selectedFilterIDs))
	
	return selectedFilterIDs
}

function getCurrentFilterPanelFilterIDsWithDefaults(elemPrefix, defaultFilterIDs,availableFilterIDs) {
	
	var availableFilterLookup = new IDLookupTable(availableFilterIDs)
	
	var filterListSelector = createPrefixedSelector(elemPrefix,'FilterRecordsPanelFilterList')
	
	
	// Start building the list of current filters based upon those explicitely selected in the panel.
	var currentFilters = getSelectedFilterPanelFilterIDs(filterListSelector)
	
	// Add any default filters which are not part of the availableFitlerIDs (i.e., any default filters which are
	// not shown for selection in the filter panel). In other words, defaultFilterIDs which are not part of 
	// availableFilterIDs serve as a base set of filters which are always enabled. 
	for(var defaultFilterIndex = 0; defaultFilterIndex < defaultFilterIDs.length; defaultFilterIndex++) {
		var currDefaultFilterID = defaultFilterIDs[defaultFilterIndex]
		if(!availableFilterLookup.hasID(currDefaultFilterID)) {
			currentFilters.push(currDefaultFilterID)
		}
	}
	
	return currentFilters
}


function filterPaneUnselectAllFilters(filterListSelector) {
	$(filterListSelector + ' input[type=checkbox]:checked').each(function() {
		$(this).prop("checked",false)
	});
}

function addFilterToFilterPanelList(filterPaneParams, defaultFilterLookup, filterRef) {
	
	var filterID = filterRef.filterID;
	var itemPrefix = filterPaneParams.elemPrefix + "filterPanelFilterItem_"
	
	var filterListFilterCheckbox = createIDWithSelector(itemPrefix + filterID)
	var filterListSelector = createPrefixedSelector(filterPaneParams.elemPrefix,'FilterRecordsPanelFilterList')
	
	var filterItemHTML = '' +
		'<div class="checkbox list-group-item filterPanelFilterItem">' +
			'<label>' +
				'<input type="checkbox" id="' + filterListFilterCheckbox.id + '"></input>'+
				'<span class="noselect">' + filterRef.name + '</span>' +
			'</label>' +
		'</div>'
	
	$(filterListSelector).append(filterItemHTML)
	
	// If the filter is part of the default selection, then select it initially.
	if(defaultFilterLookup.hasID(filterID)) {
		$(filterListFilterCheckbox.selector).prop("checked",true)
	}
	
	$(filterListFilterCheckbox.selector).data("filterRef",filterRef)
	$(filterListFilterCheckbox.selector).click(function(event) {
		var isChecked = $(this).prop("checked")
		
		var filterRef = $(this).data("filterRef")
		console.log("Checkbox clicked:" + filterRef.name + " checked = " + isChecked)
		filterPaneParams.refilterCallbackFunc()
	})
	
}

function populateFilterPanelFilterList(filterPaneParams,filterList) {
	
	
	var availableFilterLookup = new IDLookupTable(filterPaneParams.availableFilterIDs)
	var defaultFilterLookup = new IDLookupTable(filterPaneParams.defaultFilterIDs)
	
	var filterListSelector = createPrefixedSelector(filterPaneParams.elemPrefix,'FilterRecordsPanelFilterList')
	
	$(filterListSelector).empty()	
	$.each(filterList, function(filterIndex, filterRef) {
		// Only show the filter if it is in the available filter IDs list.
		if(availableFilterLookup.hasID(filterRef.filterID)) {
			addFilterToFilterPanelList(filterPaneParams,defaultFilterLookup,filterRef)		
		}
	}) 
}

function initRecordFilterPanel(filterPaneParams) {
		
	var filterListSelector = createPrefixedSelector(filterPaneParams.elemPrefix,'FilterRecordsPanelFilterList')
	
	jsonAPIRequest("filter/getList",{parentTableID:filterPaneParams.tableID},function(filterList) {
		populateFilterPanelFilterList(filterPaneParams,filterList)
	})
	
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
		filterPaneUnselectAllFilters(filterListSelector)
		filterPaneParams.refilterCallbackFunc()
		$(this).blur();
	    e.preventDefault();// prevent the default anchor functionality
	});
	
}