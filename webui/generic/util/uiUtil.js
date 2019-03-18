// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.


function createPrefixedSelector(elemPrefix, elemSuffix) {
	var elemSelector = '#' + elemPrefix + elemSuffix
	return elemSelector
}

function createPrefixedContainerObj(elemPrefix, elemSuffix) {
	return $(createPrefixedSelector(elemPrefix,elemSuffix))
}

// When the same HTML template appears on the same page, an element prefix must be used to distinguish
// the id's of the 2 DOM elements. Given a previx and suffix, this function generates both the 
// base id and selector for addressing these types of elements.
function createPrefixedTemplElemInfo(elemPrefix,elemSuffix) {
	
	var elemSelector = '#' + elemPrefix + elemSuffix
	return {
		id: elemPrefix + elemSuffix,
		selector: elemSelector,
		
		// Convenience method for the typical case when creating references
		// to form controls. Provides a short-hand notation to retrieve the form control's value.
		val: function () {
			return $(elemSelector).val()
		}
	}
}

function createIDWithSelector(elemID) {
	return {
		id: elemID,
		selector: '#' + elemID
	}
	
}

function headerWithBodyHTML(header, body) {
	return '<div class="header">' + header  + '</div>' + body
}



function getFormStringValue(formID,fieldID) {
	console.log("getFormStringValue: " + formID + " - " + fieldID + ": val= " + 
			$(formID).form('get value',fieldID))
	return $(formID).form('get value',fieldID).toString()
}

function hideSiblingsShowOne(elemID) {
	$(elemID).siblings().hide()
	$(elemID).show()
}

function getFormFloatValue(formID,fieldID) {
	console.log("getFormFloatValue: "  + formID + " - " + fieldID + ": val= " + 
		$(formID).form('get value',fieldID))
	return parseFloat($(formID).form('get value',fieldID))
}

function setElemDimensions(elem, geometry) {
	elem.css({
		width: geometry.sizeWidth,
		height: geometry.sizeHeight,
		position: "relative"
	});
}

function setElemFixedWidthFlexibleHeight($elem,sizeWidth) {
	$elem.css({
		width: sizeWidth,
		height: "auto",
		position: "relative"
	});
	
}

function setElemFixedWidthMaxHeight($elem,geometry) {
	setElemFixedWidthFlexibleHeight($elem,geometry.sizeWidth)
	var maxHeightPx = geometry.sizeHeight + "px"
	$elem.css('max-height',maxHeightPx)
}

function initButtonClickHandler(buttonSelector,handlerFunc) {
	$(buttonSelector).unbind("click")
	$(buttonSelector).click(function(e) {
		$(this).blur(); // de-select the button after it's been clicked
	    e.preventDefault() // prevent the default  functionality
		e.stopPropagation() // In general, there's no need to propagate button clicks to their parent elements
		handlerFunc()
	})
}

function initButtonControlClickHandler($button,handlerFunc) {
	$button.unbind("click")
	$button.click(function(e) {
		$(this).blur(); // de-select the button after it's been clicked
	    e.preventDefault();// prevent the default  functionality
		e.stopPropagation() // In general, there's no need to propagate button clicks to their parent elements
		handlerFunc()
	})
}

function disableButton(buttonSelector) {
	$(buttonSelector).prop('disabled', true);
}

function enableButton(buttonSelector) {
	$(buttonSelector).prop('disabled', false);
}


function initSelectionChangedHandler(selectionSelector, handlerFunc) {
	$(selectionSelector).unbind("change")				
	$(selectionSelector).change(function(){
		var newValue = $(selectionSelector).val()
		handlerFunc(newValue)
	})
	
}

function elemIsDisplayed($elem) {
	if ($elem.css('display') == 'none') { 
		return false
	} else {
		return true
	}
}

function initSelectControlChangeHandler($selectControl, handlerFunc) {
	$selectControl.unbind("change")				
	$selectControl.change(function(){
		var newValue = $selectControl.val()
		handlerFunc(newValue)
	})
	
}

function initNumberSelectionChangeHandler($selectionControl, handlerFunc) {
	initSelectControlChangeHandler($selectionControl,function(selectedVal) {
		var numberVal = Number(selectedVal)
		handlerFunc(numberVal)
	})
}

function initCheckboxControlChangeHandler($checkbox, initialVal, handlerFunc) {
	$checkbox.prop("checked",initialVal)
	$checkbox.unbind("click")
	$checkbox.click( function () {
		var newVal = $checkbox.prop("checked")
		handlerFunc(newVal)
	})	
	
}

function initCheckboxChangeHandler(checkboxSelector, initialVal, handlerFunc) {
	var $checkbox = $(checkboxSelector)
	initCheckboxControlChangeHandler($checkbox,initialVal, handlerFunc)
}

function calcContrainedPxVal(val, minVal,maxVal) {
	var roundedVal = Math.round(val)
	if(roundedVal > maxVal) {
		return maxVal
	} else if (roundedVal < minVal) {
		return minVal
	} else {
		return roundedVal
	}
}

// Function to dynamically return the width of the text: per the following:
// https://stackoverflow.com/questions/118241/calculate-text-width-with-javascript
function calcTextWidth(text) {
    var canvas = calcTextWidth.canvas || (calcTextWidth.canvas = document.createElement("canvas"));
    var context = canvas.getContext("2d");
	var font = "12pt arial"
    context.font = font;
    var metrics = context.measureText(text);
    return metrics.width;
}

function calcConstrainedTextWidth(text,minWidth, maxWidth) {
	var unconstrainedWidth = calcTextWidth(text)
	return calcContrainedPxVal(unconstrainedWidth,minWidth,maxWidth)
}


function insertTextAreaAtCursor(elem, newText) {
	
  console.log("appending text to formula box: " + newText)
	
  var selStart = elem.prop("selectionStart")
  var selEnd = elem.prop("selectionEnd")
  var allText = elem.val()
  var beforeSel = allText.substring(0, selStart)
  var afterSel  = allText.substring(selEnd, allText.length)
  elem.val(beforeSel + newText + afterSel)
  elem[0].selectionStart = elem[0].selectionEnd = selStart + newText.length
  elem.focus()
}


var escapeHTMLEntityMap = {
  "&": "&amp;",
  "<": "&lt;",
  ">": "&gt;",
  '"': '&quot;',
  "'": '&#39;',
  "/": '&#x2F;'
};

// See: http://stackoverflow.com/questions/24816/escaping-html-strings-with-jquery
function escapeHTML(string) {
  return String(string).replace(/[&<>"'\/]/g, function (s) {
    return escapeHTMLEntityMap[s];
  });
}

function openNewWindowWithElectronOptions(url) {
	// When the app is run within Electron, window.open() actually returns
	// a BrowserWindowPropxy() object instead of the standard Javascript object.
	// This object will inherit the window properties of the main window.
	//
	// 		(see https://electronjs.org/docs/api/window-open)
	//
	// By default, the way the main window in Electron works with the splashcreen,
    // and to avoid flashing upon open, is to have it's initial options set to 
	// not be visible, then show it once the content is initialized.
	//
	// To make the window actually appear, the "show=true" argument has
	// to be passed to the window.open() function. This overrides the parent
	// windows options appropriately.
	//
	// TBD - In the browser version, this function is always opening a new window, rather than a
	// tab. This may be problematic if popups are blocked.
	// 
	var win = window.open(url,"_blank","show=true")
	win.focus()
	
}