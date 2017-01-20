function loadTextBoxProperties(textBoxRef) {
	console.log("loading text box properties")
	
	var elemPrefix = "textBox_"
	
	var formatSelectionParams = {
		elemPrefix: elemPrefix,
		initialFormat: textBoxRef.properties.valueFormat.format,
		formatChangedCallback: function (newFormat) {
			console.log("Format changed for text box: " + newFormat)
			
			var newValueFormat = {
				format: newFormat
			}
			var formatParams = {
				parentFormID: textBoxRef.parentFormID,
				textboxID: textBoxRef.textBoxID,
				valueFormat: newValueFormat
			}
			jsonAPIRequest("frm/textBox/setValueFormat", formatParams, function(updatedTextBox) {
				setElemObjectRef(updatedTextBox.textBoxID,updatedTextBox)
			})	
			
		}
	}
	initNumberFormatSelection(formatSelectionParams)
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#textBoxProps')
		
	toggleFormulaEditorForComponent(textBoxRef.properties.componentLink)
		
}