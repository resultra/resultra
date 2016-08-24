

function createPrefixedSelector(elemPrefix, elemSuffix) {
	var elemSelector = '#' + elemPrefix + elemSuffix
	return elemSelector
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

function setElemGeometry(elem, geometry) {
	elem.css({
		top: geometry.positionTop,
		left: geometry.positionLeft,
		width: geometry.sizeWidth,
		height: geometry.sizeHeight,
		position: "absolute"
	});
	
}

function initButtonClickHandler(buttonSelector,handlerFunc) {
	$(buttonSelector).unbind("click")
	$(buttonSelector).click(function(e) {
		$(this).blur(); // de-select the button after it's been clicked
	    e.preventDefault();// prevent the default  functionality
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