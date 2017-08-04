function initNumberConditionalFormatPropertyPanel() {
	
	var $ruleList = $("#numberConditionFormatRuleList")
	$ruleList.empty()
	
	function updateConditionProperties() {
		var conditions = []
		$ruleList.find(".numberConditionalFormatRuleListItem").each(function() {
			var condFunc = $(this).data("getConditionPropsFunc")
			var condProp = condFunc()
			if (condProp != null) {
				conditions.push(condProp)
			}
		})
		console.log("Number conditional format: " + JSON.stringify(conditions))
		
		return conditions
	}
	
	function addConditionRule() {
		
		var ruleInfoByRuleID = {
			blank: {
				hasParam:false
			},
			negative: {
				hasParam:true
			},
			positive: {
				hasParam:true
			},
			greater: {
				hasParam:true
			},
			less: {
				hasParam:true
			}
		}
		
		
		var $ruleListItem = $('#numberConditionalFormatRuleListItem').clone()
		$ruleListItem.attr("id","")
		var $condSelection = $ruleListItem.find(".conditionTypeSelection")
		var $colorSchemeSelection = $ruleListItem.find(".conditionColorScheme")
		var $paramInput = $ruleListItem.find(".conditionParam")
		
		
		initSelectControlChangeHandler($condSelection,function(newVal) {
			var ruleInfo = ruleInfoByRuleID[newVal]
			if (ruleInfo.hasParam) {
				$paramInput.show()
				$paramInput.val("") // reset the value
			} else {
				$paramInput.hide()
			}
			updateConditionProperties()
		})

		initSelectControlChangeHandler($colorSchemeSelection,function(newVal) {
			updateConditionProperties()
		})

		
		$paramInput.blur(function() {
			updateConditionProperties()			
		})
		
		var $deleteRuleButton = $ruleListItem.find(".conditionDeleteRuleButton")
		initButtonControlClickHandler($deleteRuleButton,function() {
			$ruleListItem.remove()
			updateConditionProperties()			
		})
		
		$ruleListItem.data("getConditionPropsFunc",function() {
			var cond = $condSelection.val()
			var scheme = $colorSchemeSelection.val()
			if ((cond === null) || (cond === "") || (scheme===null) || (scheme==="")) {
				return null
			}
			var props = {
				condition:cond,
				colorScheme: scheme
			}
			var ruleInfo = ruleInfoByRuleID[cond]
			if (ruleInfo.hasParam) {
				var numberVal = convertStringToNumber($paramInput.val())
				if(numberVal === null) {
					return null	
				}
				props.param = numberVal
			}
			return props
		})
		
		$ruleList.append($ruleListItem)
	}
	
	var $addConditionButton = $('#conditionalNumberFormatAddConditionButton')
	initButtonControlClickHandler($addConditionButton,function() {
		addConditionRule()
	})
}