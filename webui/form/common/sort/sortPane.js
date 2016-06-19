var sortPaneContext = {}

function getSortPaneSortRules() {
	var sortRules = []
	$(".recordSortPaneRuleListItem").each(function() {
		var elemPrefix = $(this).data("elemPrefix")
		var sortDirection = $(sortRuleDirectionCheckedRadioSelector(elemPrefix)).val()
		var selectedFieldID = $(sortRuleFieldSelectionMenuSelector(elemPrefix)).val()
		console.log("Sort pane rule: " + elemPrefix + " field ID = " + selectedFieldID
			+ " direction=" + sortDirection)
		
		if(selectedFieldID != null && selectedFieldID.length > 0) {
			sortRules.push({
				fieldID: selectedFieldID,
				direction: sortDirection
			})
			
		}
	})
	
	console.log("Sort rules: " + JSON.stringify(sortRules))
	
	return sortRules;
}

function sortPaneRuleListChanged() {
	sortPaneContext.resortCallback()
}

function sortFunctionSelectionID(elemPrefix) {
	var fieldSelectionID = elemPrefix + "SortFieldSelection"
	return fieldSelectionID
}

function sortRuleFieldSelectionMenuSelector(elemPrefix) {
	var selectionID = sortFunctionSelectionID(elemPrefix)
	var menuSelector = '#' + selectionID
	return menuSelector
}

function sortFieldSelectionHTML(elemPrefix) {
	
	var selectionID = sortFunctionSelectionID(elemPrefix)
		
	return '' + 
		'<div class="row">' +
			'<select class="form-control input-sm" id="'+ selectionID + '"></select>' +
		'</div>';
}

function sortRuleDirectionRadioName(elemPrefix) {
	var radioName = elemPrefix + "SortDirectionRadio"
	return radioName
}

function sortRuleDirectionCheckedRadioSelector(elemPrefix) {
	var radioName = sortRuleDirectionRadioName(elemPrefix) 
	var radioSelector = 'input[type=radio][name='+radioName+']:checked'
	return radioSelector
}


function sortDirectionButtonsHTML(elemPrefix) {
	
		var radioName = sortRuleDirectionRadioName(elemPrefix)
	
		return '' + 
			'<div class="row recordSortDirectionRow">' +
				'<div class="btn-group" data-toggle="buttons">' +
					  '<label class="btn btn-default active btn-sm">' +
							'<span class="glyphicon glyphicon-sort-by-attributes" aria-hidden="true"></span>' +
					    	'<input type="radio" name="'+ radioName + '" value="asc" autocomplete="off" checked> Ascending' +
					  '</label>' +
					  '<label class="btn btn-default btn-sm">' +
						 	'<span class="glyphicon glyphicon-sort-by-attributes-alt" aria-hidden="true"></span>' +
					    	'<input type="radio" name="'+ radioName + '"  value = "desc" autocomplete="off"> Descending' +
					  '</label>' +
				'</div>' +
			'</div>';
}

function sortPanelRuleListItemID(elemPrefix) {
	return elemPrefix + "SortRuleListItem"
}

function sortPaneListItemHTML(elemPrefix) {
	
	var listItemID = sortPanelRuleListItemID(elemPrefix)
	
	return '' +
		'<div class="list-group-item recordSortPaneRuleListItem" id="'+listItemID+'">' +
			'<div class="container-fluid">' +
				sortFieldSelectionHTML(elemPrefix) +
				sortDirectionButtonsHTML(elemPrefix) +
			'</div>' +
		'</div>';
	
}

function populateSortByFieldMenu(fieldsByID, elemPrefix) {
	
	var selectionID = sortFunctionSelectionID(elemPrefix)
	var menuSelector = '#' + selectionID
	
	$(menuSelector).empty()
	$(menuSelector).append(defaultSelectOptionPromptHTML("Sort By"))
	$.each(fieldsByID, function(fieldID, fieldInfo) {
		$(menuSelector).append(selectFieldHTML(fieldID, fieldInfo.name))		
	})
	
	$(menuSelector).change(function(){
		var fieldID = $(menuSelector).val()
        console.log("Sort rule: list elem = " + $(this).attr('id')+ " selected field id = " + fieldID )
		sortPaneRuleListChanged()
    }); // change
	
	
}

function addSortRuleListItem(elemPrefix) {
	
	var fieldsByID = getFieldsByID()
	
	$('#sortPaneSortRuleList').append(sortPaneListItemHTML(elemPrefix))
	
	populateSortByFieldMenu(fieldsByID,elemPrefix)
	
	var listItemID = sortPanelRuleListItemID(elemPrefix)
	var listItemSelector = '#' + listItemID
	$(listItemSelector).data("elemPrefix",elemPrefix)
	
	var radioName = sortRuleDirectionRadioName(elemPrefix) 
	var radioSelector = 'input[type=radio][name='+radioName+']'
	
	$(radioSelector).change(function() {
		console.log("Sort direction changed: radio name = " + radioName + " direction = " + this.value)
		sortPaneRuleListChanged()
	});
	
}

// Generate unique element IDs for the different sort rule list items.
var currRecordSortPaneID = 0;
function generateSortRulePrefix() {
	currRecordSortPaneID++;
	return "sortRule" + currRecordSortPaneID + "_"
}

function initFormSortRecordsPane(resortCallback) {
	
	sortPaneContext.resortCallback = resortCallback
	
	addSortRuleListItem(generateSortRulePrefix())
	addSortRuleListItem(generateSortRulePrefix())
	
}