
function getCheckboxControlFromCheckboxContainer($checkboxContainer) {
	var $checkboxControl = $checkboxContainer.find(".checkboxFormComponentControl")
	assert($checkboxControl !== undefined, "getCheckboxControlFromCheckboxContainer: Can't get control")
	return $checkboxControl
}



// Support the generation of unique IDs for each individual checkbox. This isn't used
// to idenfity which field the checkbox is connected to, but to connect the checkbox
// to its label, so clicking on the label will check/uncheck the checkbox as well.
var uniqueCheckboxIDForLabel = 1
function generateUniqueCheckboxIDForLabel() {
	uniqueCheckboxIDForLabel++
	return "checkboxComponent_" + uniqueCheckboxIDForLabel
}

function checkBoxContainerHTML(elementID)
{	
	
	var uniqueID = generateUniqueCheckboxIDForLabel()
	
	var containerHTML = ''+
		'<div class=" layoutContainer checkBoxFormContainer">' +
			'<div class="checkbox">' +
				'<input type="checkbox" id="'+uniqueID+'"class="checkboxFormComponentControl">' +
				'<label for="'+  uniqueID + '"class="checkboxFormComponentLabel">New Checkbox</label>' + 
			'</div>' +
		'</div>';
				
	console.log ("Checkbox HTML: " + containerHTML)
		
	return containerHTML
}

function setCheckBoxComponentLabel($checkboxContainer,checkboxRef) {
	var $label = $checkboxContainer.find('.checkboxFormComponentLabel')
	
	setFormComponentLabel($label,checkboxRef.properties.fieldID,
			checkboxRef.properties.labelFormat)	
	
}