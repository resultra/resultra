function numberInputContainerHTML(elementID)
{
	var containerHTML = ''+
		'<div class="layoutContainer numberInputComponent">' +
			'<div class="form-group">'+
				'<label>New Number Input</label>'+
				'<div class="input-group">'+
					'<input type="text" class="numberInputComponentInput form-control" placeholder="Enter">'+
    				'<div class="numberInputSpinnerControls">' +
      					'<button class="btn btn-default addButton" type="button"><i class="fa fa-caret-up"></i></button>'+
      					'<button class="btn btn-default subButton" type="button"><i class="fa fa-caret-down"></i></button>'+
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