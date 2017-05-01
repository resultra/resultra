
function getToggleControlFromToggleContainer($toggleContainer) {
	var $toggleControl = $toggleContainer.find(".toggleFormComponentControl")
	assert($toggleControl !== undefined, "getToggleControlFromToggleContainer: Can't get control")
	return $toggleControl
}



// Support the generation of unique IDs for each individual toggle. This isn't used
// to idenfity which field the toggle is connected to, but to connect the toggle
// to its label, so clicking on the label will check/uncheck the toggle as well.
var uniqueToggleIDForLabel = 1
function generateUniqueToggleIDForLabel() {
	uniqueToggleIDForLabel++
	return "toggleComponent_" + uniqueToggleIDForLabel
}

function toggleContainerHTML(elementID)
{	
	
	var uniqueID = generateUniqueToggleIDForLabel()
	
	var containerHTML = ''+
		'<div class=" layoutContainer toggleFormContainer">' +
			'<div class="toggle">' +
				'<input type="checkbox" id="'+uniqueID+'"class="toggleFormComponentControl">' +
				'<label for="'+  uniqueID + '"class="toggleFormComponentLabel"> New Toggle</label>' + 
			'</div>' +
			'<div class="componentHoverFooter">' +
				smallClearDeleteButtonHTML("toggleComponentClearValueButton") + 
			'</div>' +
		'</div>';
				
	console.log ("Toggle HTML: " + containerHTML)
		
	return containerHTML
}

function setToggleComponentLabel($toggleContainer,toggleRef) {
	var $label = $toggleContainer.find('.toggleFormComponentLabel')
	
	setFormComponentLabel($label,toggleRef.properties.fieldID,
			toggleRef.properties.labelFormat)	
	
}

function getCurrentToggleComponentValue($toggleContainer) {
	var $toggle = $toggleContainer.find(".toggleFormComponentControl")
	var isIndeterminate = $toggle.prop("indeterminate")
	if (isIndeterminate) {
		return null
	} else {
		var isChecked = $toggle.prop("checked")
		return isChecked
	}
}

function toggleComponentIsDisabled($toggleContainer) {
	var $toggle = $toggleContainer.find(".toggleFormComponentControl")
	var disabled = $toggle.prop("disabled")
	return disabled
	
}