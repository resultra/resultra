// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function textBoxContainerInputControl() {
	return '<div class="input-group">'+
					'<input type="text" name="symbol" class="textBoxComponentInput form-control" placeholder="">'+
					'<div class="input-group-btn valueSelectionDropdown">'+
						'<button type="button" class="btn btn-default dropdown-toggle" ' + 
								'data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">' +
								'<span class="caret"></span></button>'+
						'<ul class="dropdown-menu valueDropdownMenu">' +
						'</ul>'+
					'</div>'+
					clearValueButtonHTML("textBoxComponentClearValueButton") +
				'</div>'
}

function textBoxContainerHTML(elementID)
{
	var containerHTML = ''+
		'<div class="layoutContainer textBoxComponent textBoxFormComponent">' +
				'<label>New Text Box</label>'+ componentHelpPopupButtonHTML() +
				textBoxContainerInputControl() +
		'</div>';
	return containerHTML
}

function textBoxTableViewContainerHTML() {
	var containerHTML = ''+
		'<div class="layoutContainer textBoxComponent textBoxTableCellComponent">' +
			textBoxContainerInputControl() +
		'</div>';
	return containerHTML
}


function setTextBoxComponentLabel($textBoxContainer, textBoxRef) {

	var $label = $textBoxContainer.find('label')
	
	setFormComponentLabel($label,textBoxRef.properties.fieldID,
			textBoxRef.properties.labelFormat)	
}

function initTextBoxClearValueControl($textBoxContainer, textBoxRef) {
	initClearValueControl($textBoxContainer,textBoxRef,".textBoxComponentClearValueButton")	
}

function initTextBoxFormComponentContainer($container,textBoxRef) {
	setTextBoxComponentLabel($container,textBoxRef)
	function dummySetVal(dropdownVal) {}
	configureTextBoxComponentValueListDropdown($container, textBoxRef,dummySetVal)
	initTextBoxClearValueControl($container, textBoxRef)
	initComponentHelpPopupButton($container, textBoxRef)
}



function configureTextBoxComponentValueListDropdown($textBoxContainer, textBoxRef, setValueCallback) {
	
	var $valSelectionDropdown = $textBoxContainer.find(".valueSelectionDropdown")
	var textBoxFieldID = textBoxRef.properties.fieldID
	
	function hideDropdownControls() {
		$valSelectionDropdown.css("display","none")
	}
	
	function showDropdownControls() {
		// The jQuery show() method will set the display to "block", which causes the controls to display on a
		// new line.
		$valSelectionDropdown.css("display","")	
	}
	
	var fieldRef = getFieldRef(textBoxFieldID)
	if(fieldRef.isCalcField) {
		hideDropdownControls()
		return
	}
	
	if(formComponentIsReadOnly(textBoxRef.properties.permissions)) {
		hideDropdownControls()
		return		
	}
	
	function createValueDropdownMenuItem(valueText) {
		var $menuItem = $('<li><a href="#"></a></li>')
		var $menuLink = $menuItem.find('a')
		$menuLink.click(function(e) {
			setValueCallback($menuLink.text())
			e.preventDefault()
		})
		$menuItem.find('a').text(valueText)
		return $menuItem
	}
	
	
	var valueListID = textBoxRef.properties.valueListID
	if (valueListID !== undefined && valueListID !== null) {
		$valSelectionDropdown.show()
		var getValListParams = { valueListID: valueListID }
		jsonAPIRequest("valueList/get",getValListParams,function(valListInfo) {
			console.log("Initializing text box with value list info: " + JSON.stringify(valListInfo))
			var values = valListInfo.properties.values
			if (values.length <= 0) {
				hideDropdownControls()
				return	
			} else {
				showDropdownControls()
				var $valDropdownMenu = $textBoxContainer.find('.valueDropdownMenu')
				$valDropdownMenu.empty()
				var $header = $('<li class="dropdown-header"></li>')
				$header.text(valListInfo.name)
				$valDropdownMenu.append($header)
				for(var valIndex = 0; valIndex < values.length; valIndex++) {
					var val = values[valIndex]
					$valDropdownMenu.append(createValueDropdownMenuItem(val.textValue))
				}		
				
			}
		})
	} else {
		hideDropdownControls()
		return		
	}
	
	showDropdownControls()
	

}