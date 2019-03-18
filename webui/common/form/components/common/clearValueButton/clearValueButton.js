// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.


function smallClearButtonHTML(iconClass, className) {
	
	// className is to uniquely identify the button with other HTML elements,
	// such that it can be found with jQuery's find() function.
	
	var buttonHTML = '<button type="button" tabindex="-1" class="btn btn-default btn-sm clearButton input-group-addon ' + 
			className + 
			'"><span class="'+iconClass+'"></span></button>'
	
	return buttonHTML
}

function smallClearDeleteButtonHTML(className) {
	return smallClearButtonHTML("glyphicon glyphicon-remove",className)
}

function clearValueButtonHTML(className) {
	
	// className is to uniquely identify the button with other HTML elements,
	// such that it can be found with jQuery's find() function.
	
	var buttonHTML = '' +
		'<span class="input-group-btn" style="display:none;">' +
			'<button type="button" tabindex="-1" class="btn btn-default btn-sm clearValueButton ' + className + '" style="display:none;">' + 
				'<small><i class="glyphicon glyphicon-remove"></i></small>' +
			'</button>' +
	    '</span'
	
	return buttonHTML
}

function smallClearComponentValHeaderButton(className) {
	
	var buttonHTML = '<button tabindex="-1" class="btn btn-default btn-sm clearButton pull-right ' + 
			className + '" style="display:none;""><span class="glyphicon glyphicon-remove"></span></button>'

	return buttonHTML
}

function clearValueControlIsEnabled(controlRef) {
	var fieldID = controlRef.properties.fieldID
	var fieldRef = getFieldRef(fieldID)
	if(fieldRef.isCalcField) {
		return false
	}
	if(formComponentIsReadOnly(controlRef.properties.permissions)) {
		return false
	}
	if(controlRef.properties.clearValueSupported) {
		return true
	}
	return false
	
}


function initClearValueControl($container,controlRef,buttonClassName) {
	var $clearValueButton = $container.find(buttonClassName)
	
	// Also hide the buttons button group, since this can add unwanted
	// space to the DOM layout, even when the button is hidden. In other
	// words, completely hide everything in the DOM related to the clear value button.
	var $inputGroupParent = $clearValueButton.parent('.input-group-btn')
	
	if (clearValueControlIsEnabled(controlRef)) {
		$clearValueButton.css("display","")
		$inputGroupParent.css("display","")
	} else {
		$clearValueButton.css("display","none")
		$inputGroupParent.css("display","none")
	}	
}