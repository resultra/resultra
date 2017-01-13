
function userSelectionIDFromElemID(elementID) {
	return "userSelection_"+elementID
}

function userSelectionContainerHTML(elementID)
{
	var controlID = userSelectionIDFromElemID(elementID)
	
	var containerHTML = ''+
		'<div class="ui-widget-content layoutContainer userSelectionFormContainer" id="'+elementID+'">' +
			'<div class="form-group">'+
				'<label for="' + controlID + '">Select User</label>'+
				'<div class="formUserSelectionControl">' + 
					'<select class="form-control" id="' + controlID + '"></select>' +
				'</div>' +
			'</div>'+
		'</div><';
										
	return containerHTML
}