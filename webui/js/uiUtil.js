

function dropdownSelectItemHTML(selItemVal, selItemText)
{
	var selectFieldRefHTML = '<div class="item" data-value="' +
	 		selItemVal + '">' +
	 		selItemText + '</div>'
	return selectFieldRefHTML
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