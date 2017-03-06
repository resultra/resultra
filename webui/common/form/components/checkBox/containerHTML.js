
function getCheckboxControlFromCheckboxContainer($checkboxContainer) {
	var $checkboxControl = $checkboxContainer.find(".checkboxFormComponentControl")
	assert($checkboxControl !== undefined, "getCheckboxControlFromCheckboxContainer: Can't get control")
	return $checkboxControl
}




function checkBoxContainerHTML(elementID)
{	
	var containerHTML = ''+
		'<div class=" layoutContainer checkBoxFormContainer">' +
			'<div class="checkbox">' +
				'<input type="checkbox" class="checkboxFormComponentControl">' +
				'<label class="checkboxFormComponentLabel"></label>' + 
			'</div>' +
		'</div><';
				
	console.log ("Checkbox HTML: " + containerHTML)
		
	return containerHTML
}
