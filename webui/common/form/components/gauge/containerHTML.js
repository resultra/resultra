function gaugeContainerHTML() {
	return '' +
		'<div class="layoutContainer gaugeComponent">' +
			'<label>Gauge</label>' + 
			'<div class="formComponentGauge">' +
				'<span class="gaugeControl"></span>'+
  			'</div>' +
		'</div>'
}

function setGaugeComponentLabel($gaugeContainer, gaugeRef) {

	var $label = $gaugeContainer.find('label')
	
	setFormComponentLabel($label,gaugeRef.properties.fieldID,
			gaugeRef.properties.labelFormat)	
}

function initGaugeComponentControl($gauge,gaugeConfig) {
	var $gaugeControlContainer = $gauge.find(".gaugeControl")
	var gaugeControl = new GaugeUIControl($gaugeControlContainer, gaugeConfig);
	gaugeControl.render()
	gaugeControl.redraw(gaugeConfig.min)
	$gauge.data("gaugeControl",gaugeControl) 
	
}

function initGaugeComponentGaugeControl($gauge,gaugeObjectRef) {
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
	
	initGaugeComponentControl($gauge,gaugeConfig)
}
