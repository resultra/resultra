function initProgressRangeProperties(params) {
	var $form = $('#progressRangePropForm')
	
	var validationSettings = createInlineFormValidationSettings({
		rules: {
			progressRangeMinVal: {
				required: true,
				floatNumber: true
			},
			progressRangeMaxVal: {
				required: true,
				floatNumber:true,
				greaterThan: '#progressRangeMinVal'
			}
		},
		messages: {
			progressRangeMaxVal: {
				greaterThan: "Value must be greater than the minimum."
			}
		}
	})	
	var validator = $form.validate(validationSettings)
	
	var $minVal = $('#progressRangeMinVal')
	$minVal.val(params.initialMinVal)
	var $maxVal = $('#progressRangeMaxVal')
	$maxVal.val(params.initialMaxVal)
	
	function setRangeIfValid() {
		if(validator.valid()) {
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