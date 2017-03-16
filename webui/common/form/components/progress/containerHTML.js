function progressContainerHTML() {
	return '' +
		'<div class="layoutContainer progressComponent">' +
			'<label>Progress</label>' + 
			'<div class="progress">' +
  				'<div class="progress-bar" role="progressbar" aria-valuenow="60" aria-valuemin="0" aria-valuemax="100" style="width: 60%;">' +
    				'<span class="sr-only">60% Complete</span>' +
				'</div>' +
  			'</div>' +
		'</div>'
}

function setProgressComponentLabel($progressContainer, progressRef) {

	var $label = $progressContainer.find('label')
	
	setFormComponentLabel($label,progressRef.properties.fieldID,
			progressRef.properties.labelFormat)	
}