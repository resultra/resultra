
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

function checkBoxContainerControlHTML() {
	
}

function checkBoxContainerHTML(elementID)
{	
	
	var uniqueID = generateUniqueCheckboxIDForLabel()
	
	var containerHTML = ''+
		'<div class=" layoutContainer checkBoxContainer checkBoxFormContainer">' +
			'<div class="checkbox">' +
				'<input type="checkbox" id="'+uniqueID+'" class="checkboxFormComponentControl">' +
				'<label for="'+  uniqueID + '"class="checkboxFormComponentLabel">New Checkbox</label>' + 
			'</div>' +
			'<div class="componentHoverFooter">' +
				smallClearDeleteButtonHTML("checkBoxComponentClearValueButton") + 
			'</div>' +
		'</div>';
				
	console.log ("Checkbox HTML: " + containerHTML)
		
	return containerHTML
}

function checkBoxTableViewCellContainerHTML() {
	
	var uniqueID = generateUniqueCheckboxIDForLabel()
	
	return 	''+
		'<div class=" layoutContainer checkBoxContainer checkBoxTableCellContainer">' +
			'<div class="checkbox">' +
				'<input type="checkbox" id="'+uniqueID+'" class="checkboxTableCellControl checkboxFormComponentControl">' +
				'<label for="'+  uniqueID + '"class="checkboxFormComponentLabel"></label>'
			'</div>' +
			'<div class="componentHoverFooter">' +
				smallClearDeleteButtonHTML("checkBoxComponentClearValueButton") + 
			'</div>'
		'</div>';

}

function initCheckBoxControl($checkbox,checkBoxObjectRef) {
	var checkboxColorSchemeClass = "checkbox-"+checkBoxObjectRef.properties.colorScheme
	$checkbox.addClass(checkboxColorSchemeClass)
	
}

function setCheckBoxComponentLabel($checkboxContainer,checkboxRef) {
	var $label = $checkboxContainer.find('.checkboxFormComponentLabel')
	
	setFormComponentLabel($label,checkboxRef.properties.fieldID,
			checkboxRef.properties.labelFormat)	
	
}



function initCheckBoxClearValueControl($checkboxContainer,checkboxRef) {
	initClearValueControl($checkboxContainer,checkboxRef,".checkBoxComponentClearValueButton")
}


function getCurrentCheckboxComponentValue($checkboxContainer) {
	var $checkbox = $checkboxContainer.find(".checkboxFormComponentControl")
	var isIndeterminate = $checkbox.prop("indeterminate")
	if (isIndeterminate) {
		return null
	} else {
		var isChecked = $checkbox.prop("checked")
		return isChecked
	}
}

function checkboxComponentIsDisabled($checkboxContainer) {
	var $checkbox = $checkboxContainer.find(".checkboxFormComponentControl")
	var disabled = $checkbox.prop("disabled")
	return disabled
	
}