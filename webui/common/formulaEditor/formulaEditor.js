function initFormulaEditor() {
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
	
//	editor.setTheme("ace/mode/javascript")
}