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
	
	var sortRules = getSortPaneSortRules()
	var saveSortRulesParams = {
		parentFormID: sortPaneContext.parentFormID,
		sortRules: sortRules
	}
	jsonAPIRequest("recordSort/saveFormSortRules",saveSortRulesParams,function(saveReply) {}) // getRecord
	
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

function populateSortByFieldMenu(elemPrefix,sortRule) {
	
	var selectionID = sortFunctionSelectionID(elemPrefix)
	var menuSelector = '#' + selectionID
	
	$(menuSelector).empty()
	$(menuSelector).append(defaultSelectOptionPromptHTML("Sort By"))
	$.each(sortPaneContext.fieldsByID, function(fieldID, fieldInfo) {
		$(menuSelector).append(selectFieldHTML(fieldID, fieldInfo.name))		
	})
	
	// Initialize the menu to the field ID of the given sortRule
	if(sortRule != null) {
		$(menuSelector).val(sortRule.fieldID)	
	}
	
	$(menuSelector).change(function(){
		var fieldID = $(menuSelector).val()
        console.log("Sort rule: list elem = " + $(this).attr('id')+ " selected field id = " + fieldID )
		sortPaneRuleListChanged()
    }); // change
	
	
}

function addSortRuleListItem(elemPrefix,sortRule) {
	
	var fieldsByID = getFilteredFieldsByID([fieldTypeNumber,fieldTypeText,fieldTypeBool,fieldTypeTime])
	
	$('#sortPaneSortRuleList').append(sortPaneListItemHTML(elemPrefix))
	
	populateSortByFieldMenu(elemPrefix,sortRule)
	
	var listItemID = sortPanelRuleListItemID(elemPrefix)
	var listItemSelector = '#' + listItemID
	$(listItemSelector).data("elemPrefix",elemPrefix)
	
	var radioName = sortRuleDirectionRadioName(elemPrefix) 
	var radioSelector = 'input[type="radio"][name="'+radioName+'"]'
	
	// Initialize the radio button selection based upon the sort direction. Using the
	// click() function is needed for bootstrap, so it will trigger the right changes
	// with CSS classes.
	if(sortRule != null) {
		$(':radio[name="'+radioName+'"][value="' + sortRule.direction + '"]').click()	
	}
	
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

function initFormSortRecordsPane(parentFormID, resortCallback, initDoneCallback) {
	
	sortPaneContext.resortCallback = resortCallback
	sortPaneContext.parentFormID = parentFormID
	sortPaneContext.fieldsByID = getFilteredFieldsByID([fieldTypeNumber,fieldTypeText,fieldTypeBool,fieldTypeTime])
	
	var getSortRulesParams = {
		parentFormID: parentFormID
	}
	
	jsonAPIRequest("recordSort/getFormSortRules",getSortRulesParams,function(formSortRules) {
		for (var sortRuleIndex = 0; sortRuleIndex < formSortRules.sortRules.length; sortRuleIndex++) {
			var sortRule = formSortRules.sortRules[sortRuleIndex]
			console.log("getFormSortRules: initializing sort rule: " + JSON.stringify(sortRule))
			addSortRuleListItem(generateSortRulePrefix(),sortRule)		
		}
		if(formSortRules.sortRules.length ==0) {
			// If no rules are already set add at least one uninitialized sort rule
			addSortRuleListItem(generateSortRulePrefix(),null)
		}
		initDoneCallback()
	}) // getFormSortRules
	
	
	$('#sortRecordsAddRuleButton').click(function(e) {
		console.log("add rule button clicked")
		$(this).blur();
	    e.preventDefault();// prevent the default anchor functionality
		
		// Add a new uninitialized sort rule
		addSortRuleListItem(generateSortRulePrefix(),null)
	})

	$('#sortRecordsClearRulesButton').click(function(e) {
		console.log("reset rule button clicked")
		$(this).blur();
	    e.preventDefault();// prevent the default anchor functionality
		
		$('#sortPaneSortRuleList').empty()
		addSortRuleListItem(generateSortRulePrefix(),null)
		sortPaneRuleListChanged()
	})
	
}