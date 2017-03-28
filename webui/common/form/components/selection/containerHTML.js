

function selectionFormControlFromSelectionFormComponent($selectionComponent) {
	return $selectionComponent.find(".selectionFormComponentSelection")
}

function selectionContainerHTML(elementID)
{	
	var containerHTML = ''+
		'<div class=" layoutContainer selectionFormComponent" id="'+elementID+'">' +
			'<div class="form-group marginBottom0">'+
				'<label>Selection</label>'+
				'<div class="selectionFormControl">' + 
					'<select class="form-control selectionFormComponentSelection"></select>' +
				'</div>' +
			'</div>'+
			'<div class="pull-right componentHoverFooter initiallyHidden">' +
				smallClearDeleteButtonHTML("selectComponentClearValueButton") + 
			'</div>' +
		'</div>';
	return containerHTML
}

function setSelectionComponentLabel($selection,selectionRef) {
	var $label = $selection.find('label')
	
	setFormComponentLabel($label,selectionRef.properties.fieldID,
			selectionRef.properties.labelFormat)	
	
}