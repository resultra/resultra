function loadRecordIntoGauge($gaugeContainer, recordRef) {
		
	var gaugeObjectRef = getContainerObjectRef($gaugeContainer)
	
	var gaugeFieldID = gaugeObjectRef.properties.fieldID
	
	function setGaugeVal(gaugeVal) {
		
		var gaugeControl = $gaugeContainer.data("gaugeControl")
		gaugeControl.redraw(gaugeVal)
	}
		
	// Populate the "intersection" of field values in the record
	// with the fields shown by the layout's containers.
	if(recordRef.fieldValues.hasOwnProperty(gaugeFieldID)) {
		var fieldVal = recordRef.fieldValues[gaugeFieldID]
		setGaugeVal(fieldVal)

	} // If record has a value for the current container's associated field ID.
	else
	{
		setGaugeVal(0.0)
	}	
	
}

function initGaugeRecordEditBehavior($gauge,componentContext,recordProxy, gaugeObjectRef) {
		
	$gauge.data("viewFormConfig", {
		loadRecord: loadRecordIntoGauge
	})
	
	function formatGaugeVal(val) {
		var formattedVal = formatNumberValue(gaugeObjectRef.properties.valueFormat.format,val)
		return formattedVal
	}
	
	var gaugeConfig = 
	{
		size: gaugeObjectRef.properties.geometry.sizeWidth,
		min: gaugeObjectRef.properties.minVal,
		max: gaugeObjectRef.properties.maxVal,
		minorTicks: 5,
		valueFormatter: formatGaugeVal
	}
	
	gaugeConfig.yellowZones = [];
	gaugeConfig.redZones = [];
	gaugeConfig.greenZones = [];
	var thresholdZones = convertStartingThresholdsToZones(gaugeObjectRef.properties.thresholdVals,
			gaugeObjectRef.properties.minVal,gaugeObjectRef.properties.maxVal)
	for (var zoneIndex = 0; zoneIndex < thresholdZones.length; zoneIndex++) {
		var currZone = thresholdZones[zoneIndex]
		switch (currZone.colorScheme) {
		case "warning":
			gaugeConfig.yellowZones.push({from:currZone.min,to:currZone.max})
			break
		case "danger":
			gaugeConfig.redZones.push({from:currZone.min,to:currZone.max})
			break
		case "success":
			gaugeConfig.greenZones.push({from:currZone.min,to:currZone.max})
			break
		}
	}
		
	var $gaugeControlContainer = $gauge.find(".gaugeControl")
	
	var gaugeControl = new GaugeUIControl($gaugeControlContainer, gaugeConfig);
	gaugeControl.render()
	gaugeControl.redraw(0)
	
	$gauge.data("gaugeControl",gaugeControl) 
	
}