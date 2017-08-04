function initNumberConditionalFormatPropertyPanel(config) {
	
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
		
		config.setConditionalFormats(conditions)
		
	}
	
	function addConditionRule(initialFormat) {
		
		var ruleInfoByRuleID = {
			blank: {
				hasParam:false
			},
			negative: {
				hasParam:false
			},
			positive: {
				hasParam:false
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
		
		function resetParamInput(conditionVal) {
			var ruleInfo = ruleInfoByRuleID[conditionVal]
			if (ruleInfo.hasParam) {
				$paramInput.show()
				$paramInput.val("") // reset the value
			} else {
				$paramInput.hide()
			}
			
		}
		
		if (initialFormat != null) {
			resetParamInput(initialFormat.condition)
			$condSelection.val(initialFormat.condition)
			$colorSchemeSelection.val(initialFormat.colorScheme)
			$paramInput.val(initialFormat.param)
		} else {
			$paramInput.val("")
			$paramInput.hide()
		}
		
		
		initSelectControlChangeHandler($condSelection,function(newVal) {
			resetParamInput(newVal)
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
	
	for (var formatIndex = 0; formatIndex < config.initialFormats.length; formatIndex++) {
		var currFormat = config.initialFormats[formatIndex]
		addConditionRule(currFormat)
	}
	
	var $addConditionButton = $('#conditionalNumberFormatAddConditionButton')
	initButtonControlClickHandler($addConditionButton,function() {
		addConditionRule(null)
	})
}