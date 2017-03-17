
function loadHtmlEditorProperties($editor, htmlEditorRef) {
	console.log("Loading html editor properties")
	
	var elemPrefix = "htmlEditor_"
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for html editor")
		var formatParams = {
			parentFormID: htmlEditorRef.parentFormID,
			htmlEditorID: htmlEditorRef.htmlEditorID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/htmlEditor/setLabelFormat", formatParams, function(updatedEditor) {
			setEditorComponentLabel($editor,updatedEditor)
			setContainerComponentInfo($editor,updatedEditor,updatedEditor.htmlEditorID)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: htmlEditorRef.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)

	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#htmlEditorProps')
		
	toggleFormulaEditorForField(htmlEditorRef.properties.fieldID)
	
}