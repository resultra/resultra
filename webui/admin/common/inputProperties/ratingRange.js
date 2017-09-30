function initRatingRangeProperties(params) {
	var $form = $('#ratingRangePropForm')
	
	var maxRangeArgs = {
		otherValSelector: '#ratingRangeMinVal',
		maxRange: 100
	}
	
	var validationSettings = createInlineFormValidationSettings({
		rules: {
			ratingRangeMinVal: {
				required: true,
				intNumber: true
			},
			ratingRangeMaxVal: {
				required: true,
				intNumber:true,
				greaterThan: '#ratingRangeMinVal',
				maxRange: maxRangeArgs
			}
		},
		messages: {
			ratingRangeMaxVal: {
				greaterThan: "Value must be greater than the minimum.",
				maxRange: "Maximum range of values is 100."
			}
		}
	})	
	var validator = $form.validate(validationSettings)
	
	var $minVal = $('#ratingRangeMinVal')
	$minVal.val(params.initialMinVal)
	var $maxVal = $('#ratingRangeMaxVal')
	$maxVal.val(params.initialMaxVal)
	
	function setRangeIfValid() {
		if($form.valid()) {
			var minVal = Number($minVal.val())
			var maxVal = Number($maxVal.val())			
			params.setRangeCallback(minVal,maxVal)
		}		
	}
	
	$minVal.unbind("blur")
	$minVal.blur(function() { setRangeIfValid() })
	$maxVal.unbind("blur")
	$maxVal.blur(function() { setRangeIfValid() })
	
}