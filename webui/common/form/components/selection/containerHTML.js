

function selectionFormControlFromSelectionFormComponent($selectionComponent) {
	return $selectionComponent.find(".selectionFormComponentSelection")
}

function selectionContainerHTML(elementID)
{	
	var containerHTML = ''+
		'<div class="ui-widget-content layoutContainer selectionFormComponent" id="'+elementID+'">' +
			'<div class="form-group">'+
				'<label>Selection</label>'+
				'<div class="selectionFormControl">' + 
					'<select class="form-control selectionFormComponentSelection"></select>' +
				'</div>' +
			'</div>'+
		'</div>';
	return containerHTML
}