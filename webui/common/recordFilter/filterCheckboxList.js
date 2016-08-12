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