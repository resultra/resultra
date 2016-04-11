

function dropdownSelectItemHTML(selItemVal, selItemText)
{
	var selectFieldRefHTML = '<div class="item" data-value="' +
	 		selItemVal + '">' +
	 		selItemText + '</div>'
	return selectFieldRefHTML
}

function selectOptionHTML(selItemVal, selItemText) {
	
	var selOptionHTML = '<option value="' + selItemVal + '">' + selItemText + '</option>'
	return selOptionHTML
	
}

function itemDivHTML(itemBody) {
	return '<div class="item">' + itemBody + '</div>'
}

function contentHTML(contentBody) {
	return	'<div class="left floated content">' + contentBody +
			'</div>'
	
}

function emptyOptionHTML(prompt) {
	return '<option value="">' + prompt + '</option>'	
}

function headerWithBodyHTML(header, body) {
	return '<div class="header">' + header  + '</div>' + body
}

function nonEmptyFieldValidation(prompt) {
	return { rules: [
		            {
		              type   : 'empty',
		              prompt : prompt
		            }
		          ]
		      }
}

function validNumberFieldValidation() {
	return { 
		rules: [
			{
				type   : 'number',
				prompt : 'Enter a number'
			},
			{
				type   : 'empty',
				prompt : 'Enter a number'
			}
		]
	}
}

function validPositiveNumberFieldValidation() {
	return { 
		rules: [
			{
				type   : 'empty',
				prompt : 'Enter a number'
			},
			{
				type : 'regExp[/(^[0][.]{1}[0-9]{0,}[1-9]+[0-9]{0,}$)|(^[1-9]+[0-9]{0,}[.]?[0-9]{0,}$)/]',
				prompt : 'Enter a positive number'
			}
		]
	}
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