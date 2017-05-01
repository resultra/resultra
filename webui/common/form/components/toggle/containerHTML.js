
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

function toggleControlHTML() {
	var uniqueID = generateUniqueToggleIDForLabel()
	return 		'<input type="checkbox" id="'+uniqueID+'"class="toggleFormComponentControl">' +
		'<label for="'+  uniqueID + '"class="toggleFormComponentLabel"> New Toggle</label>'
}

function toggleContainerHTML(elementID)
{	
	var containerHTML = ''+
		'<div class=" layoutContainer toggleFormContainer">' +
			'<div class="toggleWrapper">' +
				toggleControlHTML() + 
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

function initDummyToggleControlForDragAndDrop($dummyToggleControlForDragAndDrop) {
	var $toggleControl = getToggleControlFromToggleContainer($dummyToggleControlForDragAndDrop)

	 $toggleControl.bootstrapSwitch({
		handleWidth:'40px',
		onText:'Yes',
		offText:'No',
		labelWidth:5 ,
		 state: true,
		onColor:'success',
		offColor:'warning'
});
	
}


function initToggleComponentControl($toggleContainer,toggleRef) {
	
	var $toggleControl = getToggleControlFromToggleContainer($toggleContainer)

	 $toggleControl.bootstrapSwitch({
		handleWidth:'40px',
		indeterminate:true,
		onText:escapeHTML(toggleRef.properties.onLabel),
		 offText:escapeHTML(toggleRef.properties.offLabel),
		labelWidth:'5px',
		 animate:true,
		onColor:toggleRef.properties.onColorScheme,
		offColor:toggleRef.properties.offColorScheme
	});
	
	
		 // TODO - Dynamically size the handleWidth based upon the length of the on
	// and off labels: e.g.:	
// 	$toggleControl.bootstrapSwitch('handleWidth','80px')

}


function reInitToggleComponentControl($toggleContainer,toggleRef) {
	
	// When manipulating the toggle in the form designer, the control may change
	// colors or labels. Using the toggles 'destroy' method leaves the control
	// inoperaable. However, clearing out and re-initializing the control's DOM
	// elements works.
	var $toggleWrapper = $toggleContainer.find(".toggleWrapper")
	$toggleWrapper.empty()
	$toggleWrapper.append(toggleControlHTML)
	
	initToggleComponentControl($toggleContainer,toggleRef)
	setToggleComponentLabel($toggleContainer,toggleRef)
	
	
}


function getCurrentToggleComponentValue($toggleContainer) {
	var $toggle = $toggleContainer.find(".toggleFormComponentControl")
	var isIndeterminate = $toggle.bootstrapSwitch("indeterminate")
	if (isIndeterminate) {
		return null
	} else {
		var isChecked = $toggle.bootstrapSwitch("state")
		return isChecked
	}
}

function toggleComponentIsDisabled($toggleContainer) {
	var $toggle = $toggleContainer.find(".toggleFormComponentControl")
	var disabled = $toggle.prop("disabled")
	return disabled
}

