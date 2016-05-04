
// Only one formula editor is expected to be initialized/configured per page, so
// a global configuration should suffice.
var formulaEditorConfig;

function populateFieldRefInsertionMenu(tableID)
{
	// Populate the menu to insert field references with the list of fields
	$("#formulaFieldRefList").empty()
	loadFieldInfo(tableID, [fieldTypeAll],function(fieldsByID) {
		for (var fieldID in fieldsByID) {
		
			var fieldInfo = fieldsByID[fieldID]		
		
     	   var menuItemHTML = '<div class="item" data-value="' + fieldInfo.refName + 
				'">' + fieldInfo.name + '</div>'
				
		 	$("#formulaFieldRefList").append(menuItemHTML)			

		} // for each  field
		$("#formulaFieldRefSelector").dropdown({
			onChange: function(fieldRefName,text,$choice) {
				console.log("formula edit dropdown selection: " + fieldRefName)
				if(fieldRefName.length > 0) {
					$("#formulaFieldRefSelector").dropdown('restore default text')
					formulaEditorConfig.editor.insert("[" + fieldRefName + "]")
				}	
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
	
	$('#formulaEditor').popup({on: 'manual'})
	
	$('#formulaEditMoreDropdown').dropdown()
	
	$('#formulaErrorMsgClose').on('click', function() {
		$('#formulaEditor').popup('hide')
    })
	
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
			$('#formulaEditor').popup('hide')
			validationSucceededCallback(fieldRef,formulaText)
		} else {
			console.log("formula validation failed: " + validationResponse.errorMsg)
			$("#formulaErrorMessageMsgText").text(validationResponse.errorMsg);
			$('#formulaEditor').popup('show')
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