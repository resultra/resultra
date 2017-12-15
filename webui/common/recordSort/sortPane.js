var sortPaneContext = {}

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

function sortPaneListItem(sortPaneParams,elemPrefix) {
	
	var listItemID = sortPanelRuleListItemID(elemPrefix)
	
	var listItemHTML =  '' +
		'<div class="list-group-item row recordSortPaneRuleListItem" id="'+listItemID+'">' +
			'<div class="col-sm-10">' +
				sortFieldSelectionHTML(elemPrefix) +
				sortDirectionButtonsHTML(elemPrefix) +
			'</div>' +
			'<div class="col-sm-2">' +
				'<button type="button" class="close filterRuleListItemDeleteRuleButton">' +
					'<span aria-hidden="true">&times;</span>'+
				'</button>'+
			'</div>' +
		'</div>';
		
	var $listItem = $(listItemHTML)
		
	var $deleteSortRuleButton = $listItem.find(".filterRuleListItemDeleteRuleButton")
	initButtonControlClickHandler($deleteSortRuleButton,function() {
		$listItem.remove()
		sortPaneRuleListChanged(sortPaneParams)
		
	})
		
	return $listItem
	
}


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

function sortPaneRuleListChanged(sortPaneParams) {
	sortPaneParams.resortFunc()
	
	var sortRules = getSortPaneSortRules()
	sortPaneParams.saveUpdatedSortRulesFunc(sortRules)
	
}


function populateSortByFieldMenu(sortPaneParams,elemPrefix,sortRule) {
	
	
	var specificSortableFields = null
	if(sortPaneParams.limitToFieldList !== undefined && sortPaneParams.limitToFieldList.length > 0) {
		specificSortableFields = new IDLookupTable(sortPaneParams.limitToFieldList)
	}
	
	var selectionID = sortFunctionSelectionID(elemPrefix)
	var menuSelector = '#' + selectionID
	
	$(menuSelector).empty()
	$(menuSelector).append(defaultSelectOptionPromptHTML("Sort By"))
	
	function includeFieldInFieldMenu(fieldID) {
		if (specificSortableFields === null) {
			// Include all fields.
			return true
		} else if (specificSortableFields.hasID(fieldID)) {
			return true
		} else if ((sortRule !== null)  && (sortRule.fieldID === fieldID)){
			// Even if the field is not in the list of specific fields to allow sorting on,
			// but it is the field being selected for sorting, include the field in the list
			// of sortable fields. This can happen if sorting takes place on a table column,
			// and the table column isn't in the specific list of fields to be sortable.
			return true
		}
		return false
	}
	
	for(var fieldIndex in sortPaneParams.sortedFields) {	
		var fieldInfo = sortPaneParams.sortedFields[fieldIndex]
		if(includeFieldInFieldMenu(fieldInfo.fieldID)) {
			$(menuSelector).append(selectFieldHTML(fieldInfo.fieldID, fieldInfo.name))			
		}
	}
	
	// Initialize the menu to the field ID of the given sortRule
	if(sortRule !== null) {
		$(menuSelector).val(sortRule.fieldID)	
	}
	
	$(menuSelector).change(function(){
		var fieldID = $(menuSelector).val()
        console.log("Sort rule: list elem = " + $(this).attr('id')+ " selected field id = " + fieldID )
		sortPaneRuleListChanged(sortPaneParams)
    }); // change
	
	
}

function addSortRuleListItem(sortPaneParams,elemPrefix,sortRule) {
	
	var fieldsByID = sortPaneParams.fieldByID
	
	$('#sortPaneSortRuleList').append(sortPaneListItem(sortPaneParams, elemPrefix))
	
	populateSortByFieldMenu(sortPaneParams,elemPrefix,sortRule)
	
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
		sortPaneRuleListChanged(sortPaneParams)
	});
	
}

// Generate unique element IDs for the different sort rule list items.
var currRecordSortPaneID = 0;
function generateSortRulePrefix() {
	currRecordSortPaneID++;
	return "sortRule" + currRecordSortPaneID + "_"
}




function initSortRecordsPane(sortPaneParams) {
	
	function setSortRules(sortRules,initDoneFunc) {
		var sortableFieldTypes = [fieldTypeNumber,fieldTypeText,fieldTypeBool,fieldTypeTime]
		
		loadSortedFieldInfo(sortPaneParams.databaseID, sortableFieldTypes, function(sortedFields) {
			
				var fieldsByID = createFieldsByIDMap(sortedFields)
				sortPaneParams.fieldsByID = fieldsByID
				sortPaneParams.sortedFields = sortedFields

				$('#sortPaneSortRuleList').empty()
				for (var sortRuleIndex = 0; sortRuleIndex < sortRules.length; sortRuleIndex++) {
					var sortRule = sortRules[sortRuleIndex]
					console.log("getFormSortRules: initializing sort rule: " + JSON.stringify(sortRule))
					addSortRuleListItem(sortPaneParams,generateSortRulePrefix(),sortRule)		
				}
				if(sortRules.length ==0) {
					// If no rules are already set add at least one uninitialized sort rule
					addSortRuleListItem(sortPaneParams,generateSortRulePrefix(),null)
				}
				if(initDoneFunc != null) {
					initDoneFunc()
				}
		})		
		
	}
	
	function updateSortRules(sortRules) {
		setSortRules(sortRules,null)
	}
		
	function initDefaultSortPaneRules(initDoneFunc) {
		setSortRules(sortPaneParams.defaultSortRules,initDoneFunc)
	}
	initDefaultSortPaneRules(sortPaneParams.initDoneFunc)
	
	initButtonClickHandler('#sortRecordsAddRuleButton',function(e) {
		console.log("add rule button clicked")
		// Add a new uninitialized sort rule
		addSortRuleListItem(sortPaneParams,generateSortRulePrefix(),null)
	})
	
	initButtonClickHandler('#sortRecordResetButton',function(e) {
		console.log("reset to default button clicked")
		initDefaultSortPaneRules(null)
		sortPaneRuleListChanged(sortPaneParams)
	})
	
	this.updateSortRules = updateSortRules
	
}