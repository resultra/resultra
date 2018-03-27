

function selectionFormControlFromSelectionFormComponent($selectionComponent) {
	return $selectionComponent.find(".selectionFormComponentSelection")
}

function selectionContainerSelectionControl() {
	return 	'<div class="input-group">'+
					'<div class="selectionFormControl">' + 
						'<select class="form-control selectionFormComponentSelection"></select>' +
					'</div>' +
					clearValueButtonHTML("selectComponentClearValueButton") +
	'</div>';

}

function selectionContainerHTML(elementID)
{	
	var containerHTML = ''+
		'<div class=" layoutContainer selectionFormComponent" id="'+elementID+'">' +
			'<div class="form-group marginBottom0">'+
				'<label>Selection</label>'+
				selectionContainerSelectionControl() +
			'</div>'+
		'</div>';
	return containerHTML
}

function textSelectionTableViewContainerHTML() {
	var containerHTML = ''+
	'<div class="layoutContainer selectionTableViewComponent">' +
			selectionContainerSelectionControl() +
	'</div>';
	return containerHTML	
}

function initSelectionComponentClearValueButton($selection,selectionRef) {
	initClearValueControl($selection,selectionRef,".selectComponentClearValueButton")
}


function setSelectionComponentLabel($selection,selectionRef) {
	var $label = $selection.find('label')
	
	setFormComponentLabel($label,selectionRef.properties.fieldID,
			selectionRef.properties.labelFormat)	
	
}

function initSelectionComponentContainer($selection,selectionRef) {
	setSelectionComponentLabel($selection,selectionRef)
	initSelectionComponentClearValueButton($selection,selectionRef)
}