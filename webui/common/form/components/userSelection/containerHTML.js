
function userSelectionControlFromUserSelectionComponentContainer($userSelContainer) {
	return $userSelContainer.find(".userSelectionCompSelectionControl")
}

function userSelectionContainerHTML(elementID)
{
	var containerHTML = ''+
		'<div class="ui-widget-content layoutContainer userSelectionFormContainer">' +
			'<div class="form-group">'+
				'<label>Select User</label>'+
				'<div class="formUserSelectionControl">' + 
					'<select class="form-control userSelectionCompSelectionControl"></select>' +
				'</div>' +
			'</div>'+
		'</div><';
										
	return containerHTML
}