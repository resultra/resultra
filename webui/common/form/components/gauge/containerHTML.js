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