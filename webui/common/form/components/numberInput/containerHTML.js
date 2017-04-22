function numberInputContainerHTML(elementID)
{
	var containerHTML = ''+
		'<div class="layoutContainer numberInputComponent">' +
			'<div class="form-group">'+
				'<label>New Number Input</label>'+
				'<div class="input-group">'+
					'<input type="text" class="numberInputComponentInput form-control" placeholder="Enter">'+
    				'<div class="numberInputSpinnerControls">' +
      					'<button class="btn btn-default addButton" type="button" tabindex="-1"><i class="fa fa-caret-up"></i></button>'+
      					'<button class="btn btn-default subButton" type="button" tabindex="-1"><i class="fa fa-caret-down"></i></button>'+
   	 				'</div>' +
				'</div>'+
			'</div>'+
			'<div class="pull-right componentHoverFooter initiallyHidden">' +
				smallClearDeleteButtonHTML("numberInputComponentClearValueButton") + 
			'</div>' +
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

function numberInputComponentDisabled($numberInputContainer) {
	
	var $numberInput = $numberInputContainer.find(".numberInputComponentInput")
	return $numberInput.prop("disabled")
	
}