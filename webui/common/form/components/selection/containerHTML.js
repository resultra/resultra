

function selectionFormControlFromSelectionFormComponent($selectionComponent) {
	return $selectionComponent.find(".selectionFormComponentSelection")
}

function selectionContainerHTML(elementID)
{	
	var containerHTML = ''+
		'<div class=" layoutContainer selectionFormComponent" id="'+elementID+'">' +
			'<div class="form-group">'+
				'<label>Selection</label>'+
				'<div class="selectionFormControl">' + 
					'<select class="form-control selectionFormComponentSelection"></select>' +
				'</div>' +
			'</div>'+
		'</div>';
	return containerHTML
}