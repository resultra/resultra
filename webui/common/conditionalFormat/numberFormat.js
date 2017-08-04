function initNumberConditionalFormatPropertyPanel() {
	
	var $ruleList = $("#numberConditionFormatRuleList")
	$ruleList.empty()
	
	function addConditionRule() {
		var $ruleListItem = $('#numberConditionalFormatRuleListItem').clone()
		$ruleListItem.attr("id","")
		
		var $deleteRuleButton = $ruleListItem.find(".conditionDeleteRuleButton")
		initButtonControlClickHandler($deleteRuleButton,function() {
			$ruleListItem.remove()
		})
		
		$ruleList.append($ruleListItem)
	}
	
	var $addConditionButton = $('#conditionalNumberFormatAddConditionButton')
	initButtonControlClickHandler($addConditionButton,function() {
		addConditionRule()
	})
}