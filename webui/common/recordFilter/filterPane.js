function populateFilterPanelWithOneFilterRule(filterRuleRef)
{
	var fieldName = filterRuleRef.fieldRef.fieldInfo.name
	var ruleLabel = filterRuleRef.filterRuleDef.label
	
	// TODO - Filter rule items need better formatting & CSS style
	
	
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
