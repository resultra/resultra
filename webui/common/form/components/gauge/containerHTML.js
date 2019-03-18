// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function gaugeContainerHTML() {
	return '' +
		'<div class="layoutContainer gaugeComponent">' +
			'<label>Gauge</label>' + componentHelpPopupButtonHTML() +
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
		valueFormatter: formatGaugeVal,
		thresholdVals: gaugeObjectRef.properties.thresholdVals
	}	
	initGaugeComponentControl($gauge,gaugeConfig)
}

function initGaugeFormComponentContainer($gauge, gaugeRef) {
	setGaugeComponentLabel($gauge,gaugeRef)
	initGaugeComponentGaugeControl($gauge,gaugeRef)
	initComponentHelpPopupButton($gauge, gaugeRef)
}