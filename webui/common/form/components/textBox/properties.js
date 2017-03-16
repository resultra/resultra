function loadTextBoxProperties($textBox,textBoxRef) {
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
				setContainerComponentInfo($textBox,updatedTextBox,updatedTextBox.textBoxID)
			})	
			
		}
	}
	initNumberFormatSelection(formatSelectionParams)
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentFormID: textBoxRef.parentFormID,
			textboxID: textBoxRef.textBoxID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/textBox/setLabelFormat", formatParams, function(updatedTextBox) {
			setTextBoxComponentLabel($textBox,updatedTextBox)
			setContainerComponentInfo($textBox,updatedTextBox,updatedTextBox.textBoxID)
		})	
	}
	
	
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: textBoxRef.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#textBoxProps')
		
	toggleFormulaEditorForField(textBoxRef.properties.fieldID)
		
}