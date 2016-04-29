



function calcFieldAppendFormulaText(newText) {
	insertTextAreaAtCursor($('#calcFieldFormulaTextArea'), newText)
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
		
		var formulaFieldRef = '[' + fieldInfo.refName + ']'

	 	var selectFieldRefHTML = dropdownSelectItemHTML(
						formulaFieldRef,fieldInfo.name)
		
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
	
	$("#calcFieldFuncSelectionMenu").append(dropdownSelectItemHTML("SUM","SUM(value1,value2,...)"))
		
	$('#calcFieldInsertSelectedFuncButton').click(function(e){
		e.preventDefault();
		
		var funcName = $("#calcFieldFuncSelectionDropdown").dropdown("get value")
			 
		console.log("calcFieldInsertSelectedFieldRefButton: button clicked: " + funcName )
			 
		calcFieldAppendFormulaText(funcName) 
	});
}

function clearCalcFieldValidationMsgs() {
	$('#calcFieldErrorMsgBox').empty()
	$('#calcFieldSuccessMsgBox').empty()
}

function calcFieldSetValidationErrorMsg(errorMsg) {
	clearCalcFieldValidationMsgs()
	$('#calcFieldErrorMsgBox').append('<p><i class="warning sign icon"></i>' + errorMsg + "</p>")	
}

function calcFieldSetValidationSuccessMsg() {
	clearCalcFieldValidationMsgs()
	$('#calcFieldSuccessMsgBox').append('<p><i class="checkmark icon"></i>Formula is valid</p>')
}

function initCalcFieldFormulaTextBox() {
		
	$('#calcFieldFormulaTextArea').focusout(function () {
		
		console.log("Calculated field formula box losing focus: re-validating formula")
		
		var validationParams = {
			eqnText: $('#calcFieldFormulaTextArea').val(),
			isNewField: true
		}
		jsonAPIRequest("validateCalcFieldEqn",validationParams,function(validationResponse) {
			if(validationResponse.isValidEqn) {
				calcFieldSetValidationSuccessMsg()
			} else {
				calcFieldSetValidationErrorMsg(validationResponse.errorMsg)
			}
		})
		
	})
	
}


function initCalcFieldEditBehavior(fieldsByID) {
	
	clearCalcFieldValidationMsgs()
	initCalcFieldFormulaTextBox()
	
	initCalcFieldFieldRefSelector(fieldsByID)
	initCalcFieldFuncSelector()
	
}