

function insertTextAreaAtCursor(elem, newText) {
  var selStart = elem.prop("selectionStart")
  var selEnd = elem.prop("selectionEnd")
  var allText = elem.val()
  var beforeSel = allText.substring(0, selStart)
  var afterSel  = allText.substring(selEnd, allText.length)
  elem.val(beforeSel + newText + afterSel)
  elem[0].selectionStart = elem[0].selectionEnd = selStart + newText.length
  elem.focus()
}

function calcFieldAppendFormulaText(newText) {
	insertTextAreaAtCursor($('#calcFieldFormulaTextArea'), newText)
}

function calcFieldInsertDropdownSelectItemHTML(selItemVal, selItemText)
{
	var selectFieldRefHTML = '<div class="item" data-value="' +
	 		selItemVal + '">' +
	 		selItemText + '</div>'
	return selectFieldRefHTML
}

function initCalcFieldFieldRefSelector(fieldsByID)
{
	$("#calcFieldFieldRefDropdown").dropdown()
	$("#calcFieldFieldRefDropdownMenu").empty()
	
	// Populate the menu to insert field references with the list of fields
	
	for (var fieldID in fieldsByID) {
		
		fieldInfo = fieldsByID[fieldID]
		
		console.log("initCalcFieldEditBehavior: populating menu to insert field ref: " +
			"field id=" + fieldID +
			" field ref=" + fieldInfo.refName +
			" name = " + fieldInfo.name
		)

	 	var selectFieldRefHTML = calcFieldInsertDropdownSelectItemHTML(
						fieldInfo.refName,
						fieldInfo.refName + ' - ' + fieldInfo.name)
		
	 	$("#calcFieldFieldRefDropdownMenu").append(selectFieldRefHTML)			

	} // for each text field
	
	$('#calcFieldInsertSelectedFieldRefButton').click(function(e){
		e.preventDefault();

		var fieldRef = $("#calcFieldFieldRefDropdown").dropdown("get value")
			 
		console.log("calcFieldInsertSelectedFieldRefButton: button clicked: " + fieldRef )
			 
		calcFieldAppendFormulaText(fieldRef)
			 
	});
	
}

function initCalcFieldFuncSelector()
{
	$("#calcFieldFuncSelectionDropdown").dropdown()
	$("#calcFieldFuncSelectionMenu").empty()
	
	// Populate the menu to insert function names into the formula editing area
	
	$("#calcFieldFuncSelectionMenu").append(calcFieldInsertDropdownSelectItemHTML("SUM()","SUM(value1,value2,...)"))
		
	$('#calcFieldInsertSelectedFuncButton').click(function(e){
		e.preventDefault();
		
		var funcName = $("#calcFieldFuncSelectionDropdown").dropdown("get value")
			 
		console.log("calcFieldInsertSelectedFieldRefButton: button clicked: " + funcName )
			 
		calcFieldAppendFormulaText(funcName) 
	});
}


function initCalcFieldEditBehavior(fieldsByID) {
	
	initCalcFieldFieldRefSelector(fieldsByID)
	initCalcFieldFuncSelector()
}