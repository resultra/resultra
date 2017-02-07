
function getCheckboxControlFromCheckboxContainer($checkboxContainer) {
	var $checkboxControl = $checkboxContainer.find(".checkboxFormComponentControl")
	assert($checkboxControl !== undefined, "getCheckboxControlFromCheckboxContainer: Can't get control")
	return $checkboxControl
}




function checkBoxContainerHTML(elementID)
{	
	var containerHTML = ''+
		'<div class="ui-widget-content layoutContainer checkBoxFormContainer draggable resizable">' +
			'<div class="checkbox">' +
				'<label>' + 
				  		'<input type="checkbox" class="checkboxFormComponentControl"></input><span>Checkbox Label</span> ' +
				'</label>' +
			'</div>' +
		'</div><';
				
	console.log ("Checkbox HTML: " + containerHTML)
		
	return containerHTML
}
