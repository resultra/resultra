function textBoxContainerHTML(elementID)
{
	var containerHTML = ''+
		'<div class="layoutContainer textBoxComponent">' +
			'<div class="form-group">'+
				'<label>New Text Box</label>'+
				'<div class="input-group">'+
					'<input type="text" name="symbol" class="textBoxComponentInput form-control" placeholder="Enter">'+
					'<div class="input-group-btn">'+
						'<button type="button" class="btn btn-default dropdown-toggle" ' + 
								'data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">' +
								'<span class="caret"></span></button>'+
						'<ul class="dropdown-menu">' +
						'</ul>'+
					'</div>'+
				'</div>'+
			'</div>'+
			'<div class="componentHoverFooter">' +
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