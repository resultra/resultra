function populateFilterPanelWithOneFilterRule(filterRuleRef)
{
	var fieldName = filterRuleRef.fieldRef.fieldInfo.name
	var ruleLabel = filterRuleRef.filterRuleDef.label
	
	// TODO - Filter rule items need better formatting & CSS style
	var filterRecordRuleItem = itemDivHTML(
		contentHTML(headerWithBodyHTML(fieldName,ruleLabel)) +
		contentHTML('<button class="ui compact icon button" style="padding:3px"><i class="remove icon"></i></button>')
	)
			
	$('#filterRecordsRuleList').append(filterRecordRuleItem)
	
}


function populateFilterPanel()
{
	var getFilterRulesParams = {} // Params are initially empty. TODO - Add parameters for which rules to retrieve
	jsonAPIRequest("getRecordFilterRules",getFilterRulesParams,function(filterRuleRefs) {
		for (ruleIter in filterRuleRefs) {
			filterRuleRef = filterRuleRefs[ruleIter]
			populateFilterPanelWithOneFilterRule(filterRuleRef)
		}
	}) // set record's number field value
	
}
