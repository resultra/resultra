function loadProgressProperties($progress,progressRef) {
	console.log("Loading progress indicator properties")
	
	function initRangeProperties() {
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
		$minVal.val(progressRef.properties.minVal)
		var $maxVal = $('#progressRangeMaxVal')
		$maxVal.val(progressRef.properties.maxVal)
		
		function setRangeIfValid() {
			if(validator.valid()) {
				var minVal = Number($minVal.val())
				var maxVal = Number($maxVal.val())
				
				var setRangeParams = {
					parentFormID: progressRef.parentFormID,
					progressID: progressRef.progressID,
					minVal: minVal,
					maxVal: maxVal
				}
				console.log("Setting progress range: " + JSON.stringify(setRangeParams))
				jsonAPIRequest("frm/progress/setRange", setRangeParams, function(updatedProgress) {
					setContainerComponentInfo($progress,updatedProgress,updatedProgress.progressID)
				})	
				
			}		
		}
		
		$minVal.unbind("blur")
		$minVal.blur(function() { setRangeIfValid() })
		$maxVal.unbind("blur")
		$maxVal.blur(function() { setRangeIfValid() })
		
	}
	
	initRangeProperties()
	
	function saveProgressThresholds(newThresholdVals) {
		var setThresholdParams = {
			parentFormID: progressRef.parentFormID,
			progressID: progressRef.progressID,
			thresholdVals: newThresholdVals
		}
		jsonAPIRequest("frm/progress/setThresholds", setThresholdParams, function(updatedProgress) {
			setContainerComponentInfo($progress,updatedProgress,updatedProgress.progressID)
		})	
		
	}
	
	var thresholdParams = {
		elemPrefix: "progress_",
		saveThresholdsCallback: saveProgressThresholds,
		initialThresholdVals: progressRef.properties.thresholdVals
	}
	initThresholdValuesPropertyPanel(thresholdParams)
	
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#progressProps')
		
	toggleFormulaEditorForField(progressRef.properties.fieldID)
	
}