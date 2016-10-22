
function selectionFormControlID(elementID) {
	return elementID + "_selection"
}

function selectionContainerHTML(elementID)
{
	var controlID = selectionFormControlID(elementID)
	
	var containerHTML = ''+
		'<div class="ui-widget-content layoutContainer selectionFormComponent" id="'+elementID+'">' +
			'<div class="form-group">'+
				'<label for="' + controlID + '">Selection</label>'+
				'<select class="form-control" id="' + controlID + '"></select>'+
			'</div>'+
		'</div>';
	return containerHTML
}