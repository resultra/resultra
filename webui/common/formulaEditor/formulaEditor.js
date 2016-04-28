
// Only one formula editor is expected to be initialized/configured per page, so
// a global configuration should suffice.
var formulaEditorConfig;

function initFormulaEditor(editorConfig) {
	console.log("Initializing formula editor")
	
	var editor = ace.edit("formulaEditor")
	
	// Address a console warning message on scrolling
	editor.$blockScrolling = Infinity;
	
	
	editor.setTheme("ace/theme/tomorrow_night")
	editor.setShowPrintMargin(false);
	editor.setValue("Hello World!")
	editor.setHighlightActiveLine(false);
	
	$('#formulaEditor').popup({on: 'manual'})
	
	$('#formulaEditMoreDropdown').dropdown()
	
	$('#formulaErrorMsgClose').on('click', function() {
		$('#formulaEditor').popup('hide')
    })
	
	loadFieldInfo(function(fieldsByID) {
		console.log("Initializing formula editor field insertion menu")
		initCalcFieldFieldRefSelector(fieldsByID)
	}, [fieldTypeAll])
	
	editorConfig["editor"] = editor
	formulaEditorConfig = editorConfig
	
//	editor.setTheme("ace/mode/javascript")
}

function openFormulaEditor(fieldRef) {

	formulaEditorConfig.editor.setValue("")

	formulaEditorConfig.showEditorFunc()
	
	$('#saveFormulaButton').unbind('click');
	$('#saveFormulaButton').click(function(e){
		e.preventDefault();
		console.log("save button clicked")
	});
	
	$('#checkFormulaButton').unbind('click');
	$('#checkFormulaButton').click(function(e){
		
		e.preventDefault();
		
		var validationParams = {
			fieldID: fieldRef.fieldID,
			formulaText: formulaEditorConfig.editor.getValue()
		}
		
		jsonAPIRequest("calcField/validateFormula",validationParams,function(validationResponse) {
			if(validationResponse.isValidFormula) {
				console.log("formula validation successful")
				$('#formulaEditor').popup('hide')
			} else {
				console.log("formula validation failed: " + validationResponse.errorMsg)
				$("#formulaErrorMessageMsgText").text(validationResponse.errorMsg);
				$('#formulaEditor').popup('show')
			}
		})
	});
	
}

function closeFormulaEditor() {
	$('#formulaEditor').popup('hide')
	formulaEditorConfig.editor.setValue("")
	formulaEditorConfig.hideEditorFunc()	
}

function toggleFormulaEditorForField(fieldRef) {
	if(fieldRef.fieldInfo.isCalcField) {
		openFormulaEditor(fieldRef)
	} else {
		closeFormulaEditor()
	}
}