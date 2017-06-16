
function userSelectionControlFromUserSelectionComponentContainer($userSelContainer) {
	return $userSelContainer.find(".userSelectionCompSelectionControl")
}

function userSelectionContainerHTML(elementID)
{
	var containerHTML = ''+
		'<div class=" layoutContainer userSelectionFormContainer">' +
			'<div class="form-group marginBottom0">'+
				'<label>Select User</label>'+
				'<div class="formUserSelectionControl">' + 
					'<select class="form-control userSelectionCompSelectionControl"></select>' +
				'</div>' +
			'</div>'+
			'<div class="componentHoverFooter">' +
				smallClearDeleteButtonHTML("userSelectionComponentClearValueButton") + 
			'</div>' +
		'</div>';
										
	return containerHTML
}

function userSelectionTableCellContainerHTML() {
	var containerHTML = ''+
		'<div class=" layoutContainer userSelectionTableCellContainer">' +
			'<div class="form-group marginBottom0">'+
				'<div class="formUserSelectionControl">' + 
					'<select class="form-control userSelectionCompSelectionControl"></select>' +
				'</div>' +
			'</div>'+
			'<div class="componentHoverFooter">' +
				smallClearDeleteButtonHTML("userSelectionComponentClearValueButton") + 
			'</div>' +
		'</div>';
									
	return containerHTML
	
}

function setUserSelectionComponentLabel($userSelection,userSelection) {
	var $label = $userSelection.find('label')
	
	setFormComponentLabel($label,userSelection.properties.fieldID,
			userSelection.properties.labelFormat)	
	
}