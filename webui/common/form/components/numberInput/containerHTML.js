
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
	
	var $clearValueButton = $numberInputContainer.find(".numberInputComponentClearValueButton")
	
	function hideButton() {
		$clearValueButton.css("display","none")
		$clearValueButton.prop("disabled",true)
	}
	
	function showButton() {
		$clearValueButton.css("display","")
		$clearValueButton.prop("disabled",false)
		
	}

	var numberInputFieldID = numberInputRef.properties.fieldID
	var fieldRef = getFieldRef(numberInputFieldID)
	
	if(fieldRef.isCalcField) {
		hideButton()
		return
	}
	
	
	if(formComponentIsReadOnly(numberInputRef.properties.permissions)) {
		hideButton()
	} else {
		if(numberInputRef.properties.clearValueSupported) {
			showButton()
		} else {
			hideButton()	
		}
	}
	
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