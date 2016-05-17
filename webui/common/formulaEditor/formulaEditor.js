
// Only one formula editor is expected to be initialized/configured per page, so
// a global configuration should suffice.
var formulaEditorConfig;

function populateFieldRefInsertionMenu(tableID)
{
	// Populate the menu to insert field references with the list of fields
	$("#formulaFieldRefSelector").empty()
	
	$("#formulaFieldRefSelector").append('<option value="" disabled selected>Insert Field Reference</option>')
	loadFieldInfo(tableID, [fieldTypeAll],function(fieldsByID) {
		for (var fieldID in fieldsByID) {
		
			var fieldInfo = fieldsByID[fieldID]		
		
     	   var menuItemHTML = '<option value="' + fieldInfo.refName + 
				'">' + fieldInfo.name + '</option>'
			
			console.log("Adding selection to insert formula menu:" + menuItemHTML)
				
		 	$("#formulaFieldRefSelector").append(menuItemHTML)			

		} // for each  field
		
		$("#formulaFieldRefSelector").on('change',function() {
			var fieldRefName = $(this).find("option:selected").val();
			if(fieldRefName.length > 0) {
				formulaEditorConfig.editor.insert("[" + fieldRefName + "]")	
				$('#formulaFieldRefSelector').prop('selectedIndex',0);
			}
		})
		
	})
	
}

function initFormulaEditor(editorConfig) {
	console.log("Initializing formula editor")
	
	var editor = ace.edit("formulaEditor")
	editorConfig["editor"] = editor
	formulaEditorConfig = editorConfig
	
	// Address a console warning message on scrolling
	editor.$blockScrolling = Infinity;
	editor.setTheme("ace/theme/tomorrow_night")
	editor.setShowPrintMargin(false);
	editor.setValue("Hello World!")
	editor.setHighlightActiveLine(false);
	
	
	$('#formulaEditorErrorPopup').popover({placement:'top',trigger:'manual'})
		
	populateFieldRefInsertionMenu(editorConfig.tableID)
	
	
	
	// TODO - Setup the editor for language specific syntax highlighting, etc.
}

function validateFormula(fieldRef,validationSucceededCallback) {
	
	var formulaText =  formulaEditorConfig.editor.getValue()
	
	var validationParams = {
		fieldID: fieldRef.fieldID,
		formulaText: formulaText
	}
	
	jsonAPIRequest("calcField/validateFormula",validationParams,function(validationResponse) {
		if(validationResponse.isValidFormula) {
			console.log("formula validation successful")
			$('#formulaEditorErrorPopup').popover('hide')
			
			validationSucceededCallback(fieldRef,formulaText)
		} else {
			console.log("formula validation failed: " + validationResponse.errorMsg)
			$("#formulaEditorErrorPopup").attr("data-content",validationResponse.errorMsg);
			
			$('#formulaEditorErrorPopup').popover('show')
			
		}
	})
	
}

function saveFormula(fieldRef) {
	
	validateFormula(fieldRef,function(fieldRef,formulaText) {
		var saveFormulaParms = {
			fieldID: fieldRef.fieldID,
			formulaText: formulaText
		}
		jsonAPIRequest("calcField/setFieldFormula", saveFormulaParms, function(updatedFieldRef) {
			console.log("Saved formula: updated field = " + JSON.stringify(updatedFieldRef))				
		})
	})
	
}

function openFormulaEditor(fieldRef) {

	$('#formulaEditorErrorPopup').popover('hide')
	
	var getRawFormulaSrcParams = { fieldID: fieldRef.fieldID }
	jsonAPIRequest("calcField/getRawFormulaText",getRawFormulaSrcParams,function(formulaInfo) {
		formulaEditorConfig.editor.setValue(formulaInfo.rawFormulaText)
		
		formulaEditorConfig.showEditorFunc()
	
		$('#saveFormulaButton').unbind('click');
		$('#saveFormulaButton').click(function(e){
			e.preventDefault();
			console.log("save button clicked")
			saveFormula(fieldRef)	
		});
	
		$('#checkFormulaButton').unbind('click');
		$('#checkFormulaButton').click(function(e){	
			e.preventDefault();
			console.log("check formula button clicked")
			validateFormula(fieldRef,function(fieldRef,formulaText) {})
		});
		
		
	})


	
}

function closeFormulaEditor() {
	$('#checkFormulaButton').popover('hide')
	
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