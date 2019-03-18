// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

function getToggleControlFromToggleContainer($toggleContainer) {
	var $toggleControl = $toggleContainer.find(".toggleFormComponentControl")
	assert($toggleControl !== undefined, "getToggleControlFromToggleContainer: Can't get control")
	return $toggleControl
}




function toggleControlHTML() {
	return '<input type="checkbox" class="toggleFormComponentControl">'		
}

function toggleContainerHTML(elementID)
{	
	var containerHTML = ''+
		'<div class=" layoutContainer toggleFormContainer">' +
			'<div class="container-fluid componentHeader">' + 
				'<div class="row">' +
					'<div class="col-xs-9 componentHeaderLabelCol">' +
						'<label class="toggleFormComponentLabel marginBottom0">New Toggle</label>' +
					'</div>' +
					'<div class="col-xs-3 componentHeaderButtonCol">' +
						smallClearComponentValHeaderButton("toggleComponentClearValueButton") + 
						componentHelpPopupButtonHTML() +
					'</div>' +
				'</div>' +
			'</div>' +
			'<div class="toggleWrapper">' +
				toggleControlHTML() + 
			'</div>' +
		'</div>';
						
	return containerHTML
}


function toggleTableCellContainerHTML() {
	var containerHTML = ''+
		'<div class=" layoutContainer toggleTableCellContainer">' +
			'<div class="row">' +
				'<div class="col-xs-11" style="width:auto">' +
					'<div class="toggleWrapper">' +
						toggleControlHTML() + 
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


function initToggleComponentControl($toggleContainer,toggleRef,handleWidth) {
		
	var $toggleControl = getToggleControlFromToggleContainer($toggleContainer)
				
	 $toggleControl.bootstrapSwitch({
		handleWidth:handleWidth,
		indeterminate:true,
		 size:'normal',
		onText:escapeHTML(toggleRef.properties.onLabel),
		 offText:escapeHTML(toggleRef.properties.offLabel),
		labelWidth:'5px',
		 animate:false,
		onColor:toggleRef.properties.onColorScheme,
		offColor:toggleRef.properties.offColorScheme
	});
	
}

function initToggleFormComponentControl($toggleContainer,toggleRef) {
	
	// We want to the switch to extend to the size of the 
	// $toggleContainer.
	// 
	// The Bootstrap Switch control allocates 12px of padding on 
	// both the LHS and RHS of each handle and the label.
	//
	// At any given time, only a single handle and the label
	// is visible: i.e.:
	//
	// (12px + handleWidth + 12px) + (12px + labelWidth + 12px)
	// So assuming the labelWidth is 5px, the total non-handle
	// space is 4*12+5 = 53px.
	//
	// There is also a total of 12px padding and border
	// on the LHS and RHS $toggleContainer.
	//
	// So, to calculate the handleWidth for the Bootstrap Switch
	// (around which all the other calculations take place), 65
	// is subtracted from the size of the container.
	//
	// Finally, for aesthetic reasons the toggle control looks 
	// better if the RHS isn't flush with the RHS of the $toggleContainer's
	// perimeter. For this reason, an extra 5px is added to the total
	// "non handleWidth" space.
	// 	
	var nonHandleWidthFixedSpaceForPadding = 70
	var handleWidth = (toggleRef.properties.geometry.sizeWidth - 
				nonHandleWidthFixedSpaceForPadding) + 'px'
	
	// Similar to the negative margin used with the table cell, negative
    // margin is needed when the toggle is displayed in the form to give
	// the toggle room to expand beyond the length of the $toggleContainer.
    // For some reason, the Bootstrap Switch calculates enough negative
	// margin on the LHS but not the RHS.
	var negMargin = (-1 * toggleRef.properties.geometry.sizeWidth) + 'px'
	$toggleContainer.find(".toggleWrapper").css("margin-right",negMargin)
	
//	handleWidth = calcToggleTableCellHandleWidthPx(toggleRef)
	initToggleComponentControl($toggleContainer,toggleRef,handleWidth)
}

function initToggleTableComponentControl($toggleContainer,toggleRef) {
	var handleWidth = calcToggleTableCellHandleWidthPx(toggleRef)
	initToggleComponentControl($toggleContainer,toggleRef,handleWidth)
}

function initToggleComponentFormComponentContainer($toggleContainer,toggleRef) {
	
	initToggleFormComponentControl($toggleContainer,toggleRef)
	
	setToggleComponentLabel($toggleContainer,toggleRef)
	initToggleComponentClearValueButton($toggleContainer,toggleRef)
	initComponentHelpPopupButton($toggleContainer, toggleRef)
}

function calcToggleFormControlAddonControlsWidth(toggleRef) {
	
	if (clearValueControlIsEnabled(toggleRef)) {
		return 30
	} else {
		return 0
	}	
}

function calcToggleFormControlMinWidth(toggleRef) {
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
	var addOnWidth = calcToggleFormControlAddonControlsWidth(toggleRef)
	var minWidth =  handleWidth + negRightMarginWidth + addOnWidth
	return minWidth
}

function calcToggleFormControlMinTableCellWidth(toggleRef) {
	var handleWidth = calcToggleTableCellHandleWidth(toggleRef)
	var toggleWidth = 29
	var paddingWidth = 24
	var addOnWidth = calcToggleFormControlAddonControlsWidth(toggleRef)
	var minWidth = handleWidth + toggleWidth + addOnWidth + paddingWidth
	return minWidth
}

function initToggleComponentTableViewComponentContainer($toggleContainer,toggleRef) {
	
	initToggleTableComponentControl($toggleContainer,toggleRef)
	
	initToggleComponentClearValueButton($toggleContainer,toggleRef)

	var toggleColorSchemeClass = "checkbox-"+toggleRef.properties.colorScheme
	$toggleContainer.addClass(toggleColorSchemeClass)

	var handleWidth = calcToggleTableCellHandleWidth(toggleRef)
	var negRightMarginWidth = handleWidth
	var widthPaddingForClearButtonAndControlHandle = 80
	
	var minWidthPx = calcToggleFormControlMinWidth(toggleRef) + 'px'
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
	$toggleWrapper.append(toggleControlHTML())
	
	initToggleFormComponentControl($toggleContainer,toggleRef)
	
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

