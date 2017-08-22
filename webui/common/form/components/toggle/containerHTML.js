
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
				toggleControlHTML() + componentHelpPopupButtonHTML() +
			'</div>' +
			'<div class="componentHoverFooter">' +
				smallClearDeleteButtonHTML("toggleComponentClearValueButton") + 
			'</div>' +
		'</div>';
				
	console.log ("Toggle HTML: " + containerHTML)
		
	return containerHTML
}

function toggleControlHTMLNoLabel() {
	return 		'<input type="checkbox"class="toggleFormComponentControl">'
}


function toggleTableCellContainerHTML() {
	var containerHTML = ''+
		'<div class=" layoutContainer toggleTableCellContainer">' +
			'<div class="row">' +
				'<div class="col-xs-11" style="width:auto">' +
					'<div class="toggleWrapper">' +
						toggleControlHTMLNoLabel() + 
					'</div>' +
				'</div>' +
				'<div class="col-xs-1">' +
					smallClearDeleteButtonHTML("toggleComponentClearValueButton") + 
				'</div>' +
			'</div>' +
		'</div>';
					
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
		labelWidth:5,
		 state: true,
		onColor:'success',
		offColor:'warning'
});
	
}

function initToggleComponentClearValueButton($toggleContainer,toggleRef) {
	initClearValueControl($toggleContainer,toggleRef,".toggleComponentClearValueButton")
}

function calcToggleTableCellHandleWidth(toggleRef) {

	var labelPadding = 10
	
	// The proper way to calculate the width would be to use the jQuery width() method
	// on a DOM element which has the same attributes as the labels inside the toggle.
	// However, the following heuristic works fairly well for longer labels, and just
	// find for typically short labels.
	var widthPerChar = 9
	var onWidth = toggleRef.properties.offLabel.length*widthPerChar + labelPadding
	var offWidth = toggleRef.properties.offLabel.length*widthPerChar + labelPadding

	var handleWidth = Math.max(40,onWidth,offWidth)
	
	return handleWidth
}

function calcToggleTableCellHandleWidthPx(toggleRef) {
	var handleWidthPx = calcToggleTableCellHandleWidth(toggleRef) + 'px'
	
	return handleWidthPx
}


function initToggleComponentControl($toggleContainer,toggleRef) {
		
	var $toggleControl = getToggleControlFromToggleContainer($toggleContainer)
		
		
	 $toggleControl.bootstrapSwitch({
		handleWidth:calcToggleTableCellHandleWidthPx(toggleRef),
		indeterminate:true,
		 size:'normal',
		onText:escapeHTML(toggleRef.properties.onLabel),
		 offText:escapeHTML(toggleRef.properties.offLabel),
		labelWidth:'5px',
		 animate:true,
		onColor:toggleRef.properties.onColorScheme,
		offColor:toggleRef.properties.offColorScheme
	});
	
}

function initToggleComponentFormComponentContainer($toggleContainer,toggleRef) {
	initToggleComponentControl($toggleContainer,toggleRef)
	setToggleComponentLabel($toggleContainer,toggleRef)
	initToggleComponentClearValueButton($toggleContainer,toggleRef)
	initComponentHelpPopupButton($toggleContainer, toggleRef)
}


function initToggleComponentTableViewComponentContainer($toggleContainer,toggleRef) {
	initToggleComponentControl($toggleContainer,toggleRef)
	initToggleComponentClearValueButton($toggleContainer,toggleRef)

	var toggleColorSchemeClass = "checkbox-"+toggleRef.properties.colorScheme
	$toggleContainer.addClass(toggleColorSchemeClass)

	/*  Bootstrap switch uses negative margins and has invisible parts of the toggle overrun
	   100% of the parent's size. By allowing the size to be larger than the cell, the
	   toggle will not unnessessarily lessen it's width and expand the height of the table cell. 
	
	   See the following for more background on this: https://github.com/Bttstrp/bootstrap-switch/issues/419
		
		Using a negative margin on the RHS of the toggle control allows the hidden parts of the 
	   control to expand to their full width, without compressing the column. Otherwise, the
	   width of the column will compress and expand the height of the toggle control. 
	
		The DataTable table control also respect min-width when sizing the table columns. In this case,
		however, the min-width also needs to encompase the negative margin on the RHS of the table cell.
	*/
	var handleWidth = calcToggleTableCellHandleWidth(toggleRef)
	var negRightMarginWidth = handleWidth
	var widthPaddingForClearButtonAndControlHandle = 80
	var minWidthPx = handleWidth + negRightMarginWidth + widthPaddingForClearButtonAndControlHandle + 'px'
	$toggleContainer.css("min-width",minWidthPx)
	$toggleContainer.css("margin-right",(-1 * negRightMarginWidth)+'px')
	
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

