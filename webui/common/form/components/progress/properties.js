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
	
	var elemPrefix = "progress_"
	
	var thresholdParams = {
		elemPrefix: elemPrefix,
		saveThresholdsCallback: saveProgressThresholds,
		initialThresholdVals: progressRef.properties.thresholdVals
	}
	initThresholdValuesPropertyPanel(thresholdParams)
	
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentFormID: progressRef.parentFormID,
			progressID: progressRef.progressID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/progress/setLabelFormat", formatParams, function(updatedProgress) {
			setProgressComponentLabel($progress,updatedProgress)
			setContainerComponentInfo($progress,updatedProgress,updatedProgress.progressID)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: progressRef.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	function saveVisibilityConditions(updatedConditions) {
		var params = {
			parentFormID: progressRef.parentFormID,
			progressID: progressRef.progressID,
			visibilityConditions: updatedConditions
		}
		jsonAPIRequest("frm/progress/setVisibility",params,function(updatedProgress) {
			setContainerComponentInfo($progress,updatedProgress,updatedProgress.progressID)
		})
	}
	var visibilityParams = {
		elemPrefix: elemPrefix,
		// TODO - pass in database ID as part of the component's context, rather than reference a global.
		databaseID: designFormContext.databaseID,
		initialConditions: progressRef.properties.visibilityConditions,
		saveVisibilityConditionsCallback:saveVisibilityConditions
	}
	initFormComponentVisibilityPropertyPanel(visibilityParams)
	
	
	
	var formatSelectionParams = {
		elemPrefix: elemPrefix,
		initialFormat: progressRef.properties.valueFormat.format,
		formatChangedCallback: function (newFormat) {
			console.log("Format changed for progress bar: " + newFormat)
			
			var newValueFormat = {
				format: newFormat
			}
			var formatParams = {
				parentFormID: progressRef.parentFormID,
				progressID: progressRef.progressID,
				valueFormat: newValueFormat
			}
			jsonAPIRequest("frm/progress/setValueFormat", formatParams, function(updatedProgress) {
				setContainerComponentInfo($progress,updatedProgress,updatedProgress.progressID)
			})	
			
		}
	}
	initNumberFormatSelection(formatSelectionParams)

	var helpPopupParams = {
		initialMsg: progressRef.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentFormID: progressRef.parentFormID,
				progressID: progressRef.progressID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("frm/progress/setHelpPopupMsg",params,function(updatedProgress) {
				setContainerComponentInfo($progress,updatedProgress,updatedProgress.progressID)
				updateComponentHelpPopupMsg($progress, updatedProgress)
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)


	var deleteParams = {
		elemPrefix: elemPrefix,
		parentFormID: progressRef.parentFormID,
		componentID: progressRef.progressID,
		componentLabel: 'progress indicator',
		$componentContainer: $progress
	}
	initDeleteFormComponentPropertyPanel(deleteParams)
	
	
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#progressProps')
		
	toggleFormulaEditorForField(progressRef.properties.fieldID)
	
}