// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function progressContainerHTML() {
	return '' +
		'<div class="layoutContainer progressComponent">' +
			'<label>Progress</label>' + componentHelpPopupButtonHTML() +
			'<div class="progress">' +
  				'<div class="progress-bar" role="progressbar" aria-valuenow="60" aria-valuemin="0" aria-valuemax="100" style="width: 60%;">' +
    				'<span class="sr-only">60% Complete</span>' +
				'</div>' +
  			'</div>' +
		'</div>'
}

function progressTableCellContainerHTML() {
	return '' +
		'<div class="layoutContainer progressTableCell">' +
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

function initProgressFormComponentContainer($container,progressRef) {
	setProgressComponentLabel($container,progressRef)
	initComponentHelpPopupButton($container, progressRef)
}