

// Generate unique element IDs for the different sort rule list items.
var currValSummaryPrefix = 0;
function generateValSummaryElemPrefix() {
	currValSummaryPrefix++;
	return "valSummary" + currValSummaryPrefix + "_"
}


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
			'<select class="form-control input-sm summarizeByFieldSelection" id="'+ selectionID + '"></select>' +
		'</div>';
}

function summarizeBySelectionHTML(elemPrefix) {
	
	var selectionID = summarizeBySelectionID(elemPrefix)
		
	return '' + 
		'<div class="row">' +
			'<select class="form-control input-sm summarizeBySelection" id="'+ selectionID + '"></select>' +
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

function getSummaryColumnValSummaries(elemPrefix) {
	var valSummaries = []
	
	var listSelector = summaryColListSelector(elemPrefix)
		
	$(".summaryColumnsListItem").each(function() {
		var summaryFieldID = $(this).find(".summarizeByFieldSelection").first().val()
		var summarizeBy = $(this).find(".summarizeBySelection").first().val()
				
		if((summarizeBy.length>0) && (summaryFieldID.length > 0)) {
			valSummaries.push({
				summarizeByFieldID: summaryFieldID,
				summarizeValsWith: summarizeBy
			})
			
		}
	})
	
	console.log("Value summary properties: " + JSON.stringify(valSummaries))
	
	return valSummaries;
}



function populateSummarizeByMenu(elemPrefix,fieldsByID,fieldType,initialSummary) {
	
	var selectionID = summarizeBySelectionID(elemPrefix)
	var menuSelector = '#' + selectionID
	
	$(menuSelector).empty()
	populateSummarizeBySelection(menuSelector,fieldType)	
	
	if(initialSummary != null) {
		$(menuSelector).val(initialSummary)	
	}
	
	$(menuSelector).unbind("change")
	$(menuSelector).change(function(){
		var summarizeValWith = $(menuSelector).val()
        console.log("Value summary: list elem = " + $(this).attr('id')+ " selected summary setting = " + summarizeValWith )
		getSummaryColumnValSummaries(elemPrefix)
    }); // change
	
	
}

function populateSummaryColFieldMenu(elemPrefix,fieldsByID, initialFieldID) {
	
	var selectionID = summaryFieldSelectionID(elemPrefix)
	var menuSelector = '#' + selectionID
	
	$(menuSelector).empty()
	$(menuSelector).append(defaultSelectOptionPromptHTML("Summary Field"))
	$.each(fieldsByID, function(fieldID, fieldInfo) {
		$(menuSelector).append(selectFieldHTML(fieldID, fieldInfo.name))		
	})
	
	// Initialize the menu to the field ID of the given valSummary
	// If none is given, leave it unselected and prompt the user.
	if(initialFieldID != null) {
		$(menuSelector).val(initialFieldID)	
	}
	
	$(menuSelector).change(function(){
		var fieldID = $(menuSelector).val()
        console.log("Value Summary: list elem = " + $(this).attr('id')+ " selected field id = " + fieldID )
		var fieldInfo = fieldsByID[fieldID]
		populateSummarizeByMenu(elemPrefix,fieldsByID,fieldInfo.type,null)
    }); // change
	
	
}


function summaryColListSelector(elemPrefix) {
	var colsListSelector = '#' + elemPrefix + 'SummaryColsList'
	return colsListSelector
}

function addColumnSummaryListItem(listElemPrefix, valSummary, fieldsByID) {
	
	var elemPrefix = generateValSummaryElemPrefix()
	
	var colsListSelector = summaryColListSelector(listElemPrefix)
	$(colsListSelector).append(summaryColumnListItemHTML(elemPrefix))
	
	populateSummaryColFieldMenu(elemPrefix,fieldsByID,valSummary.summarizeByFieldID)
	var fieldInfo = fieldsByID[valSummary.summarizeByFieldID]
	populateSummarizeByMenu(elemPrefix,fieldsByID,fieldInfo.type,valSummary.summarizeValsWith)		
}

function addNewColumnSummaryListItem(listElemPrefix, fieldsByID) {
	var elemPrefix = generateValSummaryElemPrefix()

	var colsListSelector = summaryColListSelector(listElemPrefix)
	$(colsListSelector).append(summaryColumnListItemHTML(elemPrefix))

	populateSummaryColFieldMenu(elemPrefix,fieldsByID,null)
	populateSummarizeByMenu(elemPrefix,fieldsByID,null,null)		
	
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
		
		initButtonClickHandler(createPrefixedSelector(summaryTableElemPrefix, 'SummaryColsAddColButton'),function() {
			console.log("Adding summary column")
			addNewColumnSummaryListItem(summaryTableElemPrefix, valueSummaryFieldsByID)		
		})
		
	})
	
}

