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
			equal: {
				hasParam:true
			},
			greater: {
				hasParam:true
			},
			greaterequal: {
				hasParam:true
			},
			less: {
				hasParam:true
			},
			lessequal: {
				hasParam:true
			},
			
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

function getNumberConditionalFormatBackgroundColorClassForValue(conditionalFormats, numberVal) {
		
	var formatFuncByCondition = {
		blank: function(format, val) {
			if (val===null) {
				return format.colorScheme
			}
			return null // no formatting
		},
		negative: function(format, val) {
			if((val !==null) && (val < 0.0)) {
				return format.colorScheme
			}
			return null
		},
		positive:function(format, val) {
			if((val !==null) && (val > 0.0)) {
				return format.colorScheme
			}
			return null
		},
		equal: function(format, val) {
			if((val !==null) && (val == format.param)) {
				return format.colorScheme
			}
			return null
		},
		greater: function(format, val) {
			if((val !==null) && (val > format.param)) {
				return format.colorScheme
			}
			return null
		},
		greaterequal: function(format, val) {
			if((val !==null) && (val >= format.param)) {
				return format.colorScheme
			}
			return null
		},
		less: function(format, val) {
			if((val !==null) && (val < format.param)) {
				return format.colorScheme
			}
			return null
		},
		lessequal: function(format, val) {
			if((val !==null) && (val <= format.param)) {
				return format.colorScheme
			}
			return null
		}
	}
	
	
	var formatColorScheme = null
	for(var formatIndex = 0; formatIndex < conditionalFormats.length; formatIndex++) {
		var currFormat = conditionalFormats[formatIndex]
		var formatFunc = formatFuncByCondition[currFormat.condition]
		var condFormatColor = formatFunc(currFormat,numberVal)
		if (condFormatColor !== null) {
			formatColorScheme = condFormatColor
		}
	}
	return colorClassByColorScheme(formatColorScheme)
}

function setBackgroundConditionalNumberFormat($container,conditionalFormats,numberVal) {
	
	removeConditionalFormatClasses($container)
	
	var condFormatClass = getNumberConditionalFormatBackgroundColorClassForValue(conditionalFormats,numberVal)
	if (condFormatClass !== null) {
		$container.addClass(condFormatClass)
	}
}