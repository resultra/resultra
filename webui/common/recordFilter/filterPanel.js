
var filterPaneContext = {}

function refilterWithCurrentlySelectedFilters() {
	// Iterate over checkboxes which are descendants of #filterRecordsPanelFilterList
	// and build a list of currently selected filters.
	var selectedFilters = []
	var selectedFilterIDs = []
	$('#filterRecordsPanelFilterList input[type=checkbox]:checked').each(function() {
		var filterRef = $(this).data("filterRef")
		selectedFilters.push(filterRef.name)
		selectedFilterIDs.push(filterRef.filterID)
	});
	console.log("Selected filters: " + JSON.stringify(selectedFilters))
	console.log("Selected filterIDs: " + JSON.stringify(selectedFilterIDs))
	
	// The pane doesn't actually trigger the refiltering of records, but supports a
	// callback to do so.
	var refilterParams = { tableID: filterPaneContext.tableID,
		 filterIDs: selectedFilterIDs }
	filterPaneContext.refilterCallbackFunc(refilterParams)
}

function filterPaneUnselectAllFilters() {
	$('#filterRecordsPanelFilterList input[type=checkbox]:checked').each(function() {
		$(this).prop("checked",false)
	});
	refilterWithCurrentlySelectedFilters()
}

function addFilterToFilterPanelList(filterRef) {
	
	var filterID = filterRef.filterID;
	var itemPrefix = "filterPanelFilterItem_"
	
	var filterListFilterCheckbox = createIDWithSelector(itemPrefix + filterID)
	
	var filterItemHTML = '' +
		'<div class="checkbox list-group-item filterPanelFilterItem">' +
			'<label>' +
				'<input type="checkbox" id="' + filterListFilterCheckbox.id + '"></input>'+
				'<span class="noselect">' + filterRef.name + '</span>' +
			'</label>' +
		'</div>'
	
	$('#filterRecordsPanelFilterList').append(filterItemHTML)
	
	$(filterListFilterCheckbox.selector).data("filterRef",filterRef)
	$(filterListFilterCheckbox.selector).click(function(event) {
		var isChecked = $(this).prop("checked")
		
		var filterRef = $(this).data("filterRef")
		console.log("Checkbox clicked:" + filterRef.name + " checked = " + isChecked)
		refilterWithCurrentlySelectedFilters()
	})
	
}

function populateFilterPanelFilterList(filterList) {
	
	$('#filterRecordsPanelFilterList').empty()	
	$.each(filterList, function(filterIndex, filterRef) {
		addFilterToFilterPanelList(filterRef)
	}) 
}

function initRecordFilterPanel(tableID,refilterCallbackFunc) {
	
	filterPaneContext = {
		tableID: tableID,
		refilterCallbackFunc: refilterCallbackFunc
	}
	
	jsonAPIRequest("filter/getList",{parentTableID:tableID},function(filterList) {
		populateFilterPanelFilterList(filterList)
	})
	
	$('#filterRecordsManageFiltersButton').unbind("click")
	$('#filterRecordsManageFiltersButton').click(function(e) {
	    console.log("Filter dropdown: Manage filters selected")
		openRecordFilterManageFiltersDialog(tableID)
		$(this).blur();
	    e.preventDefault();// prevent the default anchor functionality
	});

	$('#filterRecordsClearFiltersButton').unbind("click")
	$('#filterRecordsClearFiltersButton').click(function(e) {
		filterPaneUnselectAllFilters()
		$(this).blur();
	    e.preventDefault();// prevent the default anchor functionality
	});
	
	
	$('#filterRecordsPanelRefilterButton').unbind("click")
	$('#filterRecordsPanelRefilterButton').click(function(e) {
		refilterWithCurrentlySelectedFilters()
		$(this).blur();
	    e.preventDefault();// prevent the default anchor functionality
	});

	
}