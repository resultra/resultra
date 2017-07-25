
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
				'<label>New Text Box</label>' + componentHelpPopupButtonHTML() +
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

function initUserSelectionClearValueButton($userSelection,userSelection) {
	
	var $clearValueButton = $userSelection.find(".userSelectionComponentClearValueButton")
	
	var fieldID = userSelection.properties.fieldID
	
	function hideClearValueButton() {
		$clearValueButton.css("display","none")
	}
	
	function showClearValueButton() {
		$clearValueButton.css("display","")
	}
	
	
	var fieldRef = getFieldRef(fieldID)
	if(fieldRef.isCalcField) {
		hideClearValueButton()
		return
	}
	
	if(formComponentIsReadOnly(userSelection.properties.permissions)) {
		hideClearValueButton()
	} else {
		if(userSelection.properties.clearValueSupported) {
			showClearValueButton()
		} else {
			hideClearValueButton()
		}
	}
	
}


function initUserSelectionFormComponentContainer($container,userSelection) {
		setUserSelectionComponentLabel($container,userSelection)
		initUserSelectionClearValueButton($container,userSelection)
		initComponentHelpPopupButton($container, userSelection)
	
}