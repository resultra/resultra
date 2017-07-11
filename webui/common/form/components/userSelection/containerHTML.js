
function userSelectionControlFromUserSelectionComponentContainer($userSelContainer) {
	return $userSelContainer.find(".userSelectionCompSelectionControl")
}

function userSelectionControlContainerHTML() {
	return '<div class="input-group">'+
				'<div class="formUserSelectionControl">' + 
					'<select class="form-control userSelectionCompSelectionControl"></select>' +
				'</div>' +
				clearValueButtonHTML("userSelectionComponentClearValueButton") +
			'</div>'
	
}

function userSelectionContainerHTML(elementID)
{
	var containerHTML = ''+
		'<div class=" layoutContainer userSelectionFormContainer">' +
			'<div class="form-group marginBottom0">'+
				'<label>New Text Box</label>'+
				userSelectionControlContainerHTML() +
			'</div>'+
		'</div>';
										
	return containerHTML
}

function userSelectionTableCellContainerHTML() {
	var containerHTML = ''+
		'<div class=" layoutContainer userSelectionTableCellContainer">' +
			userSelectionControlContainerHTML() +
		'</div>';									
	return containerHTML
	
}

function setUserSelectionComponentLabel($userSelection,userSelection) {
	var $label = $userSelection.find('label')
	
	setFormComponentLabel($label,userSelection.properties.fieldID,
			userSelection.properties.labelFormat)	
	
}