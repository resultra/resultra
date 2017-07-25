function loadGaugeProperties($gauge,gaugeRef) {
	console.log("Loading gauge properties")
	
	function setGaugeRange(minVal,maxVal) {
		var setRangeParams = {
			parentFormID: gaugeRef.parentFormID,
			gaugeID: gaugeRef.gaugeID,
			minVal: minVal,
			maxVal: maxVal
		}
		console.log("Setting gauge range: " + JSON.stringify(setRangeParams))
		jsonAPIRequest("frm/gauge/setRange", setRangeParams, function(updatedGauge) {
			setContainerComponentInfo($gauge,updatedGauge,updatedGauge.gaugeID)
			initGaugeComponentGaugeControl($gauge,updatedGauge)
		})			
	}
	var gaugeRangeParams = {
		defaultMinVal: gaugeRef.properties.minVal,
		defaultMaxVal: gaugeRef.properties.maxVal,
		setRangeCallback: setGaugeRange
	}
	initGaugeRangeProperties(gaugeRangeParams)
	
	function saveGaugeThresholds(newThresholdVals) {
		var setThresholdParams = {
			parentFormID: gaugeRef.parentFormID,
			gaugeID: gaugeRef.gaugeID,
			thresholdVals: newThresholdVals
		}
		jsonAPIRequest("frm/gauge/setThresholds", setThresholdParams, function(updatedGauge) {
			setContainerComponentInfo($gauge,updatedGauge,updatedGauge.gaugeID)
			initGaugeComponentGaugeControl($gauge,updatedGauge)
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
		jsonAPIRequest("frm/gauge/setLabelFormat", formatParams, function(updatedGauge) {
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
		jsonAPIRequest("frm/gauge/setVisibility",params,function(updatedGauge) {
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
				initGaugeComponentGaugeControl($gauge,updatedGauge)
			})	
			
		}
	}
	initNumberFormatSelection(formatSelectionParams)

	var helpPopupParams = {
		initialMsg: gaugeRef.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentFormID: gaugeRef.parentFormID,
				gaugeID: gaugeRef.gaugeID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("frm/gauge/setHelpPopupMsg",params,function(updatedGauge) {
				setContainerComponentInfo($gauge,updatedGauge,updatedGauge.gaugeID)
				updateComponentHelpPopupMsg($gauge, updatedGauge)
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)


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