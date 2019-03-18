// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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
				hasParam:false,
				hasEndDate:false,
				hasStartDate:false
			},
			past: {
				hasParam:false,
				hasEndDate:false,
				hasStartDate:false
			},
			future: {
				hasParam:false,
				hasEndDate:false,
				hasStartDate:false
			},
			after: {
				hasParam:false,
				hasEndDate:false,
				hasStartDate:true
			},
			before: {
				hasParam:false,
				hasEndDate:true,
				hasStartDate:false
			}
		}
		
		
		var $ruleListItem = $('#dateConditionalFormatRuleListItem').clone()
		$ruleListItem.attr("id","")
		var $condSelection = $ruleListItem.find(".conditionTypeSelection")
		var $colorSchemeSelection = $ruleListItem.find(".conditionColorScheme")
		var $paramInput = $ruleListItem.find(".conditionParam")
		
		var $startDatePicker = $ruleListItem.find(".dateCondFilterStartDateInput")
		$startDatePicker.datetimepicker()

		var $endDatePicker = $ruleListItem.find(".dateCondFilterEndDateInput")
		$endDatePicker.datetimepicker()
		
		function resetParamInputsForCondition(conditionVal) {
			var ruleInfo
			if (conditionVal === null) {
				ruleInfo = {
					hasParam:false,
					hasStartDate:false,
					hasEndDate:false
				}
			} else {
				ruleInfo = ruleInfoByRuleID[conditionVal]
			}
			
			$paramInput.val("") // reset the value
			if (ruleInfo.hasParam) {
				$paramInput.show()
			} else {
				$paramInput.hide()
			}
			
			$startDatePicker.data("DateTimePicker").date(null)
			if (ruleInfo.hasStartDate) {
				$startDatePicker.css("display","")
			} else {
				$startDatePicker.css("display","none")
			}
			
			$startDatePicker.data("DateTimePicker").date(null)
			if (ruleInfo.hasEndDate) {
				$endDatePicker.css("display","")
			} else {
				$endDatePicker.css("display","none")
			}
			
		}
		
				
		if (initialFormat != null) {
			resetParamInputsForCondition(initialFormat.condition)
			
			$condSelection.val(initialFormat.condition)
			
			var ruleInfo = ruleInfoByRuleID[initialFormat.condition]
			if(ruleInfo.hasParam) {
				$paramInput.val(initialFormat.param)
			}
			if(ruleInfo.hasStartDate) {
				var startDateMoment = moment(initialFormat.startDate)
				$startDatePicker.data("DateTimePicker").date(startDateMoment)
			}
			if(ruleInfo.hasEndDate) {
				var endDateMoment = moment(initialFormat.endDate)
				$endDatePicker.data("DateTimePicker").date(endDateMoment)			
			}
			
			$colorSchemeSelection.val(initialFormat.colorScheme)
			
			
		} else {
			resetParamInputsForCondition(null)
		}
		
	    $startDatePicker.on("dp.change", function (e) {
			console.log("Custom start date changed: " + e.date)
	        $endDatePicker.data("DateTimePicker").minDate(e.date);
			updateConditionProperties()
	    });
	    $endDatePicker.on("dp.change", function (e) {
			console.log("Custom end date changed: " + e.date)
	        $startDatePicker.data("DateTimePicker").maxDate(e.date);
			updateConditionProperties()
	    });
		
		initSelectControlChangeHandler($condSelection,function(newVal) {
			resetParamInputsForCondition(newVal)
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
				
			if (ruleInfo.hasStartDate) {
				var dateVal = $startDatePicker.data("DateTimePicker").date()
				if (dateVal == null) {
					return null
				}
				props.startDate = dateVal.utc()
			}
			
			if (ruleInfo.hasEndDate) {
				var dateVal = $startDatePicker.data("DateTimePicker").date()
				if (dateVal == null) {
					return null
				}
				props.endDate = dateVal.utc()
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
			var momentStartDate = moment(format.startDate)
			if((val !==null) && (val > momentStartDate.toDate()) ) {
				return format.colorScheme
			}
			return null
		},
		before: function(format, val) {
			var momentEndDate = moment(format.endDate)
			if((val !==null) && (val < momentEndDate.toDate()) ) {
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