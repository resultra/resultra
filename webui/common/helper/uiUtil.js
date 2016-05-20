

// When the same HTML template appears on the same page, an element prefix must be used to distinguish
// the id's of the 2 DOM elements. Given a previx and suffix, this function generates both the 
// base id and selector for addressing these types of elements.
function createPrefixedTemplElemInfo(elemPrefix,elemSuffix) {
	return {
		id: elemPrefix + elemSuffix,
		selector: '#' + elemPrefix + elemSuffix
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