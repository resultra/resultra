function numberInputContainerHTML(elementID)
{
	var containerHTML = ''+
		'<div class="layoutContainer numberInputComponent">' +
			'<div class="form-group">'+
				'<label>New Number Input</label>'+
				'<input type="number" class="numberInputComponentInput form-control" placeholder="Enter">'+
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