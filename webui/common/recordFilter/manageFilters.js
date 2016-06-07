var manageFilterCurrFilterID = null

function populateFilterPanelWithOneFilterRule(filterRuleRef)
{
	
	var fieldName = getFieldRef(filterRuleRef.fieldID).name
	
	// TODO - Need to retrieve table mapping rule IDs to labels.
	var ruleLabel = getFilterRecordsRuleDef(filterRuleRef.fieldID, filterRuleRef.ruleID).label
	
	var filterRecordRuleItem = 
		'<div class="list-group-item clearfix filterRecordItem">' + 
			'<div class="pull-left">' +
				'<strong>' + fieldName + '</strong>' + ' <BR> ' + ruleLabel + 
        	'</div>' +	
    		'<div class="pull-right" style="margin-top:10px;">' + 
      			'<button class="btn btn-xs btn-danger deleteFilterRuleButton">' + 
					// padding-bottom: 2px makes the button image vertically line up better.
					'<span class="glyphicon glyphicon-remove" style="padding-bottom:2px;"></span>' +
				'</button>' +
        	'</div>' +	
		'</div>'
					
	$('#recordFilterManageFilterRuleList').append(filterRecordRuleItem)
	
}

function populateFilterPanel(filterRules) {
	
	$('#recordFilterManageFilterRuleList').empty()
	$.each(filterRules, function(ruleIndex, ruleRef) {
		populateFilterPanelWithOneFilterRule(ruleRef)
	}) 
	
}

function filterClickHandler(event) {
	var filterID = event.target.id
	console.log("Manage filters: filter selected: filter ID =" + filterID)

	var filterSelector = '#'+filterID

	
	$(".filterListItem").removeClass("active")
	$(filterSelector).addClass("active")
	
	
	manageFilterCurrFilterID = filterID
	var currFilterRef = $(filterSelector).data("filterRef")
	console.log("Populating filter name: " + currFilterRef.name)
	$("#recordFilterManageFilterFilterName").val(currFilterRef.name)
	
	var getFilterRulesParams = {parentFilterID:filterID}
	jsonAPIRequest("filter/getRuleList",getFilterRulesParams,function(filterRules) {
		populateFilterPanel(filterRules)
	}) // set record's number field value
	
	
	
	event.preventDefault();// prevent the default anchor functionality
}


function addFilterRefToFilterList(filterRef) {
	var filterHTML = '<a href="#" class="list-group-item filterListItem" id="' 
		+ filterRef.filterID + '">' + filterRef.name + '</a>'
	$('#recordFilterManageFilterFilterList').append(filterHTML)	
	
	var filterSelector = '#'+filterRef.filterID
	$(filterSelector).click(filterClickHandler)
	$(filterSelector).data("filterRef",filterRef)
	
}


function populateFilterList(filterList) {
	
	$('#recordFilterManageFilterFilterList').empty()	
	$.each(filterList, function(filterIndex, filterRef) {
		addFilterRefToFilterList(filterRef)
	}) 
}

function addNewFilter(tableID) {
	jsonAPIRequest("filter/newWithPrefix",{parentTableID:tableID,namePrefix:"Untitled Filter"},function(newFilterRef) {
		console.log("Filter created: " + newFilterRef.name)
		addFilterRefToFilterList(newFilterRef)
	}) // set record's number field value
	
}

function manageFiltersAddFilterRule(newFilterRuleParams)
{
	console.log("Adding new filter rule: params = " + JSON.stringify(newFilterRuleParams))
	
	if(manageFilterCurrFilterID != null) {
		newFilterRuleParams.parentFilterID = manageFilterCurrFilterID
		jsonAPIRequest("filter/newRule",newFilterRuleParams,function(newFilterRuleRef) {
			populateFilterPanelWithOneFilterRule(newFilterRuleRef)
			// TODO - Also need to invoke a callback function to trigger an update to the view
			// (dashboard or form) which has a filter. The records shown in these views will 
			// change.
		}) // set record's number field value	
	}
}

function manageFiltersChangeFilterName() {
	if(manageFilterCurrFilterID != null) {
		
		var filterSelector = '#'+manageFilterCurrFilterID
		var currFilterRef = $(filterSelector).data("filterRef")
		
		var newFilterName = $('#recordFilterManageFilterFilterName').val()
		
		console.log("manageFiltersChangeFilterName: new name = " + newFilterName + " existing name = " + currFilterRef.name)
		
		if(newFilterName != currFilterRef.name) {
			// TODO - Validate the filter name is unique updating on server.
			var updateNameParams = { filterID: manageFilterCurrFilterID, name: newFilterName }
			jsonAPIRequest("filter/setName",updateNameParams,function(updatedFilterRef) {
				// Update the name in the filter list.
				var filterSelector = '#'+updatedFilterRef.filterID
				$(filterSelector).text(updatedFilterRef.name)
				$(filterSelector).data("filterRef",updatedFilterRef)
			}) // set record's number field value				
		}
		
	}
	
}


function openRecordFilterManageFiltersDialog(tableID) {
	
	$('#recordFilterManageFilterRuleList').empty()
	initAddFilterRuleControlPanel(tableID,manageFiltersAddFilterRule)
	
	$('#filterRecordsManageFilterAddFilterButton').unbind("click")
	$('#filterRecordsManageFilterAddFilterButton').click(function(e) {
	    console.log("Adding new filter for table = " + tableID)
		addNewFilter(tableID)
	    e.preventDefault();// prevent the default anchor functionality
	});
	
	$('#recordFilterManageFilterFilterName').unbind("blur")
	$('#recordFilterManageFilterFilterName').blur(manageFiltersChangeFilterName)

	jsonAPIRequest("filter/getList",{parentTableID:tableID},function(filterList) {
		populateFilterList(filterList)
		$('#recordFilterManageRecordsModal').modal('show')	
	}) // set record's number field value
	
}