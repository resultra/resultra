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
	
	var gaugeConfig = 
	{
		size: 120,
		label: "TBD",
		min: 0,
		max: 100,
		minorTicks: 5
	}
	
	var range = gaugeConfig.max - gaugeConfig.min;
	gaugeConfig.yellowZones = [{ from: gaugeConfig.min + range*0.75, to: gaugeConfig.min + range*0.9 }];
	gaugeConfig.redZones = [{ from: gaugeConfig.min + range*0.9, to: gaugeConfig.max }];
	
	var $gaugeControlContainer = $gauge.find(".gaugeControl")
	
	var gaugeControl = new GaugeUIControl($gaugeControlContainer, gaugeConfig);
	gaugeControl.render()
	gaugeControl.redraw(0)
	
	$gauge.data("gaugeControl",gaugeControl) 
	
}