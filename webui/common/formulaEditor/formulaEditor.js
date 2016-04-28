
// Only one formula editor is expected to be initialized/configured per page, so
// a global configuration should suffice.
var formulaEditorConfig;

function initFormulaEditor(editorConfig) {
	console.log("Initializing formula editor")
	
	var editor = ace.edit("formulaEditor")
	editor.setTheme("ace/theme/tomorrow_night")
	editor.setShowPrintMargin(false);
	editor.setValue("Hello World!")
	
	
	$('#formulaEditMoreDropdown').dropdown()
	
	loadFieldInfo(function(fieldsByID) {
		console.log("Initializing formula editor field insertion menu")
		initCalcFieldFieldRefSelector(fieldsByID)
	}, [fieldTypeAll])
	
	editorConfig["editor"] = editor
	formulaEditorConfig = editorConfig
	
//	editor.setTheme("ace/mode/javascript")
}

function openFormulaEditor(fieldRef) {
	formulaEditorConfig.editor.setValue(fieldRef.fieldInfo.name)
	formulaEditorConfig.showEditorFunc()
}

function closeFormulaEditor() {
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