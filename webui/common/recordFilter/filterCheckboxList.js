function addFilterToFilterCheckboxList(listSelector, elemPrefix,filterRef,filterSelectionFunc) {
		
	var filterListFilterCheckbox = createIDWithSelector(elemPrefix + filterRef.filterID)
	
	var filterItemHTML = '' +
		'<div class="checkbox list-group-item filterCheckboxListItem">' +
			'<label>' +
				'<input type="checkbox" id="' + filterListFilterCheckbox.id + '"></input>'+
				'<span class="noselect">' + filterRef.name + '</span>' +
			'</label>' +
		'</div>'
	
	$(listSelector).append(filterItemHTML)
	
	$(filterListFilterCheckbox.selector).data("filterRef",filterRef)
	$(filterListFilterCheckbox.selector).click(function(event) {
		var isChecked = $(this).prop("checked")
		
		var filterRef = $(this).data("filterRef")
		console.log("Filter checkbox clicked (for elem prefix = " + elemPrefix + "):" + filterRef.name + " checked = " + isChecked)
		filterSelectionFunc(filterRef.filterID,isChecked)
	})
	
}

function getFilterCheckboxListSelectedFilterIDs(listSelector) {
	
	var selectedFilterIDs = []
	
	// TODO - Is this selector too generic?
	var checkboxSelector = listSelector + " input[type=checkbox]:checked"
	
	$(checkboxSelector).each(function() {
		var filterRef = $(this).data("filterRef")
		selectedFilterIDs.push(filterRef.filterID)
	});
	
	return selectedFilterIDs
	
}

function selectFilterCheckboxListItem(elemPrefix,filterID) {
	var checkboxSelector = createPrefixedSelector(elemPrefix,filterID)
	$(checkboxSelector).prop("checked",true)
}

function initializeFilterCheckboxList(filterListSelector, elemPrefix, allFilters, 
				selectedFilterIDs,filterSelectionChangedFunc) {
	
	// Populate the checkbox list
	$(filterListSelector).empty()
	for(var filterIndex = 0; filterIndex < allFilters.length; filterIndex++) {
		var filterInfo = allFilters[filterIndex]
		addFilterToFilterCheckboxList(filterListSelector, elemPrefix,
			filterInfo,filterSelectionChangedFunc)
	}
	
	
	// Initialize the checkboxes which are selected.
	for(var filterIndex = 0; filterIndex < selectedFilterIDs.length;
		filterIndex++) {
			var filterID = selectedFilterIDs[filterIndex]
			selectFilterCheckboxListItem(elemPrefix,filterID)
	}
	
}