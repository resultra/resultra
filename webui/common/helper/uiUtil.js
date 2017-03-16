

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

function setElemFixedWidthFlexibleHeight(elem,sizeWidth) {
	elem.css({
		width: sizeWidth,
		height: "auto",
		position: "relative"
	});
	
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


function initCheckboxChangeHandler(checkboxSelector, initialVal, handlerFunc) {
	var $checkbox = $(checkboxSelector)
	$checkbox.prop("checked",initialVal)
	$checkbox.unbind("click")
	$checkbox.click( function () {
		var newVal = $checkbox.prop("checked")
		handlerFunc(newVal)
	})	
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