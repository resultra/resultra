// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

function getCheckboxControlFromCheckboxContainer($checkboxContainer) {
	var $checkboxControl = $checkboxContainer.find(".checkboxFormComponentControl")
	assert($checkboxControl !== undefined, "getCheckboxControlFromCheckboxContainer: Can't get control")
	return $checkboxControl
}



// Support the generation of unique IDs for each individual checkbox. This isn't used
// to idenfity which field the checkbox is connected to, but to connect the checkbox
// to its label, so clicking on the label will check/uncheck the checkbox as well.
var uniqueCheckboxIDForLabel = 1
function generateUniqueCheckboxIDForLabel() {
	uniqueCheckboxIDForLabel++
	return "checkboxComponent_" + uniqueCheckboxIDForLabel
}

function checkBoxContainerControlHTML() {
	
}

function checkBoxContainerHTML(elementID)
{	
	
	var uniqueID = generateUniqueCheckboxIDForLabel()
	
	var containerHTML = ''+
		'<div class=" layoutContainer checkBoxContainer checkBoxFormContainer">' +
			'<div class="pull-right">' +
				componentHelpPopupButtonHTML() +
				smallClearDeleteButtonHTML("checkBoxComponentClearValueButton") + 
			'</div>' + 
			'<div class="checkbox">' +
				'<input type="checkbox" id="'+uniqueID+'" class="checkboxFormComponentControl">' +
				'<label for="'+  uniqueID + '"class="checkboxFormComponentLabel">New Checkbox</label>' + 
			'</div>' +
		'</div>';
				
	console.log ("Checkbox HTML: " + containerHTML)
		
	return containerHTML
}

function checkBoxTableViewCellContainerHTML() {
	
	var uniqueID = generateUniqueCheckboxIDForLabel()
	
	var checkboxTableCellHTML =  ''+
		'<div class=" layoutContainer checkBoxContainer checkBoxTableCellContainer">' +
			'<div class="row">' +
				'<div class="col-xs-11">' +
					'<div class="checkbox">' +
						'<input type="checkbox" id="'+uniqueID+'" class="checkboxTableCellControl checkboxFormComponentControl">' +
						'<label for="'+  uniqueID + '"class="checkboxFormComponentLabel"></label>' +
					'</div>' +
				'</div>' +
				'<div class="col-xs-1">' +
					smallClearDeleteButtonHTML("checkBoxComponentClearValueButton") + 
				'</div>' +
			'</div>' +
		'</div>';
		
	console.log("Checkbox table cell HTML: " + checkboxTableCellHTML)
		
	return checkboxTableCellHTML

}

function initCheckBoxControl($checkbox,checkBoxObjectRef) {
	var checkboxColorSchemeClass = "checkbox-"+checkBoxObjectRef.properties.colorScheme
	$checkbox.addClass(checkboxColorSchemeClass)
	
}

function setCheckBoxComponentLabel($checkboxContainer,checkboxRef) {
	var $label = $checkboxContainer.find('.checkboxFormComponentLabel')
	
	setFormComponentLabel($label,checkboxRef.properties.fieldID,
			checkboxRef.properties.labelFormat)	
	
}



function initCheckBoxClearValueControl($checkboxContainer,checkboxRef) {
	initClearValueControl($checkboxContainer,checkboxRef,".checkBoxComponentClearValueButton")
}


function getCurrentCheckboxComponentValue($checkboxContainer) {
	var $checkbox = $checkboxContainer.find(".checkboxFormComponentControl")
	var isIndeterminate = $checkbox.prop("indeterminate")
	if (isIndeterminate) {
		return null
	} else {
		var isChecked = $checkbox.prop("checked")
		return isChecked
	}
}

function checkboxComponentIsDisabled($checkboxContainer) {
	var $checkbox = $checkboxContainer.find(".checkboxFormComponentControl")
	var disabled = $checkbox.prop("disabled")
	return disabled
	
}

function initCheckboxComponentFormContainer($checkbox, checkboxRef) {
	setCheckBoxComponentLabel($checkbox,checkboxRef)
	initCheckBoxClearValueControl($checkbox,checkboxRef)
	initComponentHelpPopupButton($checkbox, checkboxRef)
}