


function getFilterRecordsRuleDef(fieldsByID, fieldID, ruleID) {
	var fieldInfo = fieldsByID[fieldID]
	var typeRules = filterRulesByType[fieldInfo.type]
	var ruleDef = typeRules[ruleID]
	return ruleDef
}


function addFilterRule(newFilterRuleParams)
{
	console.log("Adding new filter rule: params = " + JSON.stringify(newFilterRuleParams))
	
	jsonAPIRequest("newRecordFilterRule",newFilterRuleParams,function(newFilterRuleRef) {
		populateFilterPanelWithOneFilterRule(newFilterRuleRef)
		// TODO - Also need to invoke a callback function to trigger an update to the view
		// (dashboard or form) which has a filter. The records shown in these views will 
		// change.
	}) // set record's number field value
}




function initFilterRecordsElems(parentTableID) {
	
	console.log("initFilterRecordsElems: initializing filter panel")
	
	$('#filterRecordsAddFilterButton').click(function(e){
		e.preventDefault();
		console.log("add filter button clicked")
		openAddFilterDialog(parentTableID)
	})

	$( "#filterRecordsAddFilterDialog" ).dialog({ autoOpen: false })
	
	// Populate the filter panel using a JSON call to retrieve the list of filtering rules. This will
	// get more elaborate once record filtering is actually implemented, but this suffices for 
	// prototyping.
	populateFilterPanel()
	
	console.log("initFilterRecordsElems: done initializing filter panel")	

}