function initDateConditionalFormatPropertyPanel(config) {
	
	var $ruleList = $("#dateConditionFormatRuleList")
	$ruleList.empty()
	
	function updateConditionProperties() {
		var conditions = []
		$ruleList.find(".dateConditionalFormatRuleListItem").each(function() {
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
			past: {
				hasParam:false
			},
			future: {
				hasParam:false
			},
			after: {
				hasParam:true
			}
		}
		
		
		var $ruleListItem = $('#dateConditionalFormatRuleListItem').clone()
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
	
	var $addConditionButton = $('#conditionalDateFormatAddConditionButton')
	initButtonControlClickHandler($addConditionButton,function() {
		addConditionRule(null)
	})
}

function getDateConditionalFormatBackgroundColorClassForValue(conditionalFormats, dateVal) {
	
	var formatFuncByCondition = {
		blank: function(format, val) {
			if (val===null) {
				return format.colorScheme
			}
			return null // no formatting
		},
		past: function(format, val) {
			var now = new Date()
			if((val !==null) && (val < now)) {
				return format.colorScheme
			}
			return null
		},
		future:function(format, val) {
			var now = new Date()
			if((val !==null) && (val > now)) {
				return format.colorScheme
			}
			return null
		},
		after: function(format, val) {
			if((val !==null) && (val > format.param)) {
				return format.colorScheme
			}
			return null
		}
	}
	
	
	var formatColorScheme = null
	for(var formatIndex = 0; formatIndex < conditionalFormats.length; formatIndex++) {
		var currFormat = conditionalFormats[formatIndex]
		var formatFunc = formatFuncByCondition[currFormat.condition]
		var condFormatColor = formatFunc(currFormat,dateVal)
		if (condFormatColor !== null) {
			formatColorScheme = condFormatColor
		}
	}
	return colorClassByColorScheme(formatColorScheme)
}


function setBackgroundConditionalDateFormat($container,conditionalFormats,dateVal) {
	
	removeConditionalFormatClasses($container)
	
	var condFormatClass = getDateConditionalFormatBackgroundColorClassForValue(conditionalFormats,dateVal)
	if (condFormatClass !== null) {
		$container.addClass(condFormatClass)
	}
}