function textBoxContainerHTML(elementID)
{
	var containerHTML = ''+
		'<div class="layoutContainer textBoxComponent">' +
			'<div class="form-group">'+
				'<label>New Text Box</label>'+
				'<input type="text" name="symbol" class="textBoxComponentInput form-control" placeholder="Enter">'+
			'</div>'+
			'<div class="pull-right componentHoverFooter initiallyHidden">' +
				smallClearDeleteButtonHTML("textBoxComponentClearValueButton") + 
			'</div>' +
		'</div>';
	return containerHTML
}


function setTextBoxComponentLabel($textBoxContainer, textBoxRef) {

	var $label = $textBoxContainer.find('label')
	
	setFormComponentLabel($label,textBoxRef.properties.fieldID,
			textBoxRef.properties.labelFormat)	
}