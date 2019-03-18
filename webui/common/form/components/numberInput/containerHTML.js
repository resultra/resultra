// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

function numberInputControlHTML() {
	return 	'<div class="input-group">'+
					'<input type="text" class="numberInputComponentInput form-control" placeholder="">'+
    				'<div class="numberInputSpinnerControls">' +
      					'<button class="btn btn-default addButton" type="button" tabindex="-1"><i class="fa fa-caret-up"></i></button>'+
      					'<button class="btn btn-default subButton" type="button" tabindex="-1"><i class="fa fa-caret-down"></i></button>'+
   	 				'</div>' +
       				clearValueButtonHTML("numberInputComponentClearValueButton") +
 				'</div>'

}

function numberInputContainerHTML(elementID)
{
	var containerHTML = ''+
		'<div class="layoutContainer numberInputComponent numberInputFormContainer">' +
			'<div class="form-group">'+
				'<label>New Number Input</label>' +  componentHelpPopupButtonHTML() +
				numberInputControlHTML() +
			'</div>'+
		'</div>';
	return containerHTML
}

function numberInputTableCellContainerHTML() {
	var containerHTML = ''+
		'<div class="layoutContainer numberInputComponent numberInputTableCellContainer">' +
			numberInputControlHTML() +
		'</div>';
	return containerHTML
	
}


function setNumberInputComponentLabel($numberInputContainer, numberInputRef) {

	var $label = $numberInputContainer.find('label')
	
	setFormComponentLabel($label,numberInputRef.properties.fieldID,
			numberInputRef.properties.labelFormat)	
}

function configureNumberInputButtonSpinner($numberInputContainer, numberInputRef) {
	var $spinnerControls = $numberInputContainer.find(".numberInputSpinnerControls")
	
	function hideSpinnerControls() {
		$spinnerControls.css("display","none")
	}
	
	if (!numberInputRef.properties.showValueSpinner) {
		hideSpinnerControls()
		return
	}
	
	var numberInputFieldID = numberInputRef.properties.fieldID
	var fieldRef = getFieldRef(numberInputFieldID)
	
	if(fieldRef.isCalcField) {
		hideSpinnerControls()
		return
	}
	
	if(formComponentIsReadOnly(numberInputRef.properties.permissions)) {
		hideSpinnerControls()
		return
	}
	
	// The jQuery show() method will set the display to "block", which causes the controls to display on a
	// new line.
	$spinnerControls.css("display","")
	
}

function configureNumberInputClearValueButton($numberInputContainer, numberInputRef) {
	
	initClearValueControl($numberInputContainer,numberInputRef,".numberInputComponentClearValueButton")	
}

function numberInputComponentDisabled($numberInputContainer) {
	
	var $numberInput = $numberInputContainer.find(".numberInputComponentInput")
	return $numberInput.prop("disabled")
	
}

function initNumberInputFormContainer($container,numberInput) {
		setNumberInputComponentLabel($container,numberInput)
		configureNumberInputButtonSpinner($container,numberInput)
		configureNumberInputClearValueButton($container,numberInput)
		initComponentHelpPopupButton($container, numberInput)
	
}