function loadGaugeProperties($gauge,gaugeRef) {
	console.log("Loading gauge properties")
	
	function initRangeProperties() {
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
		$minVal.val(gaugeRef.properties.minVal)
		var $maxVal = $('#gaugeRangeMaxVal')
		$maxVal.val(gaugeRef.properties.maxVal)
		
		function setRangeIfValid() {
			if(validator.valid()) {
				var minVal = Number($minVal.val())
				var maxVal = Number($maxVal.val())
				
				var setRangeParams = {
					parentFormID: gaugeRef.parentFormID,
					gaugeID: gaugeRef.gaugeID,
					minVal: minVal,
					maxVal: maxVal
				}
				console.log("Setting gauge range: " + JSON.stringify(setRangeParams))
				jsonAPIRequest("frm/gauge/setRange", setRangeParams, function(updatedGauge) {
					setContainerComponentInfo($gauge,updatedGauge,updatedGauge.gaugeID)
				})	
				
			}		
		}
		
		$minVal.unbind("blur")
		$minVal.blur(function() { setRangeIfValid() })
		$maxVal.unbind("blur")
		$maxVal.blur(function() { setRangeIfValid() })
		
	}
	
	initRangeProperties()
	
	function saveGaugeThresholds(newThresholdVals) {
		var setThresholdParams = {
			parentFormID: gaugeRef.parentFormID,
			gaugeID: gaugeRef.gaugeID,
			thresholdVals: newThresholdVals
		}
		jsonAPIRequest("frm/gauge/setThresholds", setThresholdParams, function(updatedGauge) {
			setContainerComponentInfo($gauge,updatedGauge,updatedGauge.gaugeID)
		})	
		
	}
	
	var elemPrefix = "gauge_"
	
	var thresholdParams = {
		elemPrefix: elemPrefix,
		saveThresholdsCallback: saveGaugeThresholds,
		initialThresholdVals: gaugeRef.properties.thresholdVals
	}
	initThresholdValuesPropertyPanel(thresholdParams)
	
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentFormID: gaugeRef.parentFormID,
			gaugeID: gaugeRef.gaugeID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/progress/setLabelFormat", formatParams, function(updatedGauge) {
			setProgressComponentLabel($gauge,updatedGauge)
			setContainerComponentInfo($gauge,updatedGauge,updatedGauge.gaugeID)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: gaugeRef.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	function saveVisibilityConditions(updatedConditions) {
		var params = {
			parentFormID: gaugeRef.parentFormID,
			gaugeID: gaugeRef.gaugeID,
			visibilityConditions: updatedConditions
		}
		jsonAPIRequest("frm/progress/setVisibility",params,function(updatedGauge) {
			setContainerComponentInfo($gauge,updatedGauge,updatedGauge.gaugeID)
		})
	}
	var visibilityParams = {
		elemPrefix: elemPrefix,
		// TODO - pass in database ID as part of the component's context, rather than reference a global.
		databaseID: designFormContext.databaseID,
		initialConditions: gaugeRef.properties.visibilityConditions,
		saveVisibilityConditionsCallback:saveVisibilityConditions
	}
	initFormComponentVisibilityPropertyPanel(visibilityParams)
	
	var formatSelectionParams = {
		elemPrefix: elemPrefix,
		initialFormat: gaugeRef.properties.valueFormat.format,
		formatChangedCallback: function (newFormat) {
			console.log("Format changed for gauge: " + newFormat)
			
			var newValueFormat = {
				format: newFormat
			}
			var formatParams = {
				parentFormID: gaugeRef.parentFormID,
				gaugeID: gaugeRef.gaugeID,
				valueFormat: newValueFormat
			}
			jsonAPIRequest("frm/gauge/setValueFormat", formatParams, function(updatedGauge) {
				setContainerComponentInfo($gauge,updatedGauge,updatedGauge.gaugeID)
			})	
			
		}
	}
	initNumberFormatSelection(formatSelectionParams)

	var deleteParams = {
		elemPrefix: elemPrefix,
		parentFormID: gaugeRef.parentFormID,
		componentID: gaugeRef.gaugeID,
		componentLabel: 'gauge',
		$componentContainer: $gauge
	}
	initDeleteFormComponentPropertyPanel(deleteParams)
	
	
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#gaugeProps')
		
	toggleFormulaEditorForField(gaugeRef.properties.fieldID)
	
}