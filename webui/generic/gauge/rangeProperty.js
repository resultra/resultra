// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initGaugeRangeProperties(params) {
	var $form = $('#gaugeRangePropForm')
	
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
	
	var $minVal = $('#gaugeRangeMinVal')
	$minVal.val(params.defaultMinVal)
	var $maxVal = $('#gaugeRangeMaxVal')
	$maxVal.val(params.defaultMaxVal)
	
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