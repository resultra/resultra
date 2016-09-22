

function summaryFieldSelectionID(elemPrefix) {
	var fieldSelectionID = elemPrefix + "SummaryFieldSelection"
	return fieldSelectionID
}

function summarizeBySelectionID(elemPrefix) {
	var summarizeBySelectionID = elemPrefix + "SummarizeBySelection"
	return summarizeBySelectionID
}


function summaryColumnListItemID(elemPrefix) {
	return elemPrefix + "SummaryColumnsListItem"
}

function summaryFieldSelectionHTML(elemPrefix) {
	
	var selectionID = summaryFieldSelectionID(elemPrefix)
		
	return '' + 
		'<div class="row">' +
			'<select class="form-control input-sm" id="'+ selectionID + '"></select>' +
		'</div>';
}

function summarizeBySelectionHTML(elemPrefix) {
	
	var selectionID = summarizeBySelectionID(elemPrefix)
		
	return '' + 
		'<div class="row">' +
			'<select class="form-control input-sm" id="'+ selectionID + '"></select>' +
		'</div>';
}



function summaryColumnListItemHTML(elemPrefix) {
	
	var listItemID = summaryColumnListItemID(elemPrefix)
	
	return '' +
		'<div class="list-group-item summaryColumnsListItem" id="'+listItemID+'">' +
			'<div class="container-fluid">' +
				summaryFieldSelectionHTML(elemPrefix) +
				summarizeBySelectionHTML(elemPrefix) +
			'</div>' +
		'</div>';
}

function populateSummaryColFieldMenu(elemPrefix,valSummary,fieldsByID) {
	
	var selectionID = summaryFieldSelectionID(elemPrefix)
	var menuSelector = '#' + selectionID
	
	$(menuSelector).empty()
	$(menuSelector).append(defaultSelectOptionPromptHTML("Summary Field"))
	$.each(fieldsByID, function(fieldID, fieldInfo) {
		$(menuSelector).append(selectFieldHTML(fieldID, fieldInfo.name))		
	})
	
	// Initialize the menu to the field ID of the given valSummary
	// If none is given, leave it unselected and prompt the user.
	if(valSummary != null) {
		$(menuSelector).val(valSummary.summarizeByFieldID)	
	}
	
	$(menuSelector).change(function(){
		var fieldID = $(menuSelector).val()
        console.log("Value Summary: list elem = " + $(this).attr('id')+ " selected field id = " + fieldID )
//		sortPaneRuleListChanged()
    }); // change
	
	
}

function populateSummarizeByMenu(elemPrefix,valSummary,fieldsByID) {
	
	var selectionID = summarizeBySelectionID(elemPrefix)
	var menuSelector = '#' + selectionID
	
	$(menuSelector).empty()
	if (valSummary != null) {
		var fieldInfo = fieldsByID[valSummary.summarizeByFieldID]
		if (fieldInfo != null) {
			populateSummarizeBySelection(menuSelector,fieldInfo.type)	
		}		
	}
	
	// Initialize the menu to the field ID of the given valSummary
	// If none is given, leave it unselected and prompt the user.
	if(valSummary != null) {
		$(menuSelector).val(valSummary.summarizeValsWith)	
	}
	
	$(menuSelector).change(function(){
		var summarizeValWith = $(menuSelector).val()
        console.log("Value summary: list elem = " + $(this).attr('id')+ " selected summary setting = " + summarizeValWith )
//		sortPaneRuleListChanged()
    }); // change
	
	
}



function summaryColListSelector(elemPrefix) {
	var colsListSelector = '#' + elemPrefix + 'SummaryColsList'
	return colsListSelector
}

function addColumnSummaryListItem(elemPrefix,valSummary, fieldsByID) {
	
	var colsListSelector = summaryColListSelector(elemPrefix)
	$(colsListSelector).append(summaryColumnListItemHTML(elemPrefix))
	
	populateSummaryColFieldMenu(elemPrefix,valSummary,fieldsByID)
	populateSummarizeByMenu(elemPrefix,valSummary,fieldsByID)
		
	var listItemSelector = '#' + summaryColumnListItemID(elemPrefix)
	$(listItemSelector).data("elemPrefix",elemPrefix)
		
}

function initDashboardComponentSummaryColsPropertyPanel(summaryTableElemPrefix,summaryTable) {
	
	loadFieldInfo(summaryTable.properties.dataSrcTableID,[fieldTypeNumber,fieldTypeText,fieldTypeBool,fieldTypeTime],
			function(valueSummaryFieldsByID) {
		var colValSummaries = summaryTable.properties.columnValSummaries
		$(summaryColListSelector(summaryTableElemPrefix)).empty()
		for(var colIndex = 0; colIndex < colValSummaries.length; colIndex++) {
			var colValSummary = colValSummaries[colIndex]

			addColumnSummaryListItem(summaryTableElemPrefix,colValSummary,valueSummaryFieldsByID)

		}
	})
	
	
	
}

