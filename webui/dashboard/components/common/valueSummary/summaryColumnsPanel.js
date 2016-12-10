

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



function populateSummarizeByMenu(panelParams,elemPrefix,fieldsByID,fieldType,initialSummary) {
	
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
		var valSummaries = getSummaryColumnValSummaries(panelParams.listElemPrefix)
		panelParams.setColumnsFunc(valSummaries)
    }); // change
	
	
}

function populateSummaryColFieldMenu(panelParams,elemPrefix,fieldsByID, initialFieldID) {
	
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
		populateSummarizeByMenu(panelParams,elemPrefix,fieldsByID,fieldInfo.type,null)
    }); // change
	
	
}


function summaryColListSelector(elemPrefix) {
	var colsListSelector = '#' + elemPrefix + 'SummaryColsList'
	return colsListSelector
}

function addColumnSummaryListItem(panelParams, valSummary, fieldsByID) {
	
	var elemPrefix = generateValSummaryElemPrefix()
	
	var colsListSelector = summaryColListSelector(panelParams.listElemPrefix)
	$(colsListSelector).append(summaryColumnListItemHTML(elemPrefix))
	
	populateSummaryColFieldMenu(panelParams,elemPrefix,fieldsByID,valSummary.summarizeByFieldID)
	var fieldInfo = fieldsByID[valSummary.summarizeByFieldID]
	populateSummarizeByMenu(panelParams,elemPrefix,fieldsByID,fieldInfo.type,valSummary.summarizeValsWith)		
}

function addNewColumnSummaryListItem(panelParams, fieldsByID) {
	var elemPrefix = generateValSummaryElemPrefix()

	var colsListSelector = summaryColListSelector(panelParams.listElemPrefix)
	$(colsListSelector).append(summaryColumnListItemHTML(elemPrefix))

	populateSummaryColFieldMenu(panelParams,elemPrefix,fieldsByID,null)
	populateSummarizeByMenu(panelParams,elemPrefix,fieldsByID,null,null)		
	
}

function initDashboardComponentSummaryColsPropertyPanel(panelParams) {
	
	loadFieldInfo(panelParams.databaseID,[fieldTypeNumber,fieldTypeText,fieldTypeBool,fieldTypeTime],
			function(valueSummaryFieldsByID) {
		var listSelector = summaryColListSelector(panelParams.listElemPrefix)
		$(listSelector).empty()
		var colValSummaries = panelParams.initialColumnValSummaries
		for(var colIndex = 0; colIndex < colValSummaries.length; colIndex++) {
			var colValSummary = colValSummaries[colIndex]

			addColumnSummaryListItem(panelParams,colValSummary,valueSummaryFieldsByID)

		}
		
		var addColumnButtonSelector = createPrefixedSelector(panelParams.listElemPrefix, 'SummaryColsAddColButton')
		initButtonClickHandler(addColumnButtonSelector,function() {
			console.log("Adding summary column")
			addNewColumnSummaryListItem(panelParams, valueSummaryFieldsByID)		
		})
		
	})
	
}

